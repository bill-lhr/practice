package common

// InArrayForString 字符串是否在数组里
func InArrayForString(s string, list []string) bool {
	for _, v := range list {
		if s == v {
			return true
		}
	}
	return false
}
