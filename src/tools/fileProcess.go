package tools

import "os"

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
