package main

import (
	"testing"
	"reflect"
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

func TestCliCommands(t *testing.T) {
	cases := []struct{
		input string
		want string
	}{
		{
			"exit", "exit",
		},
		{
			"help", "help",
		},
		{
			"help lh alskdh", "help",
		},
		{
			"exit lh alskdh", "exit",
		},
		{
			"ahelp", "unknown command",
		},
		{
			"lkh exit lkhlhk", "unknown command",
		},
		{
			"survey", "unknown command",
		},
		{
			"find", "unknown command",
		},
		{
			"find exit", "unknown command",
		},
	}
	for _, c := range cases {
		got := processCommand(c.input)
		if got != c.want {
			t.Fatalf("Incorrect output (%v) from given input (%v)", got, c.want)
		}
	}
}
