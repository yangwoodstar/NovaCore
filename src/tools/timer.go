package tools

import (
	"crypto/rand"
	"github.com/google/uuid"
	"time"
)

func GetTimeStamp() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

func GetSecondTimeStamp() int64 {
	return time.Now().Unix()
}
func GenerateUUID() string {
	// 生成一个新的 UUID
	uuidInstance := uuid.New()
	// 将 UUID 转换为字符串
	return uuidInstance.String()
}

func SwitchToTime(timestamp string) (int64, error) {
	layout := time.RFC3339 // 使用 RFC3339 格式
	t, err := time.Parse(layout, timestamp)
	if err != nil {
		return 0, err
	}

	// 转换为时间戳（毫秒）
	timestampMillis := t.UnixNano() / int64(time.Millisecond)
	return timestampMillis, nil
}

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func GetRandomString(length int) (string, error) {
	byteArray := make([]byte, length)
	_, err := rand.Read(byteArray)
	if err != nil {
		return "", err
	}

	// 将随机字节映射到字符集
	for i := range byteArray {
		byteArray[i] = charset[int(byteArray[i])%len(charset)]
	}

	return string(byteArray), nil
}

// 获取格式化时间字符串（默认本地时区）
func GetTimeString(format string) string {
	if format == "" {
		format = "2006-01-02 15:04:05" // 设置默认格式
	}
	return time.Now().Format(format)
}

// 获取UTC时间字符串（带时区信息）
func GetUTCTimeString(format string) string {
	if format == "" {
		format = "2006-01-02T15:04:05Z07:00" // ISO8601格式
	}
	return time.Now().UTC().Format(format)
}
