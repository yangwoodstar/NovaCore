package tools

import (
	"fmt"
	"net/url"
	"strings"
)

// ExtractPathConfig 配置参数结构体
type ExtractPathConfig struct {
	TrimLeadingSlash bool // 是否去除开头的斜杠
	DecodePath       bool // 是否解码百分比编码（如 %20 → 空格）
}

// ExtractURLPath 从URL字符串中提取路径
// 参数：
//
//	urlStr - 要解析的URL字符串
//	config - 提取配置（传 nil 使用默认配置）
//
// 返回值：
//
//	string - 提取到的路径
//	error  - 解析错误
func ExtractURLPath(urlStr string, config *ExtractPathConfig) (string, error) {
	// 解析URL
	u, err := url.Parse(urlStr)
	if err != nil {
		return "", fmt.Errorf("URL解析失败: %w", err)
	}

	// 处理配置
	var path string
	if config != nil && config.DecodePath {
		// 获取已解码的路径
		path = u.Path
	} else {
		// 获取原始编码路径
		path = u.EscapedPath()
	}

	// 处理开头的斜杠
	if config != nil && config.TrimLeadingSlash {
		path = strings.TrimPrefix(path, "/")
	}

	return path, nil
}
