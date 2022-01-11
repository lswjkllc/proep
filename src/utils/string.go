package utils

import "strings"

func JoinStrings(strs ...string) string {
	// 特殊值判断
	switch len(strs) {
	case 0:
		return ""
	case 1:
		return strs[0]
	}
	// 循环计算最终字符串长度
	n := 0
	for i := 0; i < len(strs); i++ {
		n += len(strs[i])
	}
	// 声明
	var b strings.Builder
	// 预分配内存
	b.Grow(n)
	// 循环添加字符串
	for _, str := range strs {
		b.WriteString(str)
	}
	return b.String()
}
