package reflection

import (
	"reflect"
	"testing"
)

func TestWalk(t *testing.T) {
	testCases := []struct {
		name                   string
		walking                interface{}
		yieldsWithStringFields func(t *testing.T, got []string)
	}{
		{
			name:                   "Struct with one string field",
			walking:                struct{ Name string }{"Chris"},
			yieldsWithStringFields: equalTo([]string{"Chris"}),
		},
		{
			name: "Struct with two string fields",
			walking: struct {
				Name string
				City string
			}{"Chris", "Berlin"},
			yieldsWithStringFields: equalTo([]string{"Chris", "Berlin"}),
		},
		{
			name: "Struct with one string and one int fields",
			walking: struct {
				Name string
				Age  int
			}{"Chris", 27},
			yieldsWithStringFields: equalTo([]string{"Chris"}),
		},
		{
			name:                   "Nested fields",
			walking:                Person{"Chris", Profile{27, "Berlin"}},
			yieldsWithStringFields: equalTo([]string{"Chris", "Berlin"}),
		},
		{
			name:                   "Pointer",
			walking:                &Person{"Chris", Profile{27, "Berlin"}},
			yieldsWithStringFields: equalTo([]string{"Chris", "Berlin"}),
		},
		{
			name: "Slice",
			walking: []Profile{
				{27, "Berlin"},
				{28, "London"},
			},
			yieldsWithStringFields: equalTo([]string{"Berlin", "London"}),
		},
		{
			name: "Array",
			walking: [2]Profile{
				{27, "Berlin"},
				{28, "London"},
			},
			yieldsWithStringFields: equalTo([]string{"Berlin", "London"}),
		},
		{
			name:                   "Map",
			walking:                map2Of("FooVal", "BarVal"),
			yieldsWithStringFields: containing([]string{"BarVal", "FooVal"}),
		},
		{
			name: "Channel",
			walking: channelOf(
				Profile{22, "Berlin"},
				Profile{23, "London"},
			),
			yieldsWithStringFields: equalTo([]string{"Berlin", "London"}),
		},
		{
			name:                   "Function",
			walking:                fn2Of(Profile{22, "Berlin"}, Profile{23, "London"}),
			yieldsWithStringFields: equalTo([]string{"Berlin", "London"}),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			input := tc.walking

			got := walk(input)

			tc.yieldsWithStringFields(t, got)
		})
	}
}

// returns function asserting that actual fields equal to the specified fields (exact, ordered match)
func equalTo(fields []string) func(*testing.T, []string) {
	return func(t *testing.T, got []string) {
		assertDeepEqual(t, got, fields)
	}
}

// returns function asserting that actual fields contain all the specified fields (in no particular order)
func containing(fields []string) func(*testing.T, []string) {
	return func(t *testing.T, got []string) {
		for _, field := range fields {
			assertContains(t, got, field)
		}
	}
}

type Profile struct {
	Age  int
	City string
}

type Person struct {
	Name    string
	Profile Profile
}

func map2Of(val1, val2 string) map[string]string {
	return map[string]string{
		"Foo": val1,
		"Bar": val2,
	}
}

func fn2Of(p1, p2 Profile) func() (Profile, Profile) {
	return func() (Profile, Profile) {
		return p1, p2
	}
}

func channelOf(profiles ...Profile) chan Profile {
	ch := make(chan Profile)
	go func() {
		for _, p := range profiles {
			ch <- p
		}
		close(ch)
	}()
	return ch
}

func walk(in interface{}) []string {
	var collected []string
	Walk(in, func(name string) {
		collected = append(collected, name)
	})
	return collected
}

func assertDeepEqual(t *testing.T, got, want interface{}) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Got %v, expected %v", got, want)
	}
}

func assertContains(t *testing.T, haystack []string, needle string) {
	t.Helper()
	for _, val := range haystack {
		if val == needle {
			return
		}
	}
	t.Errorf("Expected %v to contain %q", haystack, needle)
}
