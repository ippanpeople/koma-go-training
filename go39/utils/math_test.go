package utils

import "testing"

func TestAdd(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "正常處理1",
			args: args{x: 10, y: 20},
			want: 30,
		},
		{
			name: "正常處理2",
			args: args{x: 100, y: 200},
			want: 300,
		},
		{
			name: "正常處理3",
			args: args{x: 1000, y: 2000},
			want: 3000,
		},
		{
			name: "錯誤處理",
			args: args{x: 1000, y: 2000},
			want: 300,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Add(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}
