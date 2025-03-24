package tools

import (
	"fmt"
	"os"
)

func FileExists(path string) (bool, error) {
	// 获取文件信息
	_, err := os.Stat(path)

	// 文件存在的情况
	if err == nil {
		return true, nil
	}

	// 文件不存在的情况
	if os.IsNotExist(err) {
		return false, nil
	}

	// 其他错误（如权限问题）
	return false, err
}

func RemoveDir(dirPath string) error {
	err := os.RemoveAll(dirPath)
	if err != nil {
		return err
	}
	return nil
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
