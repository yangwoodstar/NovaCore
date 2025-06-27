package api

import (
	"encoding/json"
	"fmt"
	"github.com/yangwoodstar/NovaCore/src/constString"
	"github.com/yangwoodstar/NovaCore/src/httpClient"
	"github.com/yangwoodstar/NovaCore/src/modelStruct"
)

// 级别颜色映射
var LevelColor = map[string]string{
	"P0": "#FF0000", // 红色
	"P1": "#FFA500", // 橙色
	"P2": "#808080", // 黄色
	"P3": "#00FF00", // 灰色
}

var LevelTips = map[string]string{
	"P0": "紧急", // 红色
	"P1": "严重", // 橙色
	"P2": "告警", // 黄色
	"P3": "提示", // 灰色
}

func SendWaningMessage(url, message, mobile string) (string, error) {

	msg := fmt.Sprintf("%s  \n@%s", message, mobile)
	markdown := modelStruct.Markdown{
		Title: "warning [P1-严重] 服务响应超时",
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

func formatAtMobiles(mobiles []string) string {
	var result string
	for _, m := range mobiles {
		result += fmt.Sprintf("@%s ", m)
	}
	return result
}

func GenerateAlert(warningInfo modelStruct.WarningInfo) modelStruct.WebhookMessage {
	color, ok := LevelColor[warningInfo.Level]
	if !ok {
		color = LevelColor[constString.P3] // 默认使用P3颜色
	}

	tip, exist := LevelTips[warningInfo.Level]
	if !exist {
		tip = LevelTips[constString.P3]
	}

	text := fmt.Sprintf(
		"### <font color=%s>%s-%s</font>：%s\n\n"+
			"**环境**：%s\n\n"+
			"**时间**：%s\n\n"+
			"**详情**：\n> %s\n\n"+
			"**处理建议**：\n> %s\n\n"+
			"**负责人**：\n> %s\n\n"+
			"**详细信息**：\n>%s\n\n",
		color, warningInfo.Level, tip, warningInfo.Title,
		warningInfo.Env,
		warningInfo.Time,
		warningInfo.Details,
		warningInfo.Advice,
		formatAtMobiles(warningInfo.Owners),
		warningInfo.Message,
	)

	return modelStruct.WebhookMessage{
		Msgtype: "markdown",
		Markdown: modelStruct.Markdown{
			Title: fmt.Sprintf("%swarning：%s", warningInfo.Level, warningInfo.Title),
			Text:  text,
		},
		At: modelStruct.At{
			AtMobiles: warningInfo.Owners,
			IsAtAll:   false,
		},
	}
}

func SendDingTalkAlert(webhookURL string, alert modelStruct.WebhookMessage) (string, error) {
	payload, err := json.Marshal(alert)
	if err != nil {
		return "", err
	}

	response, err := httpClient.ProcessPost(webhookURL, string(payload), "", "", "")
	if err != nil {
		return "", err
	}

	return string(response), nil
}

func ServerRestart(url, mobile, name, id, startTime string) (string, error) {
	sendWarning := fmt.Sprintf("服务启动通知    \n 服务名称: %s      \n 服务标识: %s      \n 启动时间:%s      \n", name, id, startTime)
	return SendWaningMessage(url, sendWarning, mobile)
}
