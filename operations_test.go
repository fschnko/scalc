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
			name: "non normalized sets",
			sets: [][]int{{3, 3, 3}, {3, 3, 3}, {3, 3, 3}},
			want: []int{3},
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

func Benchmark_union(b *testing.B) {
	for n := 0; n < b.N; n++ {
		union([]int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			[]int{11, 12, 13, 14, 15, 16, 17, 18, 19},
			[]int{21, 22, 23, 24, 25, 26, 27, 28, 29},
			[]int{31, 32, 33, 34, 35, 36, 37, 38, 39},
		)
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
			name: "no intersection sets",
			sets: [][]int{{1, 2, 3}, {4, 5, 6}, {4, 5, 6}},
			want: []int{},
		}, {
			name: "intersection of 3",
			sets: [][]int{{1, 2, 3}, {3, 4, 5}, {5, 3, 7}, {7, 3, 9}},
			want: []int{3},
		}, {
			name: "non normalized sets",
			sets: [][]int{{3, 3, 3}, {3, 3, 3}, {3, 3, 3}},
			want: []int{3},
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

func Benchmark_intersection(b *testing.B) {
	for n := 0; n < b.N; n++ {
		intersection([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			[]int{11, 12, 13, 14, 15, 16, 17, 18, 19, 10},
			[]int{21, 22, 23, 24, 25, 26, 27, 28, 29, 10},
			[]int{31, 32, 33, 34, 35, 36, 37, 38, 39, 10},
		)
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
			want: nil,
		}, {
			name: "three equal sets",
			sets: [][]int{{1, 2, 3}, {1, 2, 3}, {1, 2, 3}},
			want: nil,
		}, {
			name: "non normalized sets",
			sets: [][]int{{1, 2}, {3, 3, 3}, {3, 3, 3}, {3, 3, 3}},
			want: []int{1, 2, 3},
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

func Benchmark_difference(b *testing.B) {
	for n := 0; n < b.N; n++ {
		difference([]int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			[]int{11, 12, 13, 14, 15, 16, 17, 18, 19},
			[]int{21, 22, 23, 24, 25, 26, 27, 28, 29},
			[]int{31, 32, 33, 34, 35, 36, 37, 38, 39},
		)
	}
}

func Test_normalize(t *testing.T) {

	tests := []struct {
		name string
		a    []int
		want []int
	}{
		{
			name: "nil set",
			a:    nil,
			want: nil,
		}, {
			name: "empty set",
			a:    []int{},
			want: []int{},
		}, {
			name: "duplication",
			a:    []int{1, 1, 2, 2, 3, 3},
			want: []int{1, 2, 3},
		}, {
			name: "triplication",
			a:    []int{1, 1, 1, 2, 2, 2, 3, 3, 3},
			want: []int{1, 2, 3},
		}, {
			name: "normalized set",
			a:    []int{1, 2, 3},
			want: []int{1, 2, 3},
		}, {
			name: "non sorted set",
			a:    []int{1, 2, 3, 1, 2, 3},
			want: []int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := normalize(tt.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("normalize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_diff(t *testing.T) {
	type args struct {
		a []int
		b []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "nil sets",
			args: args{
				a: nil,
				b: nil,
			},
			want: []int{},
		}, {
			name: "empty sets",
			args: args{
				a: []int{},
				b: []int{},
			},
			want: []int{},
		}, {
			name: "one empty set",
			args: args{
				a: []int{1, 2, 3},
				b: []int{},
			},
			want: []int{1, 2, 3},
		}, {
			name: "different sets",
			args: args{
				a: []int{1, 2, 3},
				b: []int{4, 5, 6},
			},
			want: []int{1, 2, 3, 4, 5, 6},
		}, {
			name: "equal sets",
			args: args{
				a: []int{1, 2, 3},
				b: []int{1, 2, 3},
			},
			want: []int{},
		}, {
			name: "intersected sets",
			args: args{
				a: []int{1, 2, 3, 4},
				b: []int{4, 5, 6},
			},
			want: []int{1, 2, 3, 5, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := diff(tt.args.a, tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("difference() = %v, want %v", got, tt.want)
			}
		})
	}
}
