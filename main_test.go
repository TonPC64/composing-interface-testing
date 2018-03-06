package main

import "testing"

func MockGetValue() int {
	return 6
}

func Test_useApi(t *testing.T) {
	type args struct {
		cf constructureFunctions
	}

	cf := constructureFunctions{
		GetValue: MockGetValue,
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test with mock api",
			args: args{
				cf: cf,
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := useApi(tt.args.cf); got != tt.want {
				t.Errorf("useApi() = %v, want %v", got, tt.want)
			}
		})
	}
}
