package storage

type MemoryStorage struct {
	data map[string]interface{}
}
func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{data: make(map[string]interface{})}
}

func (s *MemoryStorage) Set(key string, data interface{}) {
	s.data[key] = data
}

func (s MemoryStorage) Get(key string) interface{} {
	if value, ok := s.data[key]; ok {
		return value
	}
	return nil
}
func (s MemoryStorage) GetAll()(map[string]interface{}) {
	return s.data
}

func (s MemoryStorage) Exists(key string) (bool) {
	if _, ok := s.data[key]; !ok {
		return false
	}
	return true
}