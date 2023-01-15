package util

type SpecialBool bool

func (b *SpecialBool) UnmarshalJSON(i []byte) error {
	t := SpecialBool(true)
	f := SpecialBool(false)
	if string(i) == "0" {
		b = &f
	} else {
		b = &t
	}
	return nil
}
