package golang

type Hash interface {
	Hash(hashVal interface{}) uint64
}

type Fnv1aHash struct{}

// Implement Hash interface using Fnv1a algorithm
func (f *Fnv1aHash) Hash(hashVal interface{}) uint64 {
	a := hashVal.(string)

	var hash uint64 = 14695981039346656037
	var prime uint64 = 1099511628211

	for _, c := range a {
		hash ^= uint64(c)
		hash *= prime
	}

	return hash
}

type HashMap struct {
	mapStore map[uint64]interface{}
	hasher   Hash
}

func NewHashMap(hasher Hash) *HashMap {
	var h Hash
	if hasher == nil {
		h = &Fnv1aHash{}
	} else {
		h = hasher
	}

	return &HashMap{
		mapStore: make(map[uint64]interface{}),
		hasher:   h,
	}
}

func (h *HashMap) Put(key interface{}, value interface{}) {
	hash := h.hasher.Hash(key)
	h.mapStore[hash] = value
}

func (h *HashMap) Get(key interface{}) interface{} {
	hash := h.hasher.Hash(key)
	return h.mapStore[hash]
}
