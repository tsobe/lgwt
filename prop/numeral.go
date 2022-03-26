package prop

import (
	"strings"
)

type RomanNumeral struct {
	Value  uint16
	Symbol string
}

type RomanNumerals []RomanNumeral

var allNumerals = RomanNumerals{
	{Value: 1000, Symbol: "M"},
	{Value: 900, Symbol: "CM"},
	{Value: 500, Symbol: "D"},
	{Value: 400, Symbol: "CD"},
	{Value: 100, Symbol: "C"},
	{Value: 90, Symbol: "XC"},
	{Value: 50, Symbol: "L"},
	{Value: 40, Symbol: "XL"},
	{Value: 10, Symbol: "X"},
	{Value: 9, Symbol: "IX"},
	{Value: 5, Symbol: "V"},
	{Value: 4, Symbol: "IV"},
	{Value: 1, Symbol: "I"},
}

type SubtractiveSymbol struct {
	Value    string
	Subtract string
}

func (s SubtractiveSymbol) String() string {
	return s.Subtract + s.Value
}

var subtractiveSymbols = []SubtractiveSymbol{
	{Value: "V", Subtract: "I"},
	{Value: "X", Subtract: "I"},
	{Value: "L", Subtract: "X"},
	{Value: "C", Subtract: "X"},
	{Value: "D", Subtract: "C"},
	{Value: "M", Subtract: "C"},
}

func ConvertToRoman(arabic uint16) string {
	var result strings.Builder
	for _, numeral := range allNumerals {
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
		var value, found = findValueAhead(roman, i)
		result += value
		if !found {
			result += allNumerals.valueOf(string(roman[i]))
		} else {
			i++
		}
	}
	return result
}

func findValueAhead(roman string, index int) (value uint16, found bool) {
	nextIndex := index + 1
	if nextIndex < len(roman) {
		for _, ss := range subtractiveSymbols {
			symbol := string(roman[index])
			nextSymbol := string(roman[nextIndex])
			if ss.Value == nextSymbol && ss.Subtract == symbol {
				return allNumerals.valueOf(ss.String()), true
			}
		}
	}
	return 0, false
}

func (numerals *RomanNumerals) valueOf(symbol string) uint16 {
	for _, numeral := range *numerals {
		if numeral.Symbol == symbol {
			return numeral.Value
		}
	}
	return 0
}
