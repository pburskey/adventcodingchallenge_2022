package utility

import (
	"reflect"
	"testing"
)

func TestBytesToInt(t *testing.T) {
	type args struct {
		byteArray []byte
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BytesToInt(tt.args.byteArray); got != tt.want {
				t.Errorf("BytesToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLeastAndMax(t *testing.T) {
	type args struct {
		numbers []int
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := LeastAndMax(tt.args.numbers)
			if got != tt.want {
				t.Errorf("LeastAndMax() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("LeastAndMax() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestMean(t *testing.T) {
	type args struct {
		numbers []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "a",
			args: args{
				numbers: []int{11, 23, 30, 47, 56},
			},
			want: 33,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Mean(tt.args.numbers); got != tt.want {
				t.Errorf("Mean() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMedian(t *testing.T) {
	type args struct {
		numbers []int
	}
	tests := []struct {
		name string
		args args
		want float32
	}{
		{
			name: "a",
			args: args{
				numbers: []int{11, 23, 30, 47, 56},
			},
			want: 30,
		},
		{
			name: "b",
			args: args{
				numbers: []int{11, 23, 30, 47, 52, 56},
			},
			want: 38.5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Median(tt.args.numbers); got != tt.want {
				t.Errorf("Median() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNumbersBetween(t *testing.T) {
	type args struct {
		start int
		end   int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NumbersBetween(tt.args.start, tt.args.end); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NumbersBetween() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrderNumbersStartingWithAndEndingWith(t *testing.T) {
	type args struct {
		numbers []int
		start   int
		end     int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := OrderNumbersStartingWithAndEndingWith(tt.args.numbers, tt.args.start, tt.args.end); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrderNumbersStartingWithAndEndingWith() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseDayForInputIntoNumberRows(t *testing.T) {
	type args struct {
		day  string
		path string
	}
	tests := []struct {
		name    string
		args    args
		want    []int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseDayForInputIntoNumberRows(tt.args.day, tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseDayForInputIntoNumberRows() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseDayForInputIntoNumberRows() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseDayForInputIntoStringRows(t *testing.T) {
	type args struct {
		day  string
		path string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseDayForInputIntoStringRows(tt.args.day, tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseDayForInputIntoStringRows() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseDayForInputIntoStringRows() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseInputFileIntoStringRows(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseInputFileIntoStringRows(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseInputFileIntoStringRows() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseInputFileIntoStringRows() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringToInt(t *testing.T) {
	type args struct {
		aString string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringToInt(tt.args.aString); got != tt.want {
				t.Errorf("StringToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_assembleFilePathToDay(t *testing.T) {
	type args struct {
		day string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := assembleFilePathToDay(tt.args.day); got != tt.want {
				t.Errorf("assembleFilePathToDay() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseInputFileIntoNumberRows(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		want    []int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseInputFileIntoNumberRows(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseInputFileIntoNumberRows() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseInputFileIntoNumberRows() got = %v, want %v", got, tt.want)
			}
		})
	}
}
