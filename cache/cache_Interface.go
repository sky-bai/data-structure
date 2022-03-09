package cache

type Cache interface {
	Set(key string, value interface{})
	Get(key string) interface{}
	Delete(key string)
	DelOldest()
	Len() int
}

func Calc(value interface{}) int {
	return 1
}
