package scalc

import (
	"testing"
)

func TestParser_Process(t *testing.T) {
	const (
		a = "a.txt"
		b = "b.txt"
		c = "c.txt"
	)

	tests := []struct {
		name    string
		s       string
		want    *Expression
		wantErr bool
	}{
		{
			name: "simple string",
			s:    "[ SUM a.txt ]",
			want: &Expression{
				operator: NewOperator(Sum),
				operands: []*operand{
					&operand{file: a},
				},
			},
			wantErr: false,
		}, {
			name: "two level string",
			s:    "[ SUM [ DIF b.txt c.txt ] a.txt ]",
			want: &Expression{
				operator: NewOperator(Sum),
				operands: []*operand{
					&operand{expr: &Expression{
						operator: NewOperator(Dif),
						operands: []*operand{
							&operand{file: b},
							&operand{file: c},
						},
					}},
					&operand{file: a},
				},
			},
			wantErr: false,
		}, {
			name: "three sets string",
			s:    "[ SUM [ DIF b.txt c.txt ] a.txt [ INT a.txt b.txt ] ]",
			want: &Expression{
				operator: NewOperator(Sum),
				operands: []*operand{
					&operand{expr: &Expression{
						operator: NewOperator(Dif),
						operands: []*operand{
							&operand{file: b},
							&operand{file: c},
						},
					}},
					&operand{file: a},
					&operand{expr: &Expression{
						operator: NewOperator(Int),
						operands: []*operand{
							&operand{file: a},
							&operand{file: b},
						},
					}},
				},
			},
			wantErr: false,
		}, {
			name: "three sets string",
			s:    "[ SUM [ DIF a.txt b.txt c.txt ] [ INT b.txt c.txt ] ]",
			want: &Expression{
				operator: NewOperator(Sum),
				operands: []*operand{
					&operand{expr: &Expression{
						operator: NewOperator(Dif),
						operands: []*operand{
							&operand{file: a},
							&operand{file: b},
							&operand{file: c},
						},
					}},
					&operand{expr: &Expression{
						operator: NewOperator(Int),
						operands: []*operand{
							&operand{file: b},
							&operand{file: c},
						},
					}},
				},
			},
			wantErr: false,
		}, {
			name: "three level string",
			s:    "[ SUM [ DIF b.txt c.txt ] a.txt [ INT a.txt [ SUM b.txt c.txt ] ] ]",
			want: &Expression{
				operator: NewOperator(Sum),
				operands: []*operand{
					&operand{expr: &Expression{
						operator: NewOperator(Dif),
						operands: []*operand{
							&operand{file: b},
							&operand{file: c},
						},
					}},
					&operand{file: a},
					&operand{expr: &Expression{
						operator: NewOperator(Int),
						operands: []*operand{
							&operand{file: a},
							&operand{expr: &Expression{
								operator: NewOperator(Sum),
								operands: []*operand{
									&operand{file: b},
									&operand{file: c},
								},
							}},
						},
					}},
				},
			},
			wantErr: false,
		},
		// Fail cases
		{
			name:    "empty string",
			s:       "",
			wantErr: true,
		}, {
			name:    "empty brackets",
			s:       "[]",
			wantErr: true,
		}, {
			name:    "empty operands",
			s:       "[ DIF ]",
			wantErr: true,
		}, {
			name:    "empty operator",
			s:       "[ a.txt b.txt c.txt]",
			wantErr: true,
		}, {
			name:    "brackets inconsistency",
			s:       "[SUM [DIF a.txt b.txt] [INT a.txt b.txt]",
			wantErr: true,
		}, {
			name:    "no brackets",
			s:       "DIF a.txt b.txt",
			wantErr: true,
		}, {
			name:    "end of the line",
			s:       "[SUM [DIF a.txt b.txt] [INT a.txt b.txt",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := New(tt.s).Process()
			if (err != nil) != tt.wantErr {
				t.Errorf("parseInput() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !got.Equal(tt.want) {
				t.Errorf("parseInput() = %s, want %s", got, tt.want)
			}
		})
	}
}

func Test_cut(t *testing.T) {
	tests := []struct {
		name   string
		runes  []rune
		at, to int
		want   string
	}{
		{
			name:  "at the end",
			runes: []rune{'a', 'b', 'c', 'd', 'e'},
			at:    2,
			to:    5,
			want:  "cde",
		}, {
			name:  "at the middle",
			runes: []rune{'a', 'b', 'c', 'd', 'e'},
			at:    1,
			to:    4,
			want:  "bcd",
		}, {
			name:  "at the start",
			runes: []rune{'a', 'b', 'c', 'd', 'e'},
			at:    0,
			to:    3,
			want:  "abc",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := cut(tt.runes, tt.at, tt.to)
			if got != tt.want {
				t.Errorf("cut() got1 = %s, want %s", got, tt.want)
			}
		})
	}
}
