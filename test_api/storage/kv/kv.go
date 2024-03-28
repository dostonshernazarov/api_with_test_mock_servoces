package kv

type KV interface {
	Set(key string, value string) error
	Get(key string) (string, error)
	Delete(key string) error
	GetAll() (map[string]string, error)
}

var inst KV

func Init(store KV) {
	inst = store
}

func Set(key string, value string) error {
	return inst.Set(key, value)
}
func Get(key string) (string, error) {
	return inst.Get(key)
}
func Delete(key string) error {
	return inst.Delete(key)
}
func GetAll() (map[string]string, error) {
	return inst.GetAll()
}
