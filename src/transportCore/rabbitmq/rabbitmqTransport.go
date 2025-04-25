package rabbitmq

import (
	"errors"
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/yangwoodstar/NovaCore/src/transportCore"
	"go.uber.org/zap"
	"time"
)

type ConfigRabbitMQInfo struct {
	ID                 string
	Exchange           string
	Kind               string
	Queue              string
	BindingKey         string
	ExchangeDurable    bool
	ExchangeAutoDelete bool
	QueueDurable       bool
	QueueAutoDelete    bool
	Priority           int
	AutoAck            bool
	PrefetchCount      int
}

type TransportRabbitMQ struct {
	amqpURI       string
	conn          *amqp.Connection
	channel       *amqp.Channel
	senders       map[string]string
	receivers     map[string]<-chan amqp.Delivery
	sendersInfo   map[string]*ConfigRabbitMQInfo
	receiversInfo map[string]*ConfigRabbitMQInfo
	msgChan       chan amqp.Delivery
	errorChan     chan error
	logger        *zap.Logger
	reConnect     bool
}

// NewTransportRabbitMQ 创建一个新的 TransportRabbitMQ
func NewTransportRabbitMQ(amqpURI string, logger *zap.Logger) (*TransportRabbitMQ, error) {
	transportRabbitMQ := &TransportRabbitMQ{
		amqpURI:       amqpURI,
		senders:       make(map[string]string),
		receivers:     make(map[string]<-chan amqp.Delivery),
		sendersInfo:   make(map[string]*ConfigRabbitMQInfo),
		receiversInfo: make(map[string]*ConfigRabbitMQInfo),
		msgChan:       make(chan amqp.Delivery, 1000),
		errorChan:     make(chan error),
		reConnect:     false,
		logger:        logger,
	}

	if err := transportRabbitMQ.Connect(); err != nil {
		logger.Error("Connect error", zap.Error(err))
		return nil, err
	}

	return transportRabbitMQ, nil
}

func (rt *TransportRabbitMQ) Connect() error {
	rt.logger.Info("Connect", zap.String("amqpURI", rt.amqpURI))
	var err error
	/*	dialConfig := amqp.Config{
			Properties: amqp.Table{"connection_timeout": int32(3 * time.Second / time.Millisecond)},
			Heartbeat:  10 * time.Second,
		}
	*/
	for {
		if rt.conn, err = amqp.Dial(rt.amqpURI); err == nil {
			rt.logger.Info("Connected to RabbitMQ")
			break
		}
		rt.logger.Error("Error in creating RabbitMQ connection", zap.Error(err))
		time.Sleep(3 * time.Second)
	}

	rt.channel, err = rt.conn.Channel()
	if err != nil {
		rt.logger.Error("Channel error", zap.Error(err), zap.String("amqpURI", rt.amqpURI))
		return fmt.Errorf("Channel: %w", err)
	}

	go rt.monitorConnection()
	go rt.monitorReconnect()

	return nil
}

func (rt *TransportRabbitMQ) SetPrefetchCount(prefetchCount int) {
	rt.logger.Info("SetPrefetchCount", zap.Int("prefetchCount", prefetchCount))
	err := rt.channel.Qos(
		prefetchCount, // prefetch count
		0,             // prefetch size (0表示不限制)
		false,         // global设置 (false表示仅当前消费者生效)
	)

	if err != nil {
		rt.logger.Error("Qos error", zap.Error(err))
	}
}

func (rt *TransportRabbitMQ) monitorConnection() {
	<-rt.conn.NotifyClose(make(chan *amqp.Error))
	rt.logger.Info("Connection Closed")
	rt.errorChan <- errors.New("Connection Closed")
}

func (rt *TransportRabbitMQ) monitorReconnect() {
	if err := <-rt.errorChan; err != nil {
		rt.reConnect = true
		rt.logger.Info("Start reconnect consuming")
		if err := rt.Reconnect(); err != nil {
			rt.logger.Error("Reconnect error", zap.Error(err))
		}
	}
	rt.logger.Info("Reconnecting finish")
}

func (rt *TransportRabbitMQ) Reconnect() error {
	rt.logger.Info("Reconnect")
	if !rt.conn.IsClosed() {
		rt.Close()
	}

	if err := rt.Connect(); err != nil {
		rt.logger.Error("Reconnect error", zap.Error(err))
		return err
	}

	for _, configInfo := range rt.sendersInfo {
		if err := rt.AddSender(configInfo); err != nil {
			rt.logger.Error("AddSender error", zap.Error(err))
			return err
		}
	}

	for _, configInfo := range rt.receiversInfo {
		if err := rt.AddReceiver(configInfo); err != nil {
			rt.logger.Error("AddReceiver error", zap.Error(err))
			return err
		}
	}

	return nil
}

func (rt *TransportRabbitMQ) AddSender(configInfo *ConfigRabbitMQInfo) error {
	rt.logger.Info("AddSender", zap.String("id", configInfo.ID), zap.Any("configInfo", configInfo))
	rt.sendersInfo[configInfo.Exchange] = configInfo
	rt.senders[configInfo.Exchange] = configInfo.Exchange
	return rt.declareExchange(configInfo, rt.channel)
}

func (rt *TransportRabbitMQ) AddReceiver(configInfo *ConfigRabbitMQInfo) error {
	var err error
	var channel = rt.channel
	if configInfo.PrefetchCount > 0 {
		channel, err = rt.conn.Channel()
		if err != nil {
			rt.logger.Error("Channel error", zap.Error(err), zap.String("amqpURI", rt.amqpURI))
			return fmt.Errorf("Channel: %w", err)
		}

		err = channel.Qos(
			configInfo.PrefetchCount, // prefetch count
			0,                        // prefetch size (0表示不限制)
			false,                    // global设置 (false表示仅当前消费者生效)
		)

	}
	rt.logger.Info("AddReceiver", zap.String("id", configInfo.ID), zap.Any("configInfo", configInfo))
	rt.receiversInfo[configInfo.Exchange] = configInfo
	rt.senders[configInfo.Exchange] = configInfo.Exchange
	if err := rt.declareExchange(configInfo, channel); err != nil {
		return err
	}

	if err := rt.bindQueue(configInfo, channel); err != nil {
		return err
	}

	return rt.consumeMessages(configInfo, channel)
}

func (rt *TransportRabbitMQ) declareExchange(configInfo *ConfigRabbitMQInfo, channel *amqp.Channel) error {
	if err := channel.ExchangeDeclare(configInfo.Exchange, configInfo.Kind, configInfo.ExchangeDurable, configInfo.ExchangeAutoDelete, false, false, nil); err != nil {
		rt.logger.Error("ExchangeDeclare error", zap.Error(err), zap.Any("configInfo", configInfo))
		return err
	}
	return nil
}

func (rt *TransportRabbitMQ) bindQueue(configInfo *ConfigRabbitMQInfo, channel *amqp.Channel) error {
	var args amqp.Table = nil
	if configInfo.Priority > 0 {
		args = amqp.Table{
			"x-max-priority": configInfo.Priority,
		}
	}

	if _, err := channel.QueueDeclare(configInfo.Queue, configInfo.QueueDurable, configInfo.QueueAutoDelete, false, false, args); err != nil {
		rt.logger.Error("Queue Declare error", zap.Error(err), zap.Any("configInfo", configInfo))
		return fmt.Errorf("error in declaring the queue: %w", err)
	}

	// 使用整数作为绑定键

	if err := channel.QueueBind(configInfo.Queue, configInfo.BindingKey, configInfo.Exchange, false, nil); err != nil {
		rt.logger.Error("Queue Bind error", zap.Error(err), zap.Any("configInfo", configInfo))
		return fmt.Errorf("Queue Bind error: %w", err)
	}
	return nil
}

func (rt *TransportRabbitMQ) consumeMessages(configInfo *ConfigRabbitMQInfo, channel *amqp.Channel) error {
	deliveries, err := channel.Consume(configInfo.Queue, configInfo.Queue, configInfo.AutoAck, false, false, false, nil)
	if err != nil {
		rt.logger.Error("Consume error", zap.Error(err), zap.Any("configInfo", configInfo))
		return fmt.Errorf("Consume error: %w", err)
	}

	rt.receivers[configInfo.ID] = deliveries

	go func(msg <-chan amqp.Delivery) {
		for m := range msg {
			if rt.reConnect && string(m.Body) == "" {
				rt.logger.Error("connection closed", zap.Any("msg", m), zap.Any("configInfo", configInfo))
				return
			}
			rt.logger.Info("Read message", zap.String("msg", string(m.Body)))
			rt.msgChan <- m
		}
		rt.logger.Info("Close msg channel", zap.Any("configInfo", configInfo))
		rt.Close()
	}(deliveries)

	return nil
}

// Read 从消息通道中读取一条消息
func (rt *TransportRabbitMQ) Read() (transportCore.UnificationMessage, error) {
	select {
	case msg := <-rt.msgChan:
		return &RabbitMQMessage{message: msg, topic: msg.Exchange}, nil
	}
}

// Write 向 RabbitMQ 发送消息
func (rt *TransportRabbitMQ) Write(msg []byte, exchange, routerKey string, priority int) error {
	rt.logger.Info("Write message", zap.String("exchange", exchange), zap.String("routerKey", routerKey), zap.String("msg", string(msg)))
	var p amqp.Publishing
	if priority != 0 {
		p = amqp.Publishing{
			Headers:     amqp.Table{"type": "text/plain"},
			ContentType: "text/plain",
			Body:        msg,
			Priority:    uint8(priority),
		}
	} else {
		p = amqp.Publishing{
			Headers:     amqp.Table{"type": "text/plain"},
			ContentType: "text/plain",
			Body:        msg,
		}
	}

	_, exists := rt.senders[exchange]
	if exists {
		err := rt.channel.Publish(exchange, routerKey, false, false, p)
		if err != nil {
			rt.logger.Error("Error in Publishing", zap.Error(err), zap.Any("msg", msg), zap.String("exchange", exchange))
			// 这里可以选择重试或者其他处理方式
			return fmt.Errorf("Error in Publishing: %w", err)
		}
	} else {
		rt.logger.Error("Exchange not found", zap.String("exchange", exchange), zap.String("msg", string(msg)))
	}

	return nil
}

func (rt *TransportRabbitMQ) publishMessage(exchange, routerKey string, p amqp.Publishing) error {
	_, exists := rt.senders[exchange]
	if exists {
		err := rt.channel.Publish(exchange, routerKey, false, false, p)
		if err != nil {
			rt.logger.Error("Error in Publishing", zap.Error(err))
			return fmt.Errorf("Error in Publishing: %w", err)
		}
	}
	return nil
}

// Close 关闭 TransportRabbitMQ
func (rt *TransportRabbitMQ) Close() {
	rt.logger.Info("Close ******************************")
	if err := rt.conn.Close(); err != nil {
		rt.logger.Error("Close error", zap.Error(err))
	}
}
