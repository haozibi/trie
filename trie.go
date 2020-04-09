package trie

type Trie interface {
	Get(key string) interface{}
	Put(key string, value interface{}) bool
	Delete(key string) bool
	// Pretty()
}
