package services

const (
	DEFAULT_OFFSET = 0
	DEFAULT_LIMIT  = 10
)

func getPageInfo(data map[string]interface{}) (int, int) {
	// 声明
	var offset, limit int
	// 偏移
	ioffset, ok := data["offset"]
	if ok {
		offset = int(ioffset.(float64))
	} else {
		offset = DEFAULT_OFFSET
	}
	// 限制
	ilimit, ok := data["limit"]
	if ok {
		limit = int(ilimit.(float64))
	} else {
		limit = DEFAULT_LIMIT
	}
	return offset, limit
}
