package api

import (
	"encoding/json"
	"fmt"
	"github.com/yangwoodstar/NovaCore/src/httpClient"
	"github.com/yangwoodstar/NovaCore/src/modelStruct"
)

func SendWaningMessage(url, message, mobile string) (string, error) {

	msg := fmt.Sprintf("%s  \n@%s", message, mobile)
	markdown := modelStruct.Markdown{
		Title: "warning",
		Text:  msg,
	}
	at := modelStruct.At{
		AtMobiles: []string{mobile},
		IsAtAll:   false,
	}
	webhookMessage := modelStruct.WebhookMessage{
		Msgtype:  "markdown",
		Markdown: markdown,
		At:       at,
	}

	jsonData, err := json.Marshal(webhookMessage)
	if err != nil {
		return "", err
	}

	// 这里调用了httpApi包中的函数
	response, err := httpClient.ProcessPost(url, string(jsonData), "", "", "")
	if err != nil {
		return "", err
	}
	return string(response), nil
}

func ServerRestart(url, mobile, name, id, startTime string) (string, error) {
	sendWarning := fmt.Sprintf("服务启动通知    \n 服务名称: %s      \n 服务标识: %s      \n 启动时间:%s      \n", name, id, startTime)
	return SendWaningMessage(url, sendWarning, mobile)
}
