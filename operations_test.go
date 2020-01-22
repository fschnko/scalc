package scalc

import (
	"reflect"
	"testing"
)

func Test_union(t *testing.T) {
	tests := []struct {
		name string
		sets [][]int
		want []int
	}{
		{
			name: "nil sets",
			sets: nil,
			want: []int{},
		}, {
			name: "empty sets",
			sets: [][]int{},
			want: []int{},
		}, {
			name: "one set",
			sets: [][]int{{1, 2, 3}},
			want: []int{1, 2, 3},
		}, {
			name: "equal sets",
			sets: [][]int{{1, 2, 3}, {1, 2, 3}},
			want: []int{1, 2, 3},
		}, {
			name: "two equal sets",
			sets: [][]int{{1, 2, 3}, {4, 5, 6}, {4, 5, 6}},
			want: []int{1, 2, 3, 4, 5, 6},
		}, {
			name: "intersected ranges",
			sets: [][]int{{1, 2, 3, 5}, {4, 6, 8}, {7, 8, 9}},
			want: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		}, {
			name: "reordered sets",
			sets: [][]int{{4, 5, 6}, {1, 2, 3}},
			want: []int{1, 2, 3, 4, 5, 6},
		}, {
			name: "intersected ranges 2",
			sets: [][]int{{1, 2, 4, 5}, {3, 4}, {1, 4}},
			want: []int{1, 2, 3, 4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := union(tt.sets...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Union() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_intersection(t *testing.T) {
	tests := []struct {
		name string
		sets [][]int
		want []int
	}{
		{
			name: "nil sets",
			sets: nil,
			want: nil,
		}, {
			name: "empty sets",
			sets: [][]int{},
			want: nil,
		}, {
			name: "one set",
			sets: [][]int{{1, 2, 3}},
			want: nil,
		}, {
			name: "equal sets",
			sets: [][]int{{1, 2, 3}, {1, 2, 3}},
			want: []int{1, 2, 3},
		}, {
			name: "no intersection sets",
			sets: [][]int{{1, 2, 3}, {4, 5, 6}, {4, 5, 6}},
			want: nil,
		}, {
			name: "intersection of 3",
			sets: [][]int{{1, 2, 3}, {3, 4, 5}, {3, 5, 7}, {3, 7, 9}},
			want: []int{3},
		}, {
			name: "three different sets",
			sets: [][]int{{1, 2, 3, 5}, {2, 3, 5}, {3, 5, 6}},
			want: []int{3, 5},
		}, {
			name: "reordered sets",
			sets: [][]int{{4, 5, 6}, {1, 2, 3, 4}},
			want: []int{4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := intersection(tt.sets...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Intersection() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_difference(t *testing.T) {
	tests := []struct {
		name string
		sets [][]int
		want []int
	}{
		{
			name: "nil sets",
			sets: nil,
			want: nil,
		}, {
			name: "empty sets",
			sets: [][]int{},
			want: nil,
		}, {
			name: "one set",
			sets: [][]int{{1, 2, 3}},
			want: []int{1, 2, 3},
		}, {
			name: "different sets",
			sets: [][]int{{1, 2, 3}, {4, 5, 6}},
			want: []int{1, 2, 3, 4, 5, 6},
		}, {
			name: "equal sets",
			sets: [][]int{{1, 2, 3}, {1, 2, 3}},
			want: []int{},
		}, {
			name: "intersected ranges",
			sets: [][]int{{1, 2, 3, 5}, {4, 6, 7}},
			want: []int{1, 2, 3, 4, 5, 6, 7},
		}, {
			name: "three equal sets",
			sets: [][]int{{1, 2, 3}, {1, 2, 3}, {1, 2, 3}},
			want: []int{},
		}, {
			name: "three different sets",
			sets: [][]int{{1, 2, 3}, {2, 3, 5}, {3, 5, 6}},
			want: []int{1, 2, 5, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := difference(tt.sets...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Difference() = %v, want %v", got, tt.want)
			}
		})
	}
}
