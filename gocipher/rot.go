package gocipher

import (
	"strings"
	"unicode"
)

/*
 * ROT-5 cipher
 * ROT-13 cipher
 * ROT-18 cipher
 * ROT-47 cipher
 */

// Rot5Encipher enciphers string using ROT-5 cipher. Identical to Rot5Decipher.
// e.g. `1234567890` becomes `5678901234`.
func Rot5Encipher(text string) string {
	return rotRange(text, 5, '0', '9')
}

// Rot5Decipher deciphers string using ROT-5 cipher. Identical to Rot5Encipher.
// e.g. `5678901234` becomes `1234567890`.
func Rot5Decipher(text string) string {
	return Rot5Encipher(text)
}

// Rot13Encipher enciphers string using ROT-13 cipher. Identical to Rot13Decipher.
// e.g. `ABCDEFGHIJKLM` becomes `NOPQRSTUVWXYZ`.
func Rot13Encipher(text string) string {
	return CaesarEncipher(text, 13)
}

// Rot13Decipher deciphers string using ROT-13 cipher. Identical to Rot13Encipher.
// e.g. `NOPQRSTUVWXYZ` becomes `ABCDEFGHIJKLM`.
func Rot13Decipher(text string) string {
	return Rot13Encipher(text)
}

// Rot18Encipher enciphers string using ROT-18 cipher. Identical to Rot18Decipher.
// Note: numbers become lower case.
// e.g. `ABCXYZ012` becomes `STUFGHijk`.
func Rot18Encipher(text string) string {
	return rotAlphabetCaps(text, 18, "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
}

// Rot18Decipher deciphers string using ROT-18 cipher. Identical to Rot18Encipher.
// Note: numbers become lower case.
// e.g. `STUFGHIJK` becomes `ABCXYZ012`.
func Rot18Decipher(text string) string {
	return Rot18Encipher(text)
}

// Rot47Encipher enciphers string using ROT-47 cipher. Identical to Rot47Decipher.
// e.g. `ABCabc` becomes `pqr234`.
func Rot47Encipher(text string) string {
	return rotRange(text, 47, '!', '~')
}

// Rot47Decipher deciphers string using ROT-47 cipher. Identical to Rot47Encipher.
// e.g. `pqr234` becomes `ABCabc`.
func Rot47Decipher(text string) string {
	return Rot47Encipher(text)
}

func rotRange(text string, key int, min, max rune) string {
	size := max - min + 1
	shift := rune(key)
	runes := []rune(text)
	for i, char := range runes {
		if char >= min && char <= max {
			runes[i] = modRune(char+shift-min, size) + min
		}
	}
	return string(runes)
}

func rotAlphabet(text string, key int, alphabet string) string {
	size := len(alphabet)
	alphaRunes := []rune(alphabet)
	runes := []rune(text)
	for i, char := range runes {
		if pos := strings.IndexRune(alphabet, char); pos != -1 {
			runes[i] = alphaRunes[mod(pos+key, size)]
		}
	}
	return string(runes)
}

// Same as rotAlphabet, but preserves capitaization
func rotAlphabetCaps(text string, key int, alphabet string) string {
	size := len(alphabet)
	alphabet = strings.ToLower(alphabet)
	alphaRunes := []rune(alphabet)
	runes := []rune(text)
	for i, char := range runes {
		charLower := unicode.ToLower(char)
		if pos := strings.IndexRune(alphabet, charLower); pos != -1 {
			shifted := alphaRunes[mod(pos+key, size)]
			if unicode.IsUpper(char) {
				shifted = unicode.ToUpper(shifted)
			}
			runes[i] = shifted
		}
	}
	return string(runes)
}
