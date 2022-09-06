package base62

import (
	"errors"
	"math"
	"strings"
)

const (
	base uint64 = 62
	characterSet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

func Encode(number uint64) string {
	var encodeBuilder strings.Builder
	encodeBuilder.Grow(11)

	for ; number > 0; number = number / base {
		encodeBuilder.WriteByte(characterSet[(number % base)])
	}

	return encodeBuilder.String()
}

func Decode(encoded string) (uint64, error) {
	var number uint64

	for i, symbol := range encoded {
		characterSetPosition := strings.IndexRune(characterSet, symbol)

		if characterSetPosition == -1 {
			return uint64(characterSetPosition), errors.New("Invalid character: " + string(symbol))
		}

		number += uint64(characterSetPosition) * uint64(math.Pow(float64(base), float64(i)))
	}

	return number, nil
}
