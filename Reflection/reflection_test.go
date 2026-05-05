package reflection

import (
	"slices"
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
	cases := []struct {
		Name          string
		Input         any
		ExpectedCalls []string
	}{
		{
			"struct with one string field",
			struct{ Name string }{"Chris"},
			[]string{"Chris"},
		},
		{
			"struct with two string field",
			struct {
				Name string
				City string
			}{"Chris", "London"},
			[]string{"Chris", "London"},
		},
		{
			"struct with non string field",
			struct {
				Name string
				Age  int
			}{"Chris", 10},
			[]string{"Chris"},
		},
		{
			"struct with nested fields",
			Person{
				"Chris",
				Profile{30, "London"},
			},
			[]string{"Chris", "London"},
		},
		{
			"pointer to things",
			&Person{
				"Chris",
				Profile{30, "London"},
			},
			[]string{"Chris", "London"},
		},
		{
			"slices",
			[]Profile{
				{33, "London"},
				{23, "Delhi"},
			},
			[]string{"London", "Delhi"},
		},
		{
			"array of values",
			[2]Profile{
				{33, "London"},
				{23, "Delhi"},
			},
			[]string{"London", "Delhi"},
		},
		{
			"map of values",
			map[string]string{
				"Cow": "Moo", "Sheep": "Baa",
			},
			[]string{"Moo", "Baa"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk(test.Input, func(input string) {
				got = append(got, input)
			})

			assertContains(t, got, test.ExpectedCalls)
		})
	}
	t.Run("channels", func(t *testing.T) {
		aChannel := make(chan Profile)
		go func() {
			aChannel <- Profile{33, "London"}
			aChannel <- Profile{23, "Delhi"}
			close(aChannel)
		}()
		var got []string
		walk(aChannel, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, []string{"London", "Delhi"})
	})
	t.Run("with function", func(t *testing.T) {
		aFunction := func() (Profile, Profile){
			return Profile{33, "London"}, Profile{23, "Delhi"}
		}
		var got []string
		walk(aFunction, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, []string{"London", "Delhi"})
	})
}

func assertContains(t testing.TB, haystack []string, needles []string) {
	t.Helper()

	for _, needle := range needles {
		if !slices.Contains(haystack, needle) {
			t.Errorf("expected %+v to contain %q but it didn't", haystack, needle)
		}
	}
}
