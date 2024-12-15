package subarrayproductlessthank

import "testing"

func Test_numSubarrayProductLessTanK(t *testing.T) {
	type args struct {
		nums []int
		k    int
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
				nums: []int{10, 5, 2, 6},
				k:    100,
			},
			want: 8,
		},
		{
			name: "Test02",
			args: args{
				nums: []int{1, 2, 3},
				k:    0,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numSubarrayProductLessTanK(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("numSubarrayProductLessTanK() = %v, want %v", got, tt.want)
			}
		})
	}
}
