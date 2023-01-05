package utils

func GetMapValues(m map[string]string) (values []string) {
	// 数组默认长度为map长度,后面append时,不需要重新申请内存和拷贝,效率很高
	for _, v := range m {
		values = append(values, v)
	}

	return
}
