package utils

// HashSetString 字符串hash结构
type HashSetString struct {
	dic  map[string]interface{} // 字典
	Data []string               // 数据
}

// NewHashSetString 初始化
func NewHashSetString() *HashSetString {
	return &HashSetString{
		dic:  make(map[string]interface{}, 0),
		Data: make([]string, 0),
	}
}

// NewHashSetStringWithData 初始化
func NewHashSetStringWithData(data []string) *HashSetString {
	set := &HashSetString{
		dic:  make(map[string]interface{}, 0),
		Data: make([]string, 0),
	}
	for i := range data {
		set.Add(data[i])
	}
	return set
}

// Add 新增元素
func (h *HashSetString) Add(ele string) bool {
	_, ok := h.dic[ele]
	if ok {
		return false
	}
	h.dic[ele] = ele
	h.Data = append(h.Data, ele)
	return true
}

// Contains 是否存在字符串
func (h *HashSetString) Contains(ele string) bool {
	_, ok := h.dic[ele]
	if ok {
		return true
	}
	return false
}

// Get 获取元素
func (h *HashSetString) Get(ele string) string {
	val, ok := h.dic[ele]
	if ok {
		return val.(string)
	}
	return ""
}

// GetList 获取所有字符串
func (h *HashSetString) GetList() []string {
	return h.Data
}

// HashSetInt 数字hash结构
type HashSetInt struct {
	dic  map[int]interface{}
	Data []int
}

// NewHashSetInt 初始化
func NewHashSetInt() *HashSetInt {
	return &HashSetInt{
		dic:  make(map[int]interface{}, 0),
		Data: make([]int, 0),
	}
}

// Add 新增元素
func (h *HashSetInt) Add(ele int) bool {
	_, ok := h.dic[ele]
	if ok {
		return false
	}
	h.dic[ele] = ele
	h.Data = append(h.Data, ele)
	return true
}

// Get 获取元素
func (h *HashSetInt) Get(ele int) int {
	val, ok := h.dic[ele]
	if ok {
		return val.(int)
	}
	return -1
}

// GetList 获取所有数字
func (h *HashSetInt) GetList() []int {
	return h.Data
}
