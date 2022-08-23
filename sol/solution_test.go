package sol

import "testing"

func BenchmarkTest(b *testing.B) {
	s1 := "ab"
	s2 := "eidbaooo"
	for idx := 0; idx < b.N; idx++ {
		checkInclusion(s1, s2)
	}
}
func Test_checkInclusion(t *testing.T) {
	type args struct {
		s1 string
		s2 string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "s1 = \"ab\", s2 = \"eidbaooo\"",
			args: args{s1: "ab", s2: "eidbaooo"},
			want: true,
		},
		{
			name: "s1 = \"ab\", s2 = \"eidboaoo\"",
			args: args{s1: "ab", s2: "eidboaoo"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkInclusion(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("checkInclusion() = %v, want %v", got, tt.want)
			}
		})
	}
}
