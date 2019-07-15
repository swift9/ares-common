package converters

import (
	"encoding/json"
	"reflect"
)

func Copy(src interface{}, dst interface{}) error {
	if reflect.TypeOf(src) == reflect.TypeOf(dst) {
		dst = &src
		return nil
	}
	bytes, err := json.Marshal(src)
	if err != nil {
		return err
	}
	json.Unmarshal(bytes, &dst)
	return nil
}
