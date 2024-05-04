package util

type Entry[K comparable, V any] struct {
	Key   K
	Value V
}

type OrderedMap[K comparable, V any] struct {
	entries  []Entry[K, V]
	index    map[K]int
	defaultK K
	defaultV V
}

func NewOrderedMap[K comparable, V any]() *OrderedMap[K, V] {
	return &OrderedMap[K, V]{
		entries: make([]Entry[K, V], 0),
		index:   make(map[K]int),
	}
}

func (m *OrderedMap[K, V]) Sub(start, end int) *OrderedMap[K, V] {
	if start < 0 {
		start = len(m.entries) + start
	}
	if end < 0 {
		end = len(m.entries) + end
	}
	if start > end {
		return nil
	}
	copy := NewOrderedMap[K, V]()
	for _, entry := range m.entries[start:end] {
		if entry.Key != m.defaultK {
			copy.Put(entry.Key, entry.Value)
		} else {
			copy.PutAnonymous(entry.Value)
		}
	}
	return copy
}

func (m *OrderedMap[K, V]) Index(i int) Entry[K, V] {
	return m.entries[i]
}

func (m *OrderedMap[K, V]) First() Entry[K, V] {
	return m.entries[0]
}

func (m *OrderedMap[K, V]) Last() Entry[K, V] {
	return m.entries[len(m.entries)-1]
}

func (m *OrderedMap[K, V]) Put(key K, value V) bool {
	if _, ok := m.index[key]; ok {
		return false
	}
	m.index[key] = len(m.entries)
	m.entries = append(m.entries, Entry[K, V]{key, value})
	return true
}

func (m *OrderedMap[K, V]) MustPut(key K, value V) {
	if index, ok := m.index[key]; ok {
		m.entries[index].Value = value
	} else {
		m.index[key] = len(m.entries)
		m.entries = append(m.entries, Entry[K, V]{key, value})
	}
}

func (m *OrderedMap[K, V]) Get(key K) (V, bool) {
	if index, ok := m.index[key]; ok {
		return m.entries[index].Value, true
	}
	return m.defaultV, false
}

func (m *OrderedMap[K, V]) MustGet(key K) V {
	if index, ok := m.index[key]; ok {
		return m.entries[index].Value
	}
	return m.defaultV
}

func (m *OrderedMap[K, V]) Has(key K) bool {
	_, ok := m.index[key]
	return ok
}

func (m *OrderedMap[K, V]) Rank(key K) int {
	if index, ok := m.index[key]; ok {
		return index
	}
	return -1
}

func (m *OrderedMap[K, V]) Clear() {
	m.entries = m.entries[:0]
	m.index = make(map[K]int)
}

func (m *OrderedMap[K, V]) Remove(key K) {
	if index, ok := m.index[key]; ok {
		delete(m.index, key)
		m.entries = append(m.entries[:index], m.entries[index+1:]...)
		for i := index; i < len(m.entries); i++ {
			if m.entries[i].Key != m.defaultK {
				m.index[m.entries[i].Key] = i
			}
		}
	}
}

func (m *OrderedMap[K, V]) Len() int {
	return len(m.entries)
}

func (m *OrderedMap[K, V]) Keys() []K {
	keys := make([]K, 0, len(m.entries))
	for _, entry := range m.entries {
		keys = append(keys, entry.Key)
	}
	return keys
}

func (m *OrderedMap[K, V]) Values() []V {
	values := make([]V, 0, len(m.entries))
	for _, entry := range m.entries {
		values = append(values, entry.Value)
	}
	return values
}

func (m *OrderedMap[K, V]) Entries() []Entry[K, V] {
	return m.entries
}

// ==================== Dangerous Functions ====================

// PutAnonymous adds an anonymous entry with default key.
//
// WARNING: added entry which will never be indexed, you can only access it by iteration.
func (m *OrderedMap[K, V]) PutAnonymous(value V) {
	m.entries = append(m.entries, Entry[K, V]{Key: m.defaultK, Value: value})
}
