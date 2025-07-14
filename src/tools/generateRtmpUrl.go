package tools

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type RtmpConfig struct {
	AppID       string
	StreamName  string
	Domain      string
	SecretKey   string
	MaxIdleTime int64
	Format      string
}

func GenerateRtmpUrl(config *RtmpConfig) string {
	// 获取当前的 Unix 时间戳
	unixTime := time.Now().Unix() + config.MaxIdleTime
	// 拼接字符串以生成 volcSecret
	volcTime := fmt.Sprintf("%d", unixTime)
	stringToHash := fmt.Sprintf("/%s/%s%s%s", config.AppID, config.StreamName, config.SecretKey, volcTime)
	// 计算 MD5 哈希值
	hash := md5.Sum([]byte(stringToHash))
	volcSecret := hex.EncodeToString(hash[:])
	// 构建完整的推流地址
	rtmpURL := fmt.Sprintf("%s/%s/%s?volcTime=%s&volcSecret=%s", config.Domain, config.AppID, config.StreamName, volcTime, volcSecret)
	return rtmpURL
}

func GenerateTencentRtmpUrl(config *RtmpConfig) string {
	unixTime := time.Now().Unix() + config.MaxIdleTime
	txTime := strings.ToUpper(strconv.FormatInt(unixTime, 16))
	txSecret := md5.Sum([]byte(config.SecretKey + config.StreamName + txTime))
	txSecretStr := fmt.Sprintf("%x", txSecret)
	rtmpURL := fmt.Sprintf("%s/%s/%s%s?txSecret=%s&txTime=%s", config.Domain, config.AppID, config.StreamName, config.Format, txSecretStr, txTime)
	return rtmpURL
}
