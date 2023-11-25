package linearmap

import "fmt"

type item[K, V any] struct {
	key   K
	value V
}

type LinearMap[K comparable, V any] struct {
	items []item[K, V]
}

func New[K comparable, V any]() *LinearMap[K, V] {
	return &LinearMap[K, V]{}
}

func (m *LinearMap[K, V]) Put(key K, value V) {
	for i, it := range m.items {
		if it.key == key {
			m.items[i] = item[K, V]{key, value}
			return
		}
	}
	m.items = append(m.items, item[K, V]{key, value})
}

func (m *LinearMap[K, V]) Get(key K) (value V, found bool) {
	for _, item := range m.items {
		if item.key == key {
			return item.value, true
		}
	}
	var zero V
	return zero, false
}

func remove[T any](slice []T, index int) []T {
	slice[index] = slice[len(slice)-1]
	return slice[:len(slice)-1]
}

func (m *LinearMap[K, V]) Remove(key K) {
	for i, it := range m.items {
		if it.key == key {
			m.items = remove(m.items, i)
			return
		}
	}
}

func (m *LinearMap[K, V]) Keys() []K {
	keys := make([]K, m.Size())
	for i, item := range m.items {
		keys[i] = item.key
	}
	return keys
}

func (m *LinearMap[K, V]) Empty() bool {
	return m.Size() == 0
}

func (m *LinearMap[K, V]) Size() int {
	return len(m.items)
}

func (m *LinearMap[K, V]) Clear() {
	m.items = nil
}

func (m *LinearMap[K, V]) Values() []V {
	values := make([]V, m.Size())
	for i, item := range m.items {
		values[i] = item.value
	}
	return values
}

func (m *LinearMap[K, V]) String() string {
	return fmt.Sprintf("LinearMap{%v}", m.items)
}
