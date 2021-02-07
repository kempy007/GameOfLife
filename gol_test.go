package main

import (
	"reflect"
	"testing"
)

func Test_predictGOL(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  []string
	}{
		{name: "Test 01", input: []string{
			"4 8",
			"........",
			"....*...",
			"...**...",
			".....*..",
		}, want: []string{
			"4 8",
			"........",
			"...**...",
			"...***..",
			"....*...",
		}},
		{name: "Test 02", input: []string{
			"5 8",
			"........",
			"...**...",
			".*****..",
			"........",
			"........",
		}, want: []string{
			"5 8",
			"........",
			".....*..",
			"..*..*..",
			"..***...",
			"........",
		}},
		{name: "Test 03", input: []string{
			"5 8",
			"........",
			"....*...",
			".*...**.",
			"....*...",
			"........",
		}, want: []string{
			"5 8",
			"........",
			".....*..",
			"....**..",
			".....*..",
			"........",
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := predictGOL(tt.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("predictGOL(tt.input) = %v, want %v", got, tt.want)
			}
		})
	}
}
