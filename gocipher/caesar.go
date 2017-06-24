package gocipher

/*
 * Caesar cipher
 */

// CaesarEncipher enciphers string using Caesar cipher according to key.
func CaesarEncipher(text string, key int) string {
	shift := rune(key)
	runes := []rune(text)
	for i, char := range runes {
		if char >= 'A' && char <= 'Z' {
			runes[i] = modRune(char+shift-'A', 26) + 'A'
		} else if char >= 'a' && char <= 'z' {
			runes[i] = modRune(char+shift-'a', 26) + 'a'
		}
	}
	return string(runes)
}

// CaesarDecipher deciphers string using Caesar cipher according to key.
func CaesarDecipher(text string, key int) string {
	return CaesarEncipher(text, -key)
}
