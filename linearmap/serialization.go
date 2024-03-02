package linearmap

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// ToJSON outputs the JSON representation of the map.
func (m *LinearMap[K, T]) ToJSON() ([]byte, error) {
	elements := make(map[string]interface{})
	for _, it := range m.items {
		key, value := it.key, it.value
		elements[toString(key)] = value
	}
	return json.Marshal(&elements)
}

// FromJSON populates the map from the input JSON representation.
func (m *LinearMap[K, T]) FromJSON(data []byte) error {
	elements := make(map[K]T)
	err := json.Unmarshal(data, &elements)
	if err == nil {
		m.Clear()
		for key, value := range elements {
			m.Put(key, value)
		}
	}
	return err
}

var _ json.Marshaler = &LinearMap[int, int]{}
var _ json.Unmarshaler = &LinearMap[int, int]{}

// UnmarshalJSON @implements json.Unmarshaler
func (m *LinearMap[K, T]) UnmarshalJSON(bytes []byte) error {
	return m.FromJSON(bytes)
}

// MarshalJSON @implements json.Marshaler
func (m *LinearMap[K, T]) MarshalJSON() ([]byte, error) {
	return m.ToJSON()
}

// toString converts a value to string.
func toString(value interface{}) string {
	switch value := value.(type) {
	case string:
		return value
	case int8:
		return strconv.FormatInt(int64(value), 10)
	case int16:
		return strconv.FormatInt(int64(value), 10)
	case int32:
		return strconv.FormatInt(int64(value), 10)
	case int64:
		return strconv.FormatInt(value, 10)
	case uint8:
		return strconv.FormatUint(uint64(value), 10)
	case uint16:
		return strconv.FormatUint(uint64(value), 10)
	case uint32:
		return strconv.FormatUint(uint64(value), 10)
	case uint64:
		return strconv.FormatUint(value, 10)
	case float32:
		return strconv.FormatFloat(float64(value), 'g', -1, 64)
	case float64:
		return strconv.FormatFloat(value, 'g', -1, 64)
	case bool:
		return strconv.FormatBool(value)
	default:
		return fmt.Sprintf("%+v", value)
	}
}
