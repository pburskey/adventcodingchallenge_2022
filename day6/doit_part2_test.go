package main

import (
	"adventcodingchallenge_2022/utility"
	"reflect"
	"testing"
)

func TestSolutionPart2WithSampleDataFile(t *testing.T) {
	type args struct {
		data string
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{
			name: "Sunny Day mjqjpqmgbljsphdztnvjfqwrcgsmlb",
			args: args{
				data: "mjqjpqmgbljsphdztnvjfqwrcgsmlb",
			},
			want: 19,
		},
		{
			name: "Sunny Day bvwbjplbgvbhsrlpgdmjqwftvncz",
			args: args{
				data: "bvwbjplbgvbhsrlpgdmjqwftvncz",
			},
			want: 23,
		},
		{
			name: "Sunny Day nppdvjthqldpwncqszvftbrmjlhg",
			args: args{
				data: "nppdvjthqldpwncqszvftbrmjlhg",
			},
			want: 23,
		},
		{
			name: "Sunny Day nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
			args: args{
				data: "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
			},
			want: 29,
		},
		{
			name: "Sunny Day zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
			args: args{
				data: "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
			},
			want: 26,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inputData := utility.ParseInputStringIntoArray(tt.args.data)
			if _, got := solution_part_b(inputData); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
