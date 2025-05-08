package tools

import "strings"

// 带缓冲区的优化版本（适合高频调用）
func JoinStrings(s ...string) string {
	var builder strings.Builder
	for _, str := range s {
		builder.WriteString(str)
	}
	return builder.String()
}

func ExtractPath(fullPath string) string {
	parts := strings.Split(fullPath, "/")
	if len(parts) == 0 {
		return fullPath
	}
	// 去除最后一个部分（如playlist.mpd）
	return strings.Join(parts[:len(parts)-1], "/")
}

func TruncateString(s string, maxLen int) string {
	// 将字符串转换为rune切片（正确处理多字节字符）
	runes := []rune(s)

	// 边界检查
	if maxLen <= 0 {
		return ""
	}

	// 截断逻辑
	if len(runes) > maxLen {
		return string(runes[:maxLen])
	}
	return s
}
