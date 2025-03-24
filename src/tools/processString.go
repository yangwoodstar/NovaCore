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
