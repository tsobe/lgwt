package reflection

import (
	"reflect"
	"testing"
)

func TestWalk(t *testing.T) {
	type Profile struct {
		Age  int
		City string
	}

	type Person struct {
		Name    string
		Profile Profile
	}

	testCases := []struct {
		name            string
		input           interface{}
		hasStringFields []string
	}{
		{
			name:            "Struct with one string field",
			input:           struct{ Name string }{"Chris"},
			hasStringFields: []string{"Chris"},
		},
		{
			name: "Struct with two string fields",
			input: struct {
				Name string
				City string
			}{"Chris", "Berlin"},
			hasStringFields: []string{"Chris", "Berlin"},
		},
		{
			name: "Struct with one string and one int fields",
			input: struct {
				Name string
				Age  int
			}{"Chris", 27},
			hasStringFields: []string{"Chris"},
		},
		{
			name:            "Nested fields",
			input:           Person{"Chris", Profile{27, "Berlin"}},
			hasStringFields: []string{"Chris", "Berlin"},
		},
		{
			name:            "Pointer",
			input:           &Person{"Chris", Profile{27, "Berlin"}},
			hasStringFields: []string{"Chris", "Berlin"},
		},
		{
			name: "Slice",
			input: []Profile{
				{27, "Berlin"},
				{28, "London"},
			},
			hasStringFields: []string{"Berlin", "London"},
		},
		{
			name: "Array",
			input: [2]Profile{
				{27, "Berlin"},
				{28, "London"},
			},
			hasStringFields: []string{"Berlin", "London"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			input := tc.input
			want := tc.hasStringFields

			got := collect(input)

			if !reflect.DeepEqual(got, want) {
				t.Errorf("Got %v, expected %v", got, want)
			}
		})
	}

	t.Run("Map", func(t *testing.T) {
		input := map[string]string{
			"Foo": "FooVal",
			"Bar": "BarVal",
		}

		got := collect(input)

		assertContains(t, got, "FooVal")
		assertContains(t, got, "BarVal")
	})

	t.Run("Channel", func(t *testing.T) {
		input := make(chan Profile)
		want := []string{"Berlin", "London"}
		go func() {
			input <- Profile{22, "Berlin"}
			input <- Profile{23, "London"}
			close(input)
		}()

		got := collect(input)

		assertDeepEqual(t, got, want)
	})

	t.Run("Function", func(t *testing.T) {
		input := func() (Profile, Profile) {
			return Profile{22, "Berlin"}, Profile{23, "London"}
		}
		want := []string{"Berlin", "London"}

		got := collect(input)

		assertDeepEqual(t, got, want)
	})
}

func collect(in interface{}) []string {
	var collected []string
	walk(in, func(name string) {
		collected = append(collected, name)
	})
	return collected
}

func assertDeepEqual(t *testing.T, got, want interface{}) {
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