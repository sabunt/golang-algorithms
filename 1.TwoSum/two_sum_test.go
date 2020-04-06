package two_sum

import (
	"reflect"
	"testing"
)

func Test_twoSum1(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{{"ast", args{
		nums:   []int{2, 7, 11, 15},
		target: 18,
	},
		[]int{1, 2},
	},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoSum1(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_twoSum2(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{{"ast", args{
		nums:   []int{2, 7, 11, 15},
		target: 18,
	},
		[]int{1, 2},
	},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoSum2(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSum2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_twoSum3(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{{"ast", args{
		nums:   []int{2, 7, 11, 15},
		target: 18,
	},
		[]int{1, 2},
	},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoSum3(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSum3() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTwoNumberSum1(t *testing.T) {
	type args struct {
		array  []int
		target int
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{"Test case 1", args{[]int{4, 6}, 10}, []int{4, 6}},
		{"Test case 2", args{[]int{4, 6, 1}, 5}, []int{4, 1}},
		{"Test case 3", args{[]int{4, 6, 1, -3}, 3}, []int{6, -3}},
		{"Test case 4", args{[]int{3, 5, -4, 8, 11, 1, -1, 6}, 10}, []int{11, -1}},
		{"Test case 5", args{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 17}, []int{8, 9}},
		{"Test case 6", args{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 15}, 18}, []int{3, 15}},
		{"Test case 7", args{[]int{-7, -5, -3, -1, 0, 1, 3, 5, 7}, -5}, []int{-5, 0}},
		{"Test case 8", args{[]int{-21, 301, 12, 4, 65, 56, 210, 356, 9, -47}, 163}, []int{210, -47}},
		{"Test case 9", args{[]int{-21, 301, 12, 4, 65, 56, 210, 356, 9, -47}, 164}, []int{}},
		{"Test case 10", args{[]int{3, 5, -4, 8, 11, 1, -1, 6}, 15}, []int{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TwoNumberSum1(tt.args.array, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TwoNumberSum1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTwoNumberSum2(t *testing.T) {
	type args struct {
		array  []int
		target int
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{"Test case 1", args{[]int{4, 6}, 10}, []int{4, 6}},
		{"Test case 2", args{[]int{4, 6, 1}, 5}, []int{4, 1}},
		{"Test case 3", args{[]int{4, 6, 1, -3}, 3}, []int{6, -3}},
		{"Test case 4", args{[]int{3, 5, -4, 8, 11, 1, -1, 6}, 10}, []int{11, -1}},
		{"Test case 5", args{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 17}, []int{8, 9}},
		{"Test case 6", args{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 15}, 18}, []int{3, 15}},
		{"Test case 7", args{[]int{-7, -5, -3, -1, 0, 1, 3, 5, 7}, -5}, []int{-5, 0}},
		{"Test case 8", args{[]int{-21, 301, 12, 4, 65, 56, 210, 356, 9, -47}, 163}, []int{210, -47}},
		{"Test case 9", args{[]int{-21, 301, 12, 4, 65, 56, 210, 356, 9, -47}, 164}, []int{}},
		{"Test case 10", args{[]int{3, 5, -4, 8, 11, 1, -1, 6}, 15}, []int{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TwoNumberSum2(tt.args.array, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TwoNumberSum2() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestTwoNumberSum3(t *testing.T) {
	type args struct {
		array  []int
		target int
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{"Test case 1", args{[]int{4, 6}, 10}, []int{4, 6}},
		{"Test case 2", args{[]int{4, 6, 1}, 5}, []int{1, 4}},
		{"Test case 3", args{[]int{4, 6, 1, -3}, 3}, []int{-3, 6}},
		{"Test case 4", args{[]int{3, 5, -4, 8, 11, 1, -1, 6}, 10}, []int{-1, 11}},
		{"Test case 5", args{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 17}, []int{8, 9}},
		{"Test case 6", args{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 15}, 18}, []int{3, 15}},
		{"Test case 7", args{[]int{-7, -5, -3, -1, 0, 1, 3, 5, 7}, -5}, []int{-5, 0}},
		{"Test case 8", args{[]int{-21, 301, 12, 4, 65, 56, 210, 356, 9, -47}, 163}, []int{-47, 210}},
		{"Test case 9", args{[]int{-21, 301, 12, 4, 65, 56, 210, 356, 9, -47}, 164}, []int{}},
		{"Test case 10", args{[]int{3, 5, -4, 8, 11, 1, -1, 6}, 15}, []int{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TwoNumberSum3(tt.args.array, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TwoNumberSum3() = %v, want %v", got, tt.want)
			}
		})
	}
}
