package cache

import "testing"

func TestGetterFunc_Get(t *testing.T) {
	var f Getter = GetterFunc(func(key string) ([]byte, error) {
		return []byte(key),nil
	})
	if v,ok := f.Get("alfjklfj");ok==nil {
		t.Error(v)
	}
}
