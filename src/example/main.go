package main

import (
	"fmt"
	"github.com/yangwoodstar/NovaCore/src/api"
	"github.com/yangwoodstar/NovaCore/src/constString"
	"github.com/yangwoodstar/NovaCore/src/core/instanceAllocator"
	"github.com/yangwoodstar/NovaCore/src/httpClient"
	"github.com/yangwoodstar/NovaCore/src/modelStruct"
	"github.com/yangwoodstar/NovaCore/src/transportCore"
	"github.com/yangwoodstar/NovaCore/src/transportCore/kafka"
	"github.com/yangwoodstar/NovaCore/src/transportCore/rabbitmq"
	"go.uber.org/zap"
	"log"
	"strconv"
	"time"
)

func Test() {
	config := zap.NewProductionConfig()
	config.Level = zap.NewAtomicLevelAt(zap.DebugLevel) // 设置日志级别为 Debug

	// 使用配置创建 logger
	logger, _ := config.Build()
	defer logger.Sync()
	var unifiedTransport *transportCore.UnifiedTransport
	rabbitMQUri := "amqp://" + "" + ":" + "" + "@" + "" + ":" + strconv.Itoa(5672) + "/"

	rabbitmqInstance, err := rabbitmq.NewTransportRabbitMQ(rabbitMQUri, logger)
	if err != nil {
		logger.Error("Failed to create RabbitMQ transport", zap.Error(err))
		return
	}
	broke := []string{""}
	kafkaInstance, err := kafka.GetKafkaProducer(broke, 2, logger)
	if err != nil {
		logger.Error("Failed to create Kafka transport", zap.Error(err))
		return
	}
	send01Config := &rabbitmq.ConfigRabbitMQInfo{
		ID:                 "1",
		Exchange:           "test01",
		Kind:               "fanout",
		Queue:              "test01",
		BindingKey:         "",
		ExchangeDurable:    true,
		ExchangeAutoDelete: true,
		QueueDurable:       true,
		QueueAutoDelete:    true,
		Priority:           0,
	}
	err = rabbitmqInstance.AddSender(send01Config)
	if err != nil {
		logger.Error("Failed to add sender", zap.Error(err))
		return
	}
	send02Config := &rabbitmq.ConfigRabbitMQInfo{
		ID:                 "2",
		Exchange:           "test02",
		Kind:               "fanout",
		Queue:              "test02",
		BindingKey:         "",
		ExchangeDurable:    true,
		ExchangeAutoDelete: true,
		QueueDurable:       true,
		QueueAutoDelete:    true,
		Priority:           0,
	}
	err = rabbitmqInstance.AddSender(send02Config)
	if err != nil {
		logger.Error("Failed to add sender", zap.Error(err))
		return
	}
	receiver01Config := &rabbitmq.ConfigRabbitMQInfo{
		ID:                 "2",
		Exchange:           "test02",
		Kind:               "fanout",
		Queue:              "test02",
		BindingKey:         "",
		ExchangeDurable:    true,
		ExchangeAutoDelete: true,
		QueueDurable:       true,
		QueueAutoDelete:    true,
		Priority:           0,
	}
	err = rabbitmqInstance.AddReceiver(receiver01Config)
	if err != nil {
		logger.Error("Failed to add receiver", zap.Error(err))
		return
	}
	receiver02Config := &rabbitmq.ConfigRabbitMQInfo{
		ID:                 "1",
		Exchange:           "test01",
		Kind:               "fanout",
		Queue:              "test01",
		BindingKey:         "",
		ExchangeDurable:    true,
		ExchangeAutoDelete: true,
		QueueDurable:       true,
		QueueAutoDelete:    true,
		Priority:           0,
	}
	err = rabbitmqInstance.AddReceiver(receiver02Config)
	if err != nil {
		logger.Error("Failed to add receiver", zap.Error(err))
		return
	}

	unifiedTransport = transportCore.NewUnifiedTransport()
	unifiedTransport.AddSender("test01", rabbitmqInstance)
	unifiedTransport.AddSender("test02", rabbitmqInstance)
	unifiedTransport.AddSender("testconsistent01", rabbitmqInstance)
	unifiedTransport.AddSender("testconsistent02", rabbitmqInstance)
	unifiedTransport.AddSender("", kafkaInstance)
	unifiedTransport.AddReceiver("mq", rabbitmqInstance)

	go func() {
		// Create a ticker that fires every second
		ticker := time.NewTicker(3 * time.Second)
		defer ticker.Stop() // Ensure ticker is stopped when goroutine exits
		index := 0
		for {
			// Wait for the next tick
			<-ticker.C

			// Send messages
			unifiedTransport.Write([]byte("test01"), "test01", "test01", 0)
			unifiedTransport.Write([]byte("test02"), "test02", "test02", 0)
			//unifiedTransport.Write([]byte("testconsistent01"), "testconsistent01", "consistent01", 0)
			//unifiedTransport.Write([]byte("testconsistent02"), "testconsistent02", "consistent02", 0)
			unifiedTransport.Write([]byte(strconv.Itoa(index)), "", "123", 0)
			index++
			// Optional: Add logging to confirm messages are sent
			fmt.Println("Messages sent at:", time.Now())
		}
	}()

	for {
		msg, readErr := unifiedTransport.Read()
		if readErr != nil {
			logger.Error("Error reading message", zap.Error(readErr))
			continue
		}
		logger.Info("Received a message", zap.String("msg", string(msg.GetBody())), zap.String("exchange", msg.GetTopic()))
	}
	defer unifiedTransport.Close()
	select {}
}

func TestAppID() {
	//InitConfig()
	appIDMap := map[string]instanceAllocator.AppIDMapConfig{
		"app1": {
			AppID:  "app1_id",
			AppKey: "app1_key",
		},
		"app2": {
			AppID:  "app2_id",
			AppKey: "app2_key",
		},
	}

	// 初始化 InstanceManager
	ak := "your_access_key"
	sk := "your_secret_key"
	region := "your_region"
	manager := instanceAllocator.GetInstanceManager(appIDMap, ak, sk, region)

	// 测试获取实例
	key := "app1"
	instance, err := manager.GetInstance(key)
	if err != nil {
		log.Fatalf("Failed to get instance for key %s: %v", key, err)
	}
	fmt.Printf("Instance for key %s created successfully: %+v\n", key, instance.Config)

	// 测试获取另一个实例
	key = "app2"
	instance, err = manager.GetInstance(key)
	if err != nil {
		log.Fatalf("Failed to get instance for key %s: %v", key, err)
	}
	fmt.Printf("Instance for key %s created successfully: %+v\n", key, instance.Config)

	// 测试获取不存在的实例
	key = "app3"
	instance, err = manager.GetInstance(key)
	if err != nil {
		fmt.Printf("Expected error for key %s: %v\n", key, err)
	} else {
		log.Fatalf("Unexpected success for key %s: %+v", key, instance.Config)
	}

	instance, err = manager.GetAppIDInstance("app1_id")
	if err != nil {
		fmt.Printf("Expected error for key %s: %v\n", key, err)
	} else {
		log.Fatalf("Unexpected success for key 1111111 %s: %+v", key, instance.Config)
	}

	// 列出所有实例
	keys := manager.ListInstances()
	fmt.Printf("All instances: %v\n", keys)

	// 移除一个实例
	manager.RemoveInstance("app1")
	keys = manager.ListInstances()
	fmt.Printf("Instances after removal: %v\n", keys)
}

type DingTalkMessage struct {
	MsgType string `json:"msgtype"`
	Message string `json:"markdown"`
}

func TestDingTalk() {
	//url := "https://oapi.dingtalk.com/robot/send?access_token=your_access_token"
	url := ""

	/*	mobile := ""
			dingTalkMessage := DingTalkMessage{
				MsgType: "live",
				Message: "this is a test message",
			}

			message, err := json.Marshal(dingTalkMessage)
			if err != nil {
				log.Fatalf("Failed to marshal message: %v", err)
			}

		message = []byte("### 测试消息\nHello, this is a test message.") // 直接传递Markdown内容
		message = []byte("### <font color=orange>[P1-严重]</font> 订单服务响应超时\\n> **影响范围**: 用户支付功能\\n> **建议操作**: 检查网关及服务日志") // 直接传递Markdown内容
		response, err := api.SendWaningMessage(url, string(message), mobile)
		if err != nil {
			log.Fatalf("Failed to send message: %v", err)
		}
		fmt.Printf("Response: %s\n", response)
	*/
	warningInfo := modelStruct.WarningInfo{
		Level:   constString.P0,
		Title:   "测试标题",
		Time:    time.Now().Format(time.RFC3339),
		Details: "测试详情",
		Advice:  "测试建议",
		Env:     "测试环境",
		Message: "测试消息",
		Owners:  []string{"1234567890", "1234567890"},
	}
	alertMsg := api.GenerateAlert(warningInfo)
	res, err := api.SendDingTalkAlert(url, alertMsg)
	if err != nil {
		log.Fatalf("Failed to send message: %v", err)
	}
	fmt.Printf("Response: %s\n", res)

}

func WarningTest() {

}

func TestHttpMethod() {
	url := "http://172.17.56.159:3000/getToken"
	response, err := httpClient.ProcessPost(url, "", "", "", "")
	if err != nil {
		fmt.Printf("Failed to send POST request: %v\n", err)
	} else {
		fmt.Printf("Response: %s\n", response)
	}

	url = "http://172.17.56.159:3000/template1"
	response, err = httpClient.ProcessGet(url, "", "", "", nil)
	if err != nil {
		fmt.Printf("Failed to send GET request: %v\n", err)
	} else {
		fmt.Printf("Response: %s\n", response)
	}

}
func main() {
	//test.CreateLiveApiTest()
	//test.ListLiveApiTest()
	//test.DeleteLiveApiTest()
	//Test()
	//TestDingTalk()
	TestHttpMethod()
}
