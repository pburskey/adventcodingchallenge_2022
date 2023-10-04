package main

import (
	"adventcodingchallenge_2022/utility"
	"reflect"
	"testing"
)

func TestSolutionPart2WithSampleDataFile(t *testing.T) {
	type args struct {
		fileName string
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{
			name: "Part 2 Test 1",
			args: args{
				fileName: "test_part_1_a.txt",
			},
			want: 2713310158,
		},
		//{
		//	name: "Sunny Day 2",
		//	args: args{
		//		fileName: "test_part_2_a.txt",
		//	},
		//	want: 36,
		//},
		{
			name: "Part 2 Solution",
			args: args{
				fileName: "input.txt",
			},
			want: 25272176808,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inputData, err := utility.ParseDayForInputIntoStringRows(day, tt.args.fileName)
			if err != nil {
				t.Error(err)
			}

			if _, got := solution_part_b(inputData); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
