package number

import "testing"

var obj Number

func init() {
	obj = new(NumberImpl)
}

const (
	binaryValue      = "101011"
	octalValue       = "53"
	decimalValue     = int64(43)
	hexadecimalValue = "2b"
)

func TestNumberImpl_Binary(t *testing.T) {
	input := binaryValue
	if ret, err := obj.BinaryToOctal(input); err == nil {
		t.Log(ret)
	}
	if ret, err := obj.BinaryToDecimal(input); err == nil {
		t.Log(ret)
	}
	if ret, err := obj.BinaryToHexadecimal(input); err == nil {
		t.Log(ret)
	}
}

func TestNumberImpl_Octal(t *testing.T) {
	input := octalValue
	if ret, err := obj.OctalToBinary(input); err == nil {
		t.Log(ret)
	}
	if ret, err := obj.OctalToDecimal(input); err == nil {
		t.Log(ret)
	}
	if ret, err := obj.OctalToHexadecimal(input); err == nil {
		t.Log(ret)
	}
}

func TestNumberImpl_Decimal(t *testing.T) {
	input := decimalValue
	if ret, err := obj.DecimalToBinary(input); err == nil {
		t.Log(ret)
	}
	if ret, err := obj.DecimalToOctal(input); err == nil {
		t.Log(ret)
	}
	if ret, err := obj.DecimalToHexadecimal(input); err == nil {
		t.Log(ret)
	}
}

func TestNumberImpl_Hexadecimal(t *testing.T) {
	input := hexadecimalValue
	if ret, err := obj.HexadecimalToBinary(input); err == nil {
		t.Log(ret)
	}
	if ret, err := obj.HexadecimalToOctal(input); err == nil {
		t.Log(ret)
	}
	if ret, err := obj.HexadecimalToDecimal(input); err == nil {
		t.Log(ret)
	}
}
