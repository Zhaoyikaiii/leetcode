package removeduplicatesfromsortedarray

import (
	"testing"
)

func Test_removeDuolicates(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "Test01",
			args: args{
				nums: []int{1, 2, 3, 3, 4, 4, 4, 5},
			},
			want: 5,
		},
		{
			name: "Test02",
			args: args{
				nums: []int{1, 1, 2},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeDuolicates(tt.args.nums); got != tt.want {
				t.Errorf("removeDuolicates() = %v, want %v", got, tt.want)
			}
		})
	}
}
