package cache

import "fmt"

// Cache 缓存接口
type Cache interface {
	Set(key string, value interface{})
	Get(key string) interface{}
	Del(key string)
	DelOldest()
	Len() int
}

type Value interface {
	Len() int
}

func CalcLen(value interface{}) int {
	var n int
	switch v := value.(type) {
	case Value:
		n = v.Len()
	case string:
	case bool, uint8, int8:
	case int16, uint16:
	case int32, uint32, float32:
	case int64, uint64,float64:
	case int, uint:
	case complex64:
	case complex128:
	default:
		panic(fmt.Sprintf("%T is not implement cache.Value", value))
	}
	return n
}





