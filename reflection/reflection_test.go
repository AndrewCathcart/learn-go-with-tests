package reflection

import (
	"reflect"
	"testing"
)

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}

func TestWalk(t *testing.T) {
	t.Run("without maps", func(t *testing.T) {

		cases := []struct {
			Name          string
			Input         interface{}
			ExpectedCalls []string
		}{
			{
				"Struct with one string field",
				struct {
					Name string
				}{"Andy"},
				[]string{"Andy"},
			},
			{
				"Struct with two string field",
				struct {
					Name string
					City string
				}{"Andy", "Newcastle"},
				[]string{"Andy", "Newcastle"},
			},
			{
				"Struct with non-string field",
				struct {
					Name string
					Age  int
				}{"Andy", 25},
				[]string{"Andy"},
			},
			{
				"Nested fields",
				Person{
					"Andy",
					Profile{25, "Newcastle"},
				},
				[]string{"Andy", "Newcastle"},
			},
			{
				"Pointers to things",
				&Person{
					"Andy",
					Profile{25, "Newcastle"},
				},
				[]string{"Andy", "Newcastle"},
			},
			{
				"Slices",
				[]Profile{
					{18, "New York"},
					{19, "Hawaii"},
				},
				[]string{"New York", "Hawaii"},
			},
			{
				"Arrays",
				[2]Profile{
					{33, "London"},
					{34, "Reykjavík"},
				},
				[]string{"London", "Reykjavík"},
			},
		}

		for _, test := range cases {
			t.Run(test.Name, func(t *testing.T) {
				var received []string
				walk(test.Input, func(input string) {
					received = append(received, input)
				})

				if !reflect.DeepEqual(received, test.ExpectedCalls) {
					t.Errorf("received %v, expected %v", received, test.ExpectedCalls)
				}
			})
		}
	})

	t.Run("with maps", func(t *testing.T) {
		aMap := map[string]string{
			"Foo": "Bar",
			"Baz": "Boz",
		}

		var received []string
		walk(aMap, func(input string) {
			received = append(received, input)
		})

		assertContains(t, received, "Bar")
		assertContains(t, received, "Boz")
	})

	t.Run("with channels", func(t *testing.T) {
		aChannel := make(chan Profile)

		go func() {
			aChannel <- Profile{33, "Berlin"}
			aChannel <- Profile{34, "Katowice"}
			close(aChannel)
		}()

		var received []string
		expected := []string{"Berlin", "Katowice"}

		walk(aChannel, func(input string) {
			received = append(received, input)
		})

		if !reflect.DeepEqual(received, expected) {
			t.Errorf("received %v, expected %v", received, expected)
		}
	})

	t.Run("with function", func(t *testing.T) {
		aFunction := func() (Profile, Profile) {
			return Profile{33, "Berlin"}, Profile{34, "Katowice"}
		}

		var received []string
		expected := []string{"Berlin", "Katowice"}

		walk(aFunction, func(input string) {
			received = append(received, input)
		})

		if !reflect.DeepEqual(received, expected) {
			t.Errorf("received %v, expected %v", received, expected)
		}
	})
}

func assertContains(t *testing.T, haystack []string, needle string) {
	t.Helper()
	contains := false
	for _, x := range haystack {
		if x == needle {
			contains = true
		}
	}
	if !contains {
		t.Errorf("expected %+v to contain %q but it didn't", haystack, needle)
	}
}
