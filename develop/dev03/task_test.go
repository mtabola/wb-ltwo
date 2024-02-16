package main

import (
	"reflect"
	"testing"
)

func Test_mySort(t *testing.T) {
	type args struct {
		line []string
		key  Sort
		opt  []int
	}
	tests := []struct {
		name       string
		args       args
		wantResult []string
		wantErr    bool
	}{
		{
			name: "-k",
			args: args{line: []string{
				"1 C",
				"2 B",
				"3 A",
			}, key: &K{}, opt: []int{2}},
			wantResult: []string{
				"3 A",
				"2 B",
				"1 C",
			},
			wantErr: false,
		},

		{
			name: "-n",
			args: args{line: []string{
				"10",
				"2",
				"25",
				"1",
			}, key: &N{}},
			wantResult: []string{
				"1",
				"2",
				"10",
				"25",
			},
			wantErr: false,
		},

		{
			name: "-r",
			args: args{line: []string{
				"Z",
				"A",
				"M",
			}, key: &R{}},
			wantResult: []string{
				"Z",
				"M",
				"A",
			},
			wantErr: false,
		},

		{
			name: "-u",
			args: args{line: []string{
				"Apple",
				"Banana",
				"Apple",
			}, key: &U{}},
			wantResult: []string{
				"Apple",
				"Banana",
			},
			wantErr: false,
		},

		{
			name: "-b",
			args: args{line: []string{
				"   A",
				"B  ",
				"C",
			}, key: &B{}},
			wantResult: []string{
				"A",
				"B",
				"C",
			},
			wantErr: false,
		},

		{
			name: "-c",
			args: args{line: []string{
				"B",
				"A",
				"C",
			}, key: &C{}},
			wantResult: nil,
			wantErr:    true,
		},

		{
			name: "-h",
			args: args{line: []string{
				"5M",
				"10K",
				"2M",
				"1K",
			}, key: &H{}},
			wantResult: []string{
				"1K",
				"10K",
				"2M",
				"5M",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResult, err := mySort(tt.args.line, tt.args.key, tt.args.opt...)
			if (err != nil) != tt.wantErr {
				t.Errorf("mySort() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("mySort() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
