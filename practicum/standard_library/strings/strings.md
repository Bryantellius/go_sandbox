# The `strings` Standard Package

The Goland standard `strings`, `strconv`, and `unicode` packages contain many functions for comparing, manipulating, building, mapping, and splitting string values.

## Quick Hits

- `len()` returns the length of characters in a string
- `string()` converts the argument to a string
- `strings.Compare(a, b string)` compares to string values and returns 0 if equal, -1 if a is less than b, and 1 if b is less than a
- `strings.EqualFold(a, b string)` compares equality with case insensitivity
- `strings.ToTitle(s string)` converts the characters in a string to unicode title case
- `strings.ToUpper(s string)` converts the characters in a string to upper case
- `strings.ToLower(s string)` converts the characters in a string to lower case
- `strings.Contains(s string, substr string)` determines if the substring is found in the string
- `strings.ContainsAny(s string, chars string)` determines if any of the characters are found in the string
- `strings.Index(s string, substr string)` returns the starting index of the substring in the string; returns -1 if not found
- `strings.IndexAny(s string, chars string)` returns first index of any given character in the string; returns -1 if not found
- `strings.HasPrefix(s string, pre string)` determines if the strings starts with the substring
- `strings.HasSuffix(s string, suf string)` determines if the strings ends with the substring
- `strings.Count(s string, substr s)` returns the non-overlapping count of the given substring in the string
- `strings.Map(f func(rune)rune, s string)` returns a string with each rune mapped by a given mapping function
- `strings.Builder` is a struct type that has many methods for building and writing string values
- `strconv.Itoa(i int)` converts an integer to a string
- `strconv.Atoi(s string)` converts a string representation of an integer to an integer
- `strconv.ParseFloat(s string, bitsize int)` converts a string representation of a floating point value to a float at the give bitsize
- `unicode.IsLower(char rune)` determines if the rune is a lowercase letter
- `unicode.IsUpper(char rune)` determines if the rune is a lowercase letter
- `unicode.IsPunct(char rune)` determines if the rune is a punctuation
- `unicode.IsDigit(char rune)` determines if the rune is a digit
- `unicode.IsSpace(char rune)` determines if the rune is a space
- `unicode.Is(rangeTab *unicode.RangeTable, r rune)` determines if the rune is found on a given range table
