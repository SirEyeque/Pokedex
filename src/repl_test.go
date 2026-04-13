package main

import (
	"testing"
	"reflect"
	"os"
	"io"
	"fmt"
	"bytes"
)


func TestCleanInput(t *testing.T) {
	cases := []struct{
		input string
		want []string
	}{
		{
			input: "This is the part",
			want: []string{"this", "is", "the", "part"},
		},
		{
			input: " pika         pi  ",
			want: []string{"pika","pi"},
		},
		{
			input: " I like  fortnite ",
			want: []string{"i","like","fortnite"},
		},
	}
	for _, c := range cases {
		got := CleanInput(c.input)
		if len(got) != len(c.want) {
			t.Fatalf("Lengths do not match. Want: len(%v), Got: len(%v)", len(c.want), len(got)) 
		} else if !reflect.DeepEqual(got, c.want) {
			t.Fatalf("String slices do not match.\nWant: %v, Got: %v", c.want, got)
		}
	}
}

var out io.Writer = os.Stdout

func TestCliCommands(t *testing.T) {
	for _, test := range []struct {
		Args   []string
		Output string
	}{
		{
			Args:   []string{"./Pokedex", "5", "7", "9"},
			Output: "",
		},
		{
			Args:   []string{"./calc", "-mode", "multiply", "3", "2", "5"},
			Output: "",
		},
	} {
		t.Run("", func(t *testing.T) {
			os.Args = test.Args
			out = bytes.NewBuffer(nil)
			main()


			if actual := out.(*bytes.Buffer).String(); actual != test.Output {
				fmt.Println(actual, test.Output)
				t.Errorf("expected %s, but got %s", test.Output, actual)
			}
		})
	}
}
