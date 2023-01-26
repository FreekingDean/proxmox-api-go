package util

import "net/url"

type SpecialBool bool

func (b *SpecialBool) EncodeValues(key string, v *url.Values) error {
	v.Set(key, b.String())
	return nil
}

func (b *SpecialBool) UnmarshalJSON(i []byte) error {
	*b = string(i) == "1"
	return nil
}

func (b *SpecialBool) String() string {
	if *b {
		return "1"
	}
	return "0"
}
