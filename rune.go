package trie

type runeTrie struct {
	value interface{}
	chile map[rune]*runeTrie
}

// Trie Tree, 只读是线程安全的，同时读写是线程不安全
func NewRuneTrie() Trie {
	return &runeTrie{}
}

func (r *runeTrie) Get(key string) interface{} {

	node := r
	for _, v := range key {
		node = node.chile[v]
		if node == nil {
			return nil
		}
	}

	return node.value
}

func (r *runeTrie) Put(key string, value interface{}) bool {

	node := r
	for _, v := range key {
		child := node.chile[v]
		if child == nil {
			// 初始化未初始化节点
			if node.chile == nil {
				node.chile = make(map[rune]*runeTrie)
			}
			child = new(runeTrie)
			node.chile[v] = child
		}
		node = child
	}
	isNewVal := node.value == nil
	node.value = value
	return isNewVal
}

// Delete 删除 key，如果是叶叶节点，需要层层删除，保证清洁
func (r *runeTrie) Delete(key string) bool {
	// 记录父节点
	path := make([]nodeRune, len([]rune(key)))
	node := r
	for k, v := range key {
		path[k] = nodeRune{
			node: node,
			r:    v,
		}
		node = node.chile[v]
		if node == nil {
			// not exist
			return false
		}
	}

	node.value = nil
	// 如果是叶节点，则从低向上查找，删除无效的边
	if node.isLeaf() {
		for i := len(path) - 1; i >= 0; i-- {
			parent := path[i].node
			r := path[i].r
			delete(parent.chile, r)
			// 如果父节点还有其他子节点则放弃
			if !parent.isLeaf() {
				break
			}
			parent.chile = nil
			// 当前节点也是有效的值
			if parent.value != nil {
				break
			}
		}
	}

	return true
}

func (r *runeTrie) isLeaf() bool {
	return len(r.chile) == 0
}

type nodeRune struct {
	node *runeTrie
	r    rune
}
