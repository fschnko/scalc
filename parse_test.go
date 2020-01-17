package scalc

import (
	"testing"
)

func Test_parseInput(t *testing.T) {
	var (
		a = "a.txt"
		b = "b.txt"
		c = "c.txt"
	)

	tests := []struct {
		name    string
		s       string
		want    *calc
		wantErr bool
	}{
		{
			name: "simple string",
			s:    "[ SUM a.txt ]",
			want: &calc{
				Expr: &expression{
					Operator: Sum,
					Sets: []*set{
						&set{File: a},
					},
				},
			},
			wantErr: false,
		}, {
			name: "two level string",
			s:    "[ SUM [ DIF b.txt c.txt ] a.txt ]",
			want: &calc{
				Expr: &expression{
					Operator: Sum,
					Sets: []*set{
						&set{Expr: &expression{
							Operator: Dif,
							Sets: []*set{
								&set{File: b},
								&set{File: c},
							},
						}},
						&set{File: a},
					},
				},
			},
			wantErr: false,
		}, {
			name: "three sets string",
			s:    "[ SUM [ DIF b.txt c.txt ] a.txt [ INT a.txt b.txt ] ]",
			want: &calc{
				Expr: &expression{
					Operator: Sum,
					Sets: []*set{
						&set{Expr: &expression{
							Operator: Dif,
							Sets: []*set{
								&set{File: b},
								&set{File: c},
							},
						}},
						&set{File: a},
						&set{Expr: &expression{
							Operator: Int,
							Sets: []*set{
								&set{File: a},
								&set{File: b},
							},
						}},
					},
				},
			},
			wantErr: false,
		}, {
			name: "three sets string",
			s:    "[ SUM [ DIF a.txt b.txt c.txt ] [ INT b.txt c.txt ] ]",
			want: &calc{
				Expr: &expression{
					Operator: Sum,
					Sets: []*set{
						&set{Expr: &expression{
							Operator: Dif,
							Sets: []*set{
								&set{File: a},
								&set{File: b},
								&set{File: c},
							},
						}},
						&set{Expr: &expression{
							Operator: Int,
							Sets: []*set{
								&set{File: b},
								&set{File: c},
							},
						}},
					},
				},
			},
			wantErr: false,
		}, {
			name: "three level string",
			s:    "[ SUM [ DIF b.txt c.txt ] a.txt [ INT a.txt [ SUM b.txt c.txt ] ] ]",
			want: &calc{
				Expr: &expression{
					Operator: Sum,
					Sets: []*set{
						&set{Expr: &expression{
							Operator: Dif,
							Sets: []*set{
								&set{File: b},
								&set{File: c},
							},
						}},
						&set{File: a},
						&set{Expr: &expression{
							Operator: Int,
							Sets: []*set{
								&set{File: a},
								&set{Expr: &expression{
									Operator: Sum,
									Sets: []*set{
										&set{File: b},
										&set{File: c},
									},
								}},
							},
						}},
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseInput(tt.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseInput() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !got.Expr.Equal(tt.want.Expr) {
				t.Errorf("parseInput() = %#v, want %#v", got, tt.want)
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
