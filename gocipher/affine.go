package gocipher

/*
 * Affine cipher
 */

import "errors"

type AffineKey struct {
	a, b, inva int
}

// NewAffineKey creates an AffineKey.
// `a` is the multiplicative part of the key (allowable values are: 1, 3, 5, 7, 9, 11, 15, 17, 19, 21, 23, and 25).
// `b` is the additive part of the key (allowable values are integers 0-25).
func NewAffineKey(a, b int) (*AffineKey, error) {
	inva := -1
	for i := 1; i < 26; i += 2 {
		if mod(a*i, 26) == 1 {
			inva = i
		}
	}
	if 0 > inva || inva > 25 {
		return nil, errors.New("invalid key: a=" + string(a) + ", no inverse exists (mod 26)")
	}
	return &AffineKey{a, b, inva}, nil
}

// AffineEncipher enciphers string using Affine cipher according to key.
func AffineEncipher(text string, key *AffineKey) string {
	runes := []rune(text)
	for i, char := range runes {
		if isAlpha, isUpper := isAlpha(char); isAlpha {
			runes[i] = i2a(key.a*a2i(char)+key.b, isUpper)
		}
	}
	return string(runes)
}

// AffineDecipher deciphers string using Affine cipher according to key.
func AffineDecipher(text string, key *AffineKey) string {
	runes := []rune(text)
	for i, char := range runes {
		if isAlpha, isUpper := isAlpha(char); isAlpha {
			runes[i] = i2a(key.inva*(a2i(char)-key.b), isUpper)
		}
	}
	return string(runes)
}
