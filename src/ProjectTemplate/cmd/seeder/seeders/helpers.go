package seeders

import (
	"github.com/best-expendables/common-utils/util/uuid_generator"
	"math/rand"
	"reflect"
	"time"
)

func randDate() time.Time {
	min := time.Date(1970, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	max := time.Date(2010, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	delta := max - min

	sec := rand.Int63n(delta) + min
	return time.Unix(sec, 0)
}

func setIdIfNotExists(model interface{}) {
	ps := reflect.ValueOf(model)
	if ps.Kind() != reflect.Ptr {
		return
	}
	e := ps.Elem()
	if e.Kind() != reflect.Struct {
		return
	}
	f := e.FieldByName("ID")
	if !f.IsValid() {
		return
	}
	if f.String() == "" && f.CanSet() {
		f.SetString(uuid_generator.NewUUIDV4())
	}
}
