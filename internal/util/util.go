package util

import (
	"fmt"
	"net/url"
	"reflect"
	"strings"

	"github.com/google/go-querystring/query"
)

type PVEBool bool

func (b *PVEBool) EncodeValues(key string, v *url.Values) error {
	v.Set(key, b.String())
	return nil
}

func (b *PVEBool) UnmarshalJSON(i []byte) error {
	*b = string(i) == "1"
	return nil
}

func (b *PVEBool) String() string {
	if *b {
		return "1"
	}
	return "0"
}

func EncodeArray(key string, v *url.Values, array interface{}) error {
	key = strings.TrimSuffix(key, "[n]")
	switch reflect.TypeOf(array).Kind() {
	case reflect.Ptr:
		return EncodeArray(key, v, reflect.Indirect(reflect.ValueOf(array)).Interface())
	case reflect.Slice:
		s := reflect.ValueOf(array)

		for i := 0; i < s.Len(); i++ {
			derefed := s.Index(i).Interface()
			if derefed != nil {
				if s.Index(i).Type().Kind() == reflect.Ptr {
					if !s.Index(i).Elem().IsNil() {
						derefed = s.Index(i).Elem().Interface()
					} else {
						continue
					}
				}
			}
			elem := struct {
				Item interface{} `url:"item"`
			}{
				Item: derefed,
			}
			d := url.Values{}
			var err error
			if e, ok := elem.Item.(query.Encoder); ok {
				err = e.EncodeValues("item", &d)
			} else {
				d, err = query.Values(elem)
			}

			if err != nil {
				return err
			}
			v.Set(fmt.Sprintf("%s%d", key, i), d.Get("item"))
		}
	default:
		return fmt.Errorf("bad slice type %T", array)
	}
	return nil
}

func EncodeString(key string, v *url.Values, item interface{}, _ string) error {
	newValues, err := query.Values(item)
	if err != nil {
		return err
	}
	values := []string{}
	for key, val := range newValues {
		if len(val) != 1 {
			return fmt.Errorf("Incorrect value count for query")
		}
		values = append(
			values,
			fmt.Sprintf("%s=%s", key, val[0]),
		)
	}
	v.Set(key, strings.Join(values, ","))
	return nil
}
