package piscine

import "github.com/01-edu/z01"

func PrintStr(s string) {
	casted := []rune(s)
	for _, i := range casted {
		z01.PrintRune(i)
	}

}
