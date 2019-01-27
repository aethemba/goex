package main

import (
	"fmt"
)

func main() {
	var length, delta int
	var input string

	fmt.Scanf("%d\n", &length)
	fmt.Scanf("%s\n", &input)
	fmt.Scanf("%d\n", &delta)

	// alphabetLower := "abcdefghijklmnopqrstuvwxyz"
	// alphabetUpper := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	var result []rune
	for _, ch := range input {
		result = append(result, cipher(ch, delta))
	}
	fmt.Println(string(result))
}

func cipher(r rune, delta int) rune {
	if r >= 'A' && r <= 'Z' {
		return rotate(r, 'A', delta)
	}

	if r >= 'a' && r <= 'z' {
		return rotate(r, 'a', delta)
	}
	return r
}

func rotate(r rune, base, delta int) rune {
	tmp := int(r) - base
	tmp = (tmp + delta) % 26
	return rune(tmp + base)
}

// func rotate(r rune, delta int, key []rune) rune {
// 	idx := strings.IndexRune(string(key), r)

// 	idx = (idx + delta) % len(key)
// 	return key[idx]
// }

// func caesarCipher(s string, k int) string {
// 	var result []string

// 	normalizedK := k
// 	if k > 26 {
// 		normalizedK = int(math.Abs(math.Remainder(float64(26), float64(k))))
// 	}

// 	for _, rune := range s {
// 		if unicode.IsLetter(rune) == true {
// 			var newRune int
// 			if rune >= 97 && rune <= 122 {
// 				newRune = int(rune) + normalizedK
// 				if newRune > 122 {
// 					diff := newRune - 122
// 					newRune = 96 + diff
// 				}
// 			}
// 			if rune >= 65 && rune <= 90 {
// 				newRune = int(rune) + normalizedK
// 				if newRune > 90 {
// 					diff := newRune - 90
// 					newRune = 64 + diff
// 				}
// 			}
// 			result = append(result, string(newRune))
// 		} else {
// 			result = append(result, string(rune))
// 		}
// 	}

// 	return strings.Join(result, "")
// }
