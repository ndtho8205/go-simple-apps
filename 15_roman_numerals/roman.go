package roman

import (
	"strings"
)

type Numeral struct {
	Value  uint16
	Symbol string
}

type Numerals []Numeral

func (n Numerals) ValueOf(symbol string) uint16 {
	for _, roman := range n {
		if roman.Symbol == symbol {
			return roman.Value
		}
	}

	return 0
}

var allRomanNumerals = Numerals{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"}, {1, "I"},
}

func ConvertToRoman(arabic uint16) string {
	var result strings.Builder

	for _, numeral := range allRomanNumerals {
		for arabic >= numeral.Value {
			result.WriteString(numeral.Symbol)
			arabic -= numeral.Value
		}
	}

	return result.String()
}

func ConvertToArabic(roman string) uint16 {
	var result uint16

	for i := 0; i < len(roman); i++ {
		currentRomanValue := allRomanNumerals.ValueOf(string(roman[i]))

		if i+1 < len(roman) {
			nextRomanValue := allRomanNumerals.ValueOf(string(roman[i+1]))

			if currentRomanValue < nextRomanValue {
				currentRomanValue = -currentRomanValue
			}
		}

		result += currentRomanValue
	}

	return result
}
