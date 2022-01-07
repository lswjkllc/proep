package join

import "testing"

func TestJoin(t *testing.T) {
	// 基础字符串
	var base = "123456789qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASFGHJKLZXCVBNM"
	// 生成长切片
	var baseSlice []string
	for i := 0; i < 200; i++ {
		baseSlice = append(baseSlice, base)
	}
	// 开始测试
	t.Log(Sprintf(baseSlice))
	t.Log(StringBuilder(baseSlice...))
	t.Log(StringJoin(baseSlice...))
}
