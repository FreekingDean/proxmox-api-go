package main

var (
	translations = map[string]string{
		"ide":  "IDE",
		"scsi": "SCSI",
		"vmbr": "VMBR",
	}
)

func Tokenize(in string) []string {
	out := ""
	for i, c := range in {
		if i == 0 {
			out += string(UpRune(c))
			continue
		}
		if c == '_' {
			continue
		}
		if i > 0 && in[i-1] == '_' {
			out += string(UpRune(c))
		}
	}
	return []string{out}
}

func UpRune(c rune) rune {
	if c > 'z' || c < 'a' {
		return c
	}
	return c - 'a' + 'A'
}
