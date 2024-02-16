package main

import (
	"testing"
)

func Test_Unpacking(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "a4bc2d5e",
			args:    args{line: "a4bc2d5e"},
			want:    "aaaabccddddde",
			wantErr: false,
		},

		{
			name:    "abcd",
			args:    args{line: "abcd"},
			want:    "abcd",
			wantErr: false,
		},

		{
			name:    "45",
			args:    args{line: "45"},
			want:    "",
			wantErr: true,
		},

		{
			name:    "Пустая строка",
			args:    args{line: ""},
			want:    "",
			wantErr: false,
		},

		{
			name:    "qwe\\4\\5",
			args:    args{line: "qwe\\4\\5"},
			want:    "qwe45",
			wantErr: false,
		},

		{
			name:    "qwe\\45",
			args:    args{line: "qwe\\45"},
			want:    "qwe44444",
			wantErr: false,
		},

		{
			name:    "qwe\\\\5",
			args:    args{line: "qwe\\\\5"},
			want:    "qwe\\\\\\\\\\",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Unpacking(tt.args.line)
			if (err != nil) != tt.wantErr {
				t.Errorf("unpacking() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("unpacking() = %v, want %v", got, tt.want)
			}
		})
	}
}
