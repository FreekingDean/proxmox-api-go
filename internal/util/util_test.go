package util

import (
	"net/url"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncodeArray(t *testing.T) {
	tests := []struct {
		name           string
		expectedValues map[string]string
		val            func() interface{}
	}{
		{
			"slice of strings",
			map[string]string{"somekey0": "1", "somekey1": "2"},
			func() interface{} { return map[string]string{"somekey0": "1", "somekey1": "2"} },
		},
		{
			"ptr of slice of strings",
			map[string]string{"somekey0": "1", "somekey1": "2"},
			func() interface{} {
				arr := map[string]string{"somekey0": "1", "somekey1": "2"}
				return &arr
			},
		},
		{
			"ptr of slice of string ptrs",
			map[string]string{"somekey0": "1", "somekey1": "2"},
			func() interface{} {
				one := "1"
				two := "2"
				arr := map[string]*string{"somekey0": &one, "somekey1": &two}
				return &arr
			},
		},
		{
			"struct",
			map[string]string{"somekey0": "{foo1 1}", "somekey1": "{foo2 2}"},
			func() interface{} {
				return map[string]struct {
					Foo string `url:"foo"`
					Bar int    `url:"bar"`
				}{
					"somekey0": {"foo1", 1},
					"somekey1": {"foo2", 2},
				}
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			v := &url.Values{}
			err := EncodeArray("somekey", v, test.val())
			if assert.NoError(t, err) {
				for k, kv := range test.expectedValues {
					assert.Equal(t, kv, v.Get(k))
				}
			}
		})
	}
}

type Foo struct {
	Bar string `url:"bar"`
	Baz string `url:"baz"`
}

func (f Foo) EncodeValues(k string, v *url.Values) error {
	return nil
}

func TestEncodeString(t *testing.T) {
	v := &url.Values{}
	err := EncodeString("somekey", v, &Foo{"1", "2"}, "foo=<foo>")
	if assert.NoError(t, err) {
		assert.ElementsMatch(t,
			strings.Split("bar=1,baz=2", ","),
			strings.Split(v.Get("somekey"), ","),
		)
	}
}
