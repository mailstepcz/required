package validate

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

type Person struct {
	Name Required[string] `json:"name"`
	Age  Required[int]    `json:"age"`
}

func TestRequired(t *testing.T) {
	req := require.New(t)

	var p Person
	err := json.Unmarshal([]byte(`{"name":"Saoirse"}`), &p)
	req.Nil(err)
	err = Struct(&p)
	req.NotNil(err)
	req.Equal("field 'Age' in 'validate.Person' is required", err.Error())

	err = json.Unmarshal([]byte(`{"name":"Saoirse","age":25}`), &p)
	req.Nil(err)
	err = Struct(&p)
	req.Nil(err)
}

func TestRequiredPtr(t *testing.T) {
	req := require.New(t)

	var p Person
	err := json.Unmarshal([]byte(`{"name":"Saoirse"}`), &p)
	req.NoError(err)

	req.Equal(p.Name.Ptr().(*string), (*string)(p.Name.UnsafePtr()))
	req.Equal(reflect.ValueOf(p.Name.Ptr()).UnsafePointer(), p.Name.UnsafePtr())
}

var gr interface{}

func BenchmarkWithValidation(b *testing.B) {
	var lr interface{}
	bs := []byte(`{"name":"Saoirse","age":25}`)
	for i := 0; i < b.N; i++ {
		var p Person
		if err := json.Unmarshal(bs, &p); err != nil {
			b.Fatal(err)
		}
		if err := Struct(&p); err != nil {
			b.Fatal(err)
		}
		lr = &p
	}
	gr = lr
}

func BenchmarkWithoutValidation(b *testing.B) {
	var lr interface{}
	bs := []byte(`{"name":"Saoirse","age":25}`)
	for i := 0; i < b.N; i++ {
		var p Person
		if err := json.Unmarshal(bs, &p); err != nil {
			b.Fatal(err)
		}
		lr = &p
	}
	gr = lr
}
