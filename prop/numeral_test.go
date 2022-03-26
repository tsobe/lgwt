package prop

import (
	"fmt"
	"testing"
	"testing/quick"
)

var cases = []struct {
	arabic uint16
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
			got := ConvertToRoman(test.arabic)
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

func TestPropertiesOfConversion(t *testing.T) {
	assertion := func(arabic uint16) bool {
		if arabic > 3999 {
			return true
		}
		t.Log("testing", arabic)
		roman := ConvertToRoman(arabic)
		fromRoman := ConvertToArabic(roman)
		return arabic == fromRoman
	}

	if err := quick.Check(assertion, nil); err != nil {
		t.Error("Checks failed", err)
	}
}
