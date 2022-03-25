package prop

import (
	"fmt"
	"testing"
)

var cases = []struct {
	arabic int
	roman  string
}{
	{arabic: 1, roman: "I"},
	{arabic: 2, roman: "II"},
	{arabic: 3, roman: "III"},
	{arabic: 4, roman: "IV"},
	{arabic: 5, roman: "V"},
	{arabic: 6, roman: "VI"},
	{arabic: 7, roman: "VII"},
	{arabic: 8, roman: "VIII"},
	{arabic: 9, roman: "IX"},
	{arabic: 10, roman: "X"},
	{arabic: 11, roman: "XI"},
	{arabic: 14, roman: "XIV"},
	{arabic: 18, roman: "XVIII"},
	{arabic: 20, roman: "XX"},
	{arabic: 29, roman: "XXIX"},
	{arabic: 39, roman: "XXXIX"},
	{arabic: 40, roman: "XL"},
	{arabic: 48, roman: "XLVIII"},
	{arabic: 50, roman: "L"},
	{arabic: 100, roman: "C"},
	{arabic: 90, roman: "XC"},
	{arabic: 90, roman: "XC"},
	{arabic: 400, roman: "CD"},
	{arabic: 500, roman: "D"},
	{arabic: 900, roman: "CM"},
	{arabic: 1000, roman: "M"},
	{arabic: 1984, roman: "MCMLXXXIV"},
	{arabic: 3999, roman: "MMMCMXCIX"},
	{arabic: 2014, roman: "MMXIV"},
	{arabic: 1006, roman: "MVI"},
	{arabic: 798, roman: "DCCXCVIII"},
}

func TestArabicToRoman(t *testing.T) {
	for _, test := range cases {
		t.Run(fmt.Sprintf("Subtract %d converts to Value %q", test.arabic, test.roman), func(t *testing.T) {
			got := ConvertToNumeral(test.arabic)
			want := test.roman

			if got != want {
				t.Errorf("Expected %q, got %q", want, got)
			}
		})
	}
}

func TestRomanToArabic(t *testing.T) {
	for _, test := range cases {
		t.Run(fmt.Sprintf("Value %q converts to Subtract %d", test.roman, test.arabic), func(t *testing.T) {
			got := ConvertToArabic(test.roman)
			want := test.arabic

			if got != want {
				t.Errorf("Expected %d, got %d", want, got)
			}
		})
	}
}
