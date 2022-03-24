package prop

import (
	"fmt"
	"strings"
	"testing"
)

func TestRomanNumerals(t *testing.T) {
	tests := []struct {
		arabic            int
		convertsToRomanAs string
	}{
		{arabic: 1, convertsToRomanAs: "I"},
		{arabic: 2, convertsToRomanAs: "II"},
		{arabic: 3, convertsToRomanAs: "III"},
		{arabic: 4, convertsToRomanAs: "IV"},
		{arabic: 5, convertsToRomanAs: "V"},
		{arabic: 6, convertsToRomanAs: "VI"},
		{arabic: 7, convertsToRomanAs: "VII"},
		{arabic: 8, convertsToRomanAs: "VIII"},
		{arabic: 9, convertsToRomanAs: "IX"},
		{arabic: 10, convertsToRomanAs: "X"},
		{arabic: 11, convertsToRomanAs: "XI"},
		{arabic: 14, convertsToRomanAs: "XIV"},
		{arabic: 18, convertsToRomanAs: "XVIII"},
		{arabic: 20, convertsToRomanAs: "XX"},
		{arabic: 29, convertsToRomanAs: "XXIX"},
		{arabic: 39, convertsToRomanAs: "XXXIX"},
		{arabic: 40, convertsToRomanAs: "XL"},
		{arabic: 48, convertsToRomanAs: "XLVIII"},
		{arabic: 50, convertsToRomanAs: "L"},
		{arabic: 100, convertsToRomanAs: "C"},
		{arabic: 90, convertsToRomanAs: "XC"},
		{arabic: 90, convertsToRomanAs: "XC"},
		{arabic: 400, convertsToRomanAs: "CD"},
		{arabic: 500, convertsToRomanAs: "D"},
		{arabic: 900, convertsToRomanAs: "CM"},
		{arabic: 1000, convertsToRomanAs: "M"},
		{arabic: 1984, convertsToRomanAs: "MCMLXXXIV"},
		{arabic: 3999, convertsToRomanAs: "MMMCMXCIX"},
		{arabic: 2014, convertsToRomanAs: "MMXIV"},
		{arabic: 1006, convertsToRomanAs: "MVI"},
		{arabic: 798, convertsToRomanAs: "DCCXCVIII"},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("Value %d converts to Symbol %q", test.arabic, test.convertsToRomanAs), func(t *testing.T) {
			got := ConvertToNumeral(test.arabic)
			want := test.convertsToRomanAs

			if got != want {
				t.Errorf("Expected %q, got %q", want, got)
			}
		})
	}
}

type RomanNumeral struct {
	Value  int
	Symbol string
}

func ConvertToNumeral(arabic int) string {
	var result strings.Builder

	allNumerals := []RomanNumeral{
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

	for _, numeral := range allNumerals {
		for arabic >= numeral.Value {
			result.WriteString(numeral.Symbol)
			arabic -= numeral.Value
		}
	}

	return result.String()
}
