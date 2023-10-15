package main

import (
	"adventcodingchallenge_2022/utility"
	"reflect"
	"testing"
)

func TestSolutionPart1WithSampleDataFile(t *testing.T) {
	type args struct {
		fileName string
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{
			name: "Part 1 Test 1",
			args: args{
				fileName: "test_part_1_a.txt",
			},
			want: 31,
		},
		{
			name: "Part 1 Solution",
			args: args{
				fileName: "input.txt",
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inputData, err := utility.ParseDayForInputIntoStringRows(day, tt.args.fileName)
			if err != nil {
				t.Error(err)
			}

			if _, got := solution_part_a(inputData); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
