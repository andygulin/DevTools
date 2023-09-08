package number

import (
	"encoding/hex"
	"fmt"
	"strconv"
)

type Number interface {
	// DecimalToBinary 10 -> 2
	DecimalToBinary(input int64) (string, error)
	// DecimalToOctal 10 -> 8
	DecimalToOctal(input int64) (string, error)
	// DecimalToHexadecimal 10 -> 16
	DecimalToHexadecimal(input int64) (string, error)

	// BinaryToDecimal 2 -> 10
	BinaryToDecimal(input string) (int64, error)
	// BinaryToOctal 2 -> 8
	BinaryToOctal(input string) (string, error)
	// BinaryToHexadecimal 2 -> 16
	BinaryToHexadecimal(input string) (string, error)

	// OctalToBinary 8 -> 2
	OctalToBinary(input string) (string, error)
	// OctalToDecimal 8 -> 10
	OctalToDecimal(input string) (int64, error)
	// OctalToHexadecimal 8 -> 16
	OctalToHexadecimal(input string) (string, error)

	// HexadecimalToBinary 16 -> 2
	HexadecimalToBinary(input string) (string, error)
	// HexadecimalToOctal 16 -> 8
	HexadecimalToOctal(input string) (string, error)
	// HexadecimalToDecimal 16 -> 10
	HexadecimalToDecimal(input string) (int64, error)
}

type NumberImpl struct{}

func (obj *NumberImpl) DecimalToBinary(input int64) (string, error) {
	return strconv.FormatInt(input, 2), nil
}
func (obj *NumberImpl) DecimalToOctal(input int64) (string, error) {
	return strconv.FormatInt(input, 8), nil
}
func (obj *NumberImpl) DecimalToHexadecimal(input int64) (string, error) {
	return strconv.FormatInt(input, 16), nil
}

func (obj *NumberImpl) BinaryToDecimal(input string) (int64, error) {
	return strconv.ParseInt(input, 2, 64)
}
func (obj *NumberImpl) BinaryToOctal(input string) (string, error) {
	i, _ := strconv.ParseInt(input, 2, 8)
	return strconv.FormatInt(i, 8), nil
}
func (obj *NumberImpl) BinaryToHexadecimal(input string) (string, error) {
	i, _ := strconv.ParseInt(input, 2, 16)
	return strconv.FormatInt(i, 16), nil
}

func (obj *NumberImpl) OctalToBinary(input string) (string, error) {
	i, _ := strconv.ParseInt(input, 8, 64)
	return strconv.FormatInt(i, 2), nil
}
func (obj *NumberImpl) OctalToDecimal(input string) (int64, error) {
	return strconv.ParseInt(input, 8, 64)
}
func (obj *NumberImpl) OctalToHexadecimal(input string) (string, error) {
	i, _ := strconv.ParseInt(input, 8, 64)
	return strconv.FormatInt(i, 16), nil
}

func (obj *NumberImpl) HexadecimalToBinary(input string) (string, error) {
	data, err := hex.DecodeString(input)
	if err != nil {
		return "", err
	}
	binary := ""
	for _, b := range data {
		binary += fmt.Sprintf("%08b", b)
	}
	return binary, nil
}
func (obj *NumberImpl) HexadecimalToOctal(input string) (string, error) {
	octal, err := strconv.ParseInt(input, 16, 64)
	if err != nil {
		return "", err
	}
	return strconv.FormatInt(octal, 8), nil
}
func (obj *NumberImpl) HexadecimalToDecimal(input string) (int64, error) {
	decimal, err := strconv.ParseInt(input, 16, 64)
	if err != nil {
		return 0, err
	}
	num, err := strconv.Atoi(strconv.Itoa(int(decimal)))
	if err != nil {
		return 0, err
	}
	return int64(num), nil
}
