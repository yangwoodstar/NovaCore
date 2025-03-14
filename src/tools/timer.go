package tools

import (
	"crypto/rand"
	"fmt"
	"github.com/google/uuid"
	"io"
	"os"
	"path/filepath"
	"time"
)

func GetTimeStamp() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

func GenerateUUID() string {
	// 生成一个新的 UUID
	uuidInstance := uuid.New()
	// 将 UUID 转换为字符串
	return uuidInstance.String()
}

func ExistDirectory(path string) error {
	// 检查目录是否存在
	// 检查目录是否存在
	if _, err := os.Stat(path); os.IsNotExist(err) {
		// 目录不存在，创建多级目录
		err = os.MkdirAll(path, 0755) // 0755 是目录的权限
		if err != nil {
			return err
		}
	}
	return nil
}

// AppendToFile appendToFile 将数据追加到文件末尾
func AppendToFile(filePath string, data []byte) error {
	f, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer func(f *os.File) {
		err = f.Close()
		if err != nil {
			return
		}
	}(f)

	if _, err = f.Write(data); err != nil {
		return err
	}
	return nil
}

// CopyFile appendToFile 将数据追加到文件末尾
// CopyFile 拷贝文件到指定目录
func CopyFile(src string, dstDir string) error {
	// 打开源文件
	sourceFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer func(sourceFile *os.File) {
		err = sourceFile.Close()
		if err != nil {
			return
		}
	}(sourceFile)

	// 获取源文件的基本名称
	fileName := filepath.Base(src)
	// 创建目标文件的完整路径
	dst := filepath.Join(dstDir, fileName)

	// 创建目标文件
	destinationFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer func(destinationFile *os.File) {
		err = destinationFile.Close()
		if err != nil {
			return
		}
	}(destinationFile)

	// 拷贝文件内容
	_, err = io.Copy(destinationFile, sourceFile)
	if err != nil {
		return err
	}

	// 复制文件的权限
	sourceInfo, err := sourceFile.Stat()
	if err != nil {
		return err
	}
	err = os.Chmod(dst, sourceInfo.Mode())
	if err != nil {
		return err
	}
	return nil
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

// CheckAndCreateDir 检查指定的目录是否存在，如果不存在则创建它。
func CheckAndCreateDir(dirPath string) error {
	// 检查文件夹是否存在
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		// 文件夹不存在，创建它
		err = os.Mkdir(dirPath, 0755) // 0755 是文件夹的权限
		if err != nil {
			return fmt.Errorf("failed to create directory: %w", err)
		}
	}
	return nil
}
