package util

type SpecialBool bool

func (b *SpecialBool) UnmarshalJSON(i []byte) error {
	*b = string(i) == "1"
	return nil
}
