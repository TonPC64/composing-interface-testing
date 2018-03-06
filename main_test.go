package main

import (
	"testing"
)

func MockGetValue(api string) int {
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

func TestGetValue(t *testing.T) {
	type args struct {
		api string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "https://api.myjson.com/bins/8hykh",
			args: args{
				api: "https://api.myjson.com/bins/8hykh",
			},
			want: 8,
		},
		{
			name: "no api",
			args: args{
				api: "",
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetValue(tt.args.api); got != tt.want {
				t.Errorf("GetValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
