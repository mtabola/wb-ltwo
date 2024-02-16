package main

import (
	"fmt"
	"testing"
)

func Test_doSomething(t *testing.T) {
	type args struct {
		fields    []int
		delimeter string
		separated bool
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Чек",
			args: args{
				fields:    []int{1, 3},
				delimeter: ",",
				separated: false,
			},
			want: "Имя Город\nАнна Москва\nИван Санкт-Петербург\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := doSomething(tt.args.fields, tt.args.delimeter, tt.args.separated); got != tt.want {
				fmt.Println([]byte(tt.want))
				fmt.Println([]byte(got))
				t.Errorf("doSomething() = %v, want %v", got, tt.want)
			}
		})
	}
}
