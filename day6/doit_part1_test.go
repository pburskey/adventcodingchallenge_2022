package main

import (
	"adventcodingchallenge_2022/utility"
	"reflect"
	"testing"
)

func TestSolutionPart1WithSampleDataFile(t *testing.T) {
	type args struct {
		data string
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{
			name: "Sunny Day 1",
			args: args{
				data: "mjqjpqmgbljsphdztnvjfqwrcgsmlb",
			},
			want: 7,
		},
		{
			name: "Sunny Day 2",
			args: args{
				data: "bvwbjplbgvbhsrlpgdmjqwftvncz",
			},
			want: 5,
		},
		{
			name: "Sunny Day 3",
			args: args{
				data: "nppdvjthqldpwncqszvftbrmjlhg",
			},
			want: 6,
		},
		{
			name: "Sunny Day 4",
			args: args{
				data: "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
			},
			want: 10,
		},
		{
			name: "Sunny Day 5",
			args: args{
				data: "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
			},
			want: 11,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inputData := utility.ParseInputStringIntoArray(tt.args.data)

			if _, got := solution_part_a(inputData); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
