package cmp

import "testing"

func TestBFcmp(t *testing.T) {
	type args struct {
		s string
		p string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"", args{"abcd", "abcd"}, 0},
		{"", args{"abcd", "bcd"}, 1},
		{"", args{"abcd", "cd"}, 2},
		{"", args{"abcd", "d"}, 3},
		{"", args{"abcd", "cde"}, -1},
		{"", args{"aaabaaabaaabaaab", "aaaa"}, -1},
		{"", args{"aaaaaaaaaaaaaaaa", "baaa"}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BFcmp(tt.args.s, tt.args.p); got != tt.want {
				t.Errorf("BMcmp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBMcmp(t *testing.T) {
	type args struct {
		s string
		p string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"", args{"abcd", "abcd"}, 0},
		{"", args{"abcd", "bcd"}, 1},
		{"", args{"abcd", "cd"}, 2},
		{"", args{"abcd", "d"}, 3},
		{"", args{"abcd", "cde"}, -1},
		{"", args{"aaabaaabaaabaaab", "aaaa"}, -1},
		{"", args{"aaaaaaaaaaaaaaaa", "baaa"}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BMcmpWithGoodSuffix(tt.args.s, tt.args.p); got != tt.want {
				t.Errorf("BMcmp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKMPcmp(t *testing.T) {
	type args struct {
		s string
		p string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"", args{"ababaeabac", "ababacd"}, -1},
		{"", args{"abcd", "bcd"}, 1},
		{"", args{"abcd", "cd"}, 2},
		{"", args{"abcd", "d"}, 3},
		{"", args{"abcd", "cde"}, -1},
		{"", args{"aaabaaabaaabaaab", "aaaa"}, -1},
		{"", args{"aaaaaaaaaaaaaaaa", "baaa"}, -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := KMPcmp(tt.args.s, tt.args.p); got != tt.want {
				t.Errorf("BMcmp() = %v, want %v", got, tt.want)
			} else if got != -1 {
				if tt.args.s[got:got+len(tt.args.p)] != tt.args.p {
					t.Errorf("BMcmp() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func BenchmarkBFcmp(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// aaabaaabaaabaaab
		BFcmp("aaaaaaaaaaaaaaaa", "baaa")
	}
}

func BenchmarkBMcmp(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// BMcmp("abcacabdc", "abd")
		BMcmp("aaaaaaaaaaaaaaaa", "baaa")
	}
}
