package main

import (
	"reflect"
	"testing"
)

func TestA_highlight(t *testing.T) {
	type args struct {
		s      string
		subStr string
		n      int
	}
	tests := []struct {
		name string
		a    *A
		args args
		want []string
	}{
		{
			name: "-a",
			args: args{
				s: "Welcome to Linux !\n" +
					"Linux is a free and opensource Operating system that is mostly used by\n" +
					"developers and in production servers for hosting crucial components such as web\n" +
					"and database servers. Linux has also made a name for itself in PCs.\n" +
					"Beginners looking to experiment with Linux can get started with friendlier linux\n" +
					"distributions such as Ubuntu, Mint, Fedora and Elementary OS.",
				subStr: "Linux",
				n:      1,
			},
			want: []string{
				"Welcome to Linux !\nLinux is a free and opensource Operating system that is mostly used by\n",
				"Linux is a free and opensource Operating system that is mostly used by\ndevelopers and in production servers for hosting crucial components such as web\n",
				"and database servers. Linux has also made a name for itself in PCs.\nBeginners looking to experiment with Linux can get started with friendlier linux\n",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.highlight(tt.args.s, tt.args.subStr, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("A.highlight() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_c_counter(t *testing.T) {
	type args struct {
		lines  string
		subStr string
	}
	tests := []struct {
		name      string
		c         *c
		args      args
		wantCount int
	}{
		{
			name: "-c",
			args: args{
				lines: "Welcome to Linux !\n" +
					"Linux is a free and opensource Operating system that is mostly used by\n" +
					"developers and in production servers for hosting crucial components such as web\n" +
					"and database servers. Linux has also made a name for itself in PCs.\n" +
					"Beginners looking to experiment with Linux can get started with friendlier linux\n" +
					"distributions such as Ubuntu, Mint, Fedora and Elementary OS.",
				subStr: "Linux",
			},
			wantCount: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotCount := tt.c.counter(tt.args.lines, tt.args.subStr); gotCount != tt.wantCount {
				t.Errorf("c.counter() = %v, want %v", gotCount, tt.wantCount)
			}
		})
	}
}

func TestN_find(t *testing.T) {
	type args struct {
		lines  string
		subStr string
	}
	tests := []struct {
		name string
		n    *N
		args args
		want int
	}{
		{
			name: "-n",
			args: args{
				lines: "Welcome to Linux !\n" +
					"Linux is a free and opensource Operating system that is mostly used by\n" +
					"developers and in production servers for hosting crucial components such as web\n" +
					"and database servers. Linux has also made a name for itself in PCs.\n" +
					"Beginners looking to experiment with Linux can get started with friendlier linux\n" +
					"distributions such as Ubuntu, Mint, Fedora and Elementary OS.",
				subStr: "Linux",
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.n.find(tt.args.lines, tt.args.subStr); got != tt.want {
				t.Errorf("N.find() = %v, want %v", got, tt.want)
			}
		})
	}
}
