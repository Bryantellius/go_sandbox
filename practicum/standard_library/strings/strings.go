package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	s := "The quick brown fox jumps over the lazy dog"

	// length of a string
	fmt.Println("Length of s:", len(s))

	// iterate of string characters
	for _, ch := range s {
		fmt.Print(string(ch), ",")
	}

	// comparing strings
	result := strings.Compare("dog", "cat")
	fmt.Println(result)
	result = strings.Compare("dog", "dog")
	fmt.Println(result)

	// comparing with unicode equal folding
	fmt.Println(strings.EqualFold("Dog", "dog"))

	// alter casing
	fmt.Println(strings.ToTitle("dog and cat"))
	fmt.Println(strings.ToUpper("dog and cat"))
	fmt.Println(strings.ToLower("DOG AND CAT"))

	// string searching
	fmt.Println(strings.Contains(s, "fox"))
	fmt.Println(strings.ContainsAny(s, "abcd"))
	fmt.Println(strings.Index(s, "fox"))
	fmt.Println(strings.Index(s, "cat"))
	fmt.Println(strings.IndexAny("rhythm", "aeiouAEIOU"))
	fmt.Println(strings.HasPrefix(s, "The"))
	fmt.Println(strings.HasSuffix("google.com", ".com"))
	fmt.Println(strings.Count(s, "o"))

	// string manipulation
	words := strings.Split(s, " ")
	fmt.Printf("%q\n", words)

	joined := strings.Join([]string{"My", "name", "is", "Frodo"}, " ")
	fmt.Printf("%s\n", joined)

	fields := strings.Fields(s)
	fmt.Printf("%q\n", fields)

	replacer := strings.NewReplacer("fox", "---", "dog", "---")
	replaced := replacer.Replace(s)
	fmt.Println(replaced)

	// string mapping
	shift := 11
	cipher := CaesarCipher(s, shift)
	fmt.Println(cipher)
	fmt.Println(CaesarCipher(cipher, -shift))

	// string builder
	var builder strings.Builder

	builder.WriteString("My\n")
	builder.WriteString("name\n")
	builder.WriteString("is\n")
	builder.WriteString("Ben\n")

	fmt.Println(builder.String())

	// string builder's capacity
	fmt.Println("Builder's capacity:", builder.Cap())
	// growing builder's capacity
	builder.Grow(1024)
	fmt.Println("Builder's capacity:", builder.Cap())

	// builder's size
	fmt.Println("Builder's size:", builder.Len())

	// reseting a builder
	builder.Reset()

	// converting strings
	num := 122
	str := "250"
	fmt.Println(strconv.Itoa(num))
	fmt.Println(strconv.Atoi(str))
	fmt.Println(strconv.ParseFloat("3.141589", 64))

	// formatting strings
	fmt.Println(strconv.FormatBool(true))
	fmt.Println(strconv.FormatFloat(3.14159, 'E', -1, 64))
	fmt.Println(strconv.FormatInt(-42, 10))

	// tests with unicode
	fmt.Println(unicode.IsLower('A'))
	fmt.Println(unicode.IsPunct('!'))
	fmt.Println(unicode.IsDigit('0'))
	fmt.Println(unicode.IsSpace(' '))
	fmt.Println(unicode.Is(unicode.Hex_Digit, 'a'))
}

// Implements the Caesar cipher.
// Pass a positive number for shift to encrypt.
// Pass a negative number for shift to decrypt.
func CaesarCipher(s string, shift int) string {
	transform := func(r rune) rune {
		var value int
		switch {
		case r >= 'A' && r <= 'Z':
			value = int('A') + (int(r) - int('A') + shift)
			if value > 91 {
				value -= 26
			} else if value < 65 {
				value += 26
			}
		case r >= 'a' && r <= 'z':
			value = int('a') + (int(r) - int('a') + shift)
			if value > 122 {
				value -= 26
			} else if value < 97 {
				value += 26
			}
		default:
			value = int(r)
		}

		return rune(value)
	}

	return strings.Map(transform, s)
}
