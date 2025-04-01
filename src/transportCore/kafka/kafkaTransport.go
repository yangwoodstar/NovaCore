package kafka

import (
	"context"
	"errors"
	"fmt"
	"github.com/IBM/sarama"
	"github.com/yangwoodstar/NovaCore/src/transportCore"
	"go.uber.org/zap"
	"sync"
	"time"
)

type KafkaProducer struct {
	producer   sarama.AsyncProducer
	logger     *zap.Logger
	brokers    []string
	config     *sarama.Config
	mu         sync.Mutex
	isClosed   bool
	retryDelay time.Duration
	maxRetries int
	ctx        context.Context
	cancel     context.CancelFunc
}

func GetKafkaProducer(brokers []string, partition int, logger *zap.Logger) (*KafkaProducer, error) {
	ctx, cancel := context.WithCancel(context.Background())

	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true

	producer, err := sarama.NewAsyncProducer(brokers, config)
	if err != nil {
		cancel()
		return nil, fmt.Errorf("failed to create producer: %w", err)
	}

	kp := &KafkaProducer{
		producer:   producer,
		logger:     logger,
		brokers:    brokers,
		config:     config,
		retryDelay: 5 * time.Second,
		maxRetries: 3,
		ctx:        ctx,
		cancel:     cancel,
	}

	go kp.handleAsyncResults()
	return kp, nil
}

func (kp *KafkaProducer) handleAsyncResults() {
	for {
		select {
		case success, ok := <-kp.producer.Successes():
			if !ok {
				return
			}
			kp.logger.Info("Message delivered",
				zap.String("topic", success.Topic),
				zap.Int32("partition", success.Partition))
		case err, ok := <-kp.producer.Errors():
			if !ok {
				return
			}
			kp.logger.Error("Send failed",
				zap.Error(err),
				zap.String("topic", err.Msg.Topic))
		case <-kp.ctx.Done():
			return
		}
	}
}

func (kp *KafkaProducer) Read() (transportCore.UnificationMessage, error) {
	kafkaMessage := KafkaMessage{message: "", topic: "msg.Exchange"}
	return &kafkaMessage, nil
}

func (kp *KafkaProducer) Write(message []byte, topic, routerKey string, priority int) error {
	kp.mu.Lock()
	defer kp.mu.Unlock()

	if kp.isClosed {
		return errors.New("producer closed")
	}

	msg := &sarama.ProducerMessage{
		Topic: topic,
		Key:   sarama.StringEncoder(routerKey),
		Value: sarama.ByteEncoder(message),
	}

	select {
	case kp.producer.Input() <- msg:
		return nil
	case <-time.After(100 * time.Millisecond):
		return errors.New("producer queue full")
	case <-kp.ctx.Done():
		return errors.New("producer closing")
	}
}

func (kp *KafkaProducer) Close() {
	kp.mu.Lock()
	defer kp.mu.Unlock()

	if kp.isClosed {
		return
	}

	kp.isClosed = true
	kp.cancel()         // 触发协程退出
	kp.producer.Close() // 关闭生产者
}

func (kp *KafkaProducer) reconnect() error {
	kp.mu.Lock()
	defer kp.mu.Unlock()

	// 终止旧实例
	kp.cancel()
	if err := kp.producer.Close(); err != nil {
		kp.logger.Error("Close old producer failed", zap.Error(err))
	}

	// 创建新实例
	ctx, cancel := context.WithCancel(context.Background())
	producer, err := sarama.NewAsyncProducer(kp.brokers, kp.config)
	if err != nil {
		cancel()
		return fmt.Errorf("reconnect failed: %w", err)
	}

	// 更新实例
	kp.ctx = ctx
	kp.cancel = cancel
	kp.producer = producer

	go kp.handleAsyncResults()
	return nil
}
