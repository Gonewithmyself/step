package backtrack

import (
	"reflect"
	"testing"
)

func Test_newNQueen(t *testing.T) {
	nq := newNQueen()
	t.Log(nq)

	// nq.placeQueenTp()
	t.Log(nq)
}

func Test_packet(t *testing.T) {
	nq := newPacket()
	t.Log(nq)
}

func Test_subsets(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		// TODO: Add test cases.
		{"", args{[]int{0}}, [][]int{}},
		{"", args{[]int{1, 2, 3}}, [][]int{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := subsets(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("subsets() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_subsetswithsubs(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		// TODO: Add test cases.
		{"", args{[]int{0}}, [][]int{}},
		{"", args{[]int{1, 2, 2}}, [][]int{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := subsetsWithdups(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("subsets() = %v, want %v", got, tt.want)
			}
		})
	}
}
