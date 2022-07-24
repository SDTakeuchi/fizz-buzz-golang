package main

import (
	"testing"
)

func TestFizzBuzz(t *testing.T) {
	// define arguments
	type args struct {
		n int
	}

	// set up test data
	tests := []struct {
		name string
		args args 
		want string
	}{
		{
			name: "fizz",
			args: args{33},
			want: "Fizz",
		},
		{
			name: "buzz",
			args: args{155},
			want: "Buzz",
		},
		{
			name: "normal",
			args: args{119},
			want: "119",
		},
		{
			name: "fizzbuzz",
			args: args{165},
			want: "FizzBuzz",
		},
	}

	// run tests
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fizzBuzz(tt.args.n); got != tt.want {
				t.Errorf("add() = %v, want %v", got, tt.want)
			}
		})
	}
}