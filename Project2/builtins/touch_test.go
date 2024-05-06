package builtins_test

import (
	"errors"
	"github.com/asuezp/CSCE-4600/Project2/builtins"
	"testing"
)

func TestTouchFile(t *testing.T) {
	tmpDir := t.TempDir()
	tmpFile := tmpDir + "/testfile.txt"

	type args struct {
		args []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr error
	}{
		{
			name:    "no args",
			args:    args{},
			wantErr: builtins.ErrInvalidArgCount,
		},
		{
			name: "touch file",
			args: args{
				args: []string{tmpFile},
			},
		},
		// Add more test cases as needed
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := builtins.TouchFile(tt.args.args...)
			if tt.wantErr != nil {
				if !errors.Is(err, tt.wantErr) {
					t.Fatalf("TouchFile() error = %v, wantErr %v", err, tt.wantErr)
				}
				return
			} else if err != nil {
				t.Fatalf("TouchFile() unexpected error: %v", err)
			}
		})
	}
}