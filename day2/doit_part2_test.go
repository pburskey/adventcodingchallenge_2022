package main

import (
	"adventcodingchallenge_2021/utility"
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
			name: "Sunny Day 1",
			args: args{
				fileName: "data_test.txt",
			},
			want: 5,
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
