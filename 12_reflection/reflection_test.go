package reflection

import (
	"reflect"
	"testing"
)

type Profile struct {
	Age  int
	City string
}

type Person struct {
	Name string
	Profile
}

func TestWalk(t *testing.T) {
	cases := []struct {
		Name  string
		Input interface{}
		Want  []string
	}{
		{
			"Struct with one string field",
			struct {
				Name string
			}{"Tho"},
			[]string{"Tho"},
		},
		{
			"Struct with two string field",
			struct {
				Name string
				City string
			}{"Tho", "Vietnam"},
			[]string{"Tho", "Vietnam"},
		},
		{
			"Struct with non string field",
			struct {
				Name string
				Age  int
			}{"Tho", 23},
			[]string{"Tho"},
		},
		{
			"Struct with nested fileds",
			Person{"Tho", Profile{23, "Vietnam"}},
			[]string{"Tho", "Vietnam"},
		},
		{
			"Pointer to a struct",
			&Person{"Tho", Profile{23, "Vietnam"}},
			[]string{"Tho", "Vietnam"},
		},
		{
			"Slice of structs",
			[]Profile{
				{23, "Vietnam"},
				{32, "Japan"},
			},
			[]string{"Vietnam", "Japan"},
		},
		{
			"Array of structs",
			[2]Profile{
				{23, "Vietnam"},
				{32, "Japan"},
			},
			[]string{"Vietnam", "Japan"},
		},
		{
			"All cases",
			[]struct {
				Person         Person
				MapString      []map[string]string
				ArrayOfPointer [2]*Profile
			}{
				{
					Person{"Tho", Profile{23, "Vietnam"}},
					[]map[string]string{
						{"foo": "bar"},
						{"foo": "baz"},
					},
					[2]*Profile{
						{23, "Vietnam"},
						{32, "Japan"},
					},
				},
			},
			[]string{"Tho", "Vietnam", "bar", "baz", "Vietnam", "Japan"},
		},
	}

	for _, test := range cases {
		var got []string

		t.Run(test.Name, func(t *testing.T) {
			Walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.Want) {
				t.Errorf("got %q want %q", got, test.Want)
			}
		})
	}

	t.Run("Maps of string", func(t *testing.T) {
		aMap := map[string]string{
			"Tho":    "Vietnam",
			"Nagano": "Japan",
		}

		var got []string
		Walk(aMap, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, "Vietnam")
		assertContains(t, got, "Japan")
	})
}

func assertContains(t *testing.T, got []string, value string) {
	contains := false
	for _, x := range got {
		if x == value {
			contains = true
			break
		}
	}

	if !contains {
		t.Errorf("expected %+v to contain %q but it didn't", got, value)
	}
}
