package main

import "testing"

func Test_orFn(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "Deadlock checking",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := orFn(); (err != nil) != tt.wantErr {
				t.Errorf("orFn() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
