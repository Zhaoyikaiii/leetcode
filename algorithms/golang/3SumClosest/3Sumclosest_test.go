package sumclosest

import "testing"

func Test_thressSumClosest(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "test01",
			args: args{
				nums:   []int{-1, 2, 1, -4},
				target: 1,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := thressSumClosest(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("thressSumClosest() = %v, want %v", got, tt.want)
			}
		})
	}
}
