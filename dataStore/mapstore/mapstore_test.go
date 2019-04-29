package mapstore

import (
	"../../model"
	"fmt"
	"testing"
)

func TestRange(t *testing.T) {
	model.ShortToLong.Store("fdasfad","453")
	model.ShortToLong.Store("5432543","frdadfsasd")
	var res string
	model.ShortToLong.Range(func(key,value interface{}) bool {
		res += fmt.Sprintf("%s, %s,",key,value)
		return true
	})
	if len(res) == 0 {
		t.Fatal("wrong")
	}
}