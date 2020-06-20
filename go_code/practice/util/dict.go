package util

func dict() {
	// 创建字典
	m := make(map[string]int)

	// 设置kv
	m["hello"] = 1

	// 删除k
	delete(m, "hello")

	//遍历
	for k, v := range m{
		println(k, v)
	}
}
