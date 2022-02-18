package reflection

import (
	"reflect"
	"testing"
)

func TestWalk(t *testing.T) {
	testCases := []struct {
		name             string
		input            interface{}
		withStringFields func(t *testing.T, got []string)
	}{
		{
			name:             "Struct with one string field",
			input:            struct{ Name string }{"Chris"},
			withStringFields: equalTo([]string{"Chris"}),
		},
		{
			name: "Struct with two string fields",
			input: struct {
				Name string
				City string
			}{"Chris", "Berlin"},
			withStringFields: equalTo([]string{"Chris", "Berlin"}),
		},
		{
			name: "Struct with one string and one int fields",
			input: struct {
				Name string
				Age  int
			}{"Chris", 27},
			withStringFields: equalTo([]string{"Chris"}),
		},
		{
			name:             "Nested fields",
			input:            Person{"Chris", Profile{27, "Berlin"}},
			withStringFields: equalTo([]string{"Chris", "Berlin"}),
		},
		{
			name:             "Pointer",
			input:            &Person{"Chris", Profile{27, "Berlin"}},
			withStringFields: equalTo([]string{"Chris", "Berlin"}),
		},
		{
			name: "Slice",
			input: []Profile{
				{27, "Berlin"},
				{28, "London"},
			},
			withStringFields: equalTo([]string{"Berlin", "London"}),
		},
		{
			name: "Array",
			input: [2]Profile{
				{27, "Berlin"},
				{28, "London"},
			},
			withStringFields: equalTo([]string{"Berlin", "London"}),
		},
		{
			name:             "Map",
			input:            map2Of("FooVal", "BarVal"),
			withStringFields: containing([]string{"FooVal", "BarVal"}),
		},
		{
			name:             "Channel",
			input:            channelOf(Profile{22, "Berlin"}, Profile{23, "London"}),
			withStringFields: equalTo([]string{"Berlin", "London"}),
		},
		{
			name:             "Function",
			input:            fn2Of(Profile{22, "Berlin"}, Profile{23, "London"}),
			withStringFields: equalTo([]string{"Berlin", "London"}),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			input := tc.input

			got := collect(input)

			tc.withStringFields(t, got)
		})
	}
}

func equalTo(fields []string) func(*testing.T, []string) {
	return func(t *testing.T, got []string) {
		assertDeepEqual(t, got, fields)
	}
}

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
	input := map[string]string{
		"Foo": val1,
		"Bar": val2,
	}
	return input
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

func collect(in interface{}) []string {
	var collected []string
	walk(in, func(name string) {
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
	contains := false
	for _, val := range haystack {
		if val == needle {
			contains = true
			break
		}
	}
	if !contains {
		t.Errorf("Expected %v to contain %q, but didn't", haystack, needle)
	}
}
