package tools

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
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

func GetFileSize(filePath string) (int64, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return 0, err
	}
	// 确保在函数结束时关闭文件
	defer func(file *os.File) {
		err = file.Close()
		if err != nil {
			fmt.Println("Error closing file:", err)
		}
	}(file)

	// 获取文件信息
	fileInfo, err := file.Stat()
	if err != nil {
		return 0, err
	}
	// 获取文件大小
	return fileInfo.Size(), nil
}

func RemoveFile(filePath string) error {
	err := os.Remove(filePath)
	if err != nil {
		return err
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
