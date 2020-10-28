package main

import "fmt"

func main() {
	length, delta := 0, 0
	s := ""

	fmt.Scanf("%d\n", &length)
	fmt.Scanf("%s\n", &s)
	fmt.Scanf("%d\n", &delta)

	c := make([]rune, length)

	for _, char := range s {
		c = append(c, caeserCipher(char, delta))
	}

	fmt.Println(string(c))
}

func caeserCipher(r rune, d int) rune {
	if r >= 'A' && r <= 'Z' {
		return rotateWithBaseRune(r, 'A', d)
	}
	if r >= 'a' && r <= 'z' {
		return rotateWithBaseRune(r, 'a', d)
	}
	return r
}

func rotateWithBaseRune(r rune, base, d int) rune {
	tmp := int(r) - base
	tmp = (tmp + d) % 26
	return rune(tmp + base)
}
