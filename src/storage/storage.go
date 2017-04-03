package storage

type Storage interface {
	Set(key string, data interface{})
	Get(key string) interface{}
	GetAll() map[string]interface{}
	Exists(key string) bool
}
