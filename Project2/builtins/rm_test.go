package builtins_test

import (
	"errors"
	"github.com/asuezp/CSCE-4600/Project2/builtins"
	"os"
	"testing"
)

func TestRemoveFile(t *testing.T) {
	tmpDir := t.TempDir()
	tmpFile := tmpDir + "/testfile.txt"
	f, _ := os.Create(tmpFile)
	f.Close()

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
			name: "remove file",
			args: args{
				args: []string{tmpFile},
			},
		},
		// Add more test cases as needed
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := builtins.RemoveFile(tt.args.args...)
			if tt.wantErr != nil {
				if !errors.Is(err, tt.wantErr) {
					t.Fatalf("RemoveFile() error = %v, wantErr %v", err, tt.wantErr)
				}
				return
			} else if err != nil {
				t.Fatalf("RemoveFile() unexpected error: %v", err)
			}
		})
	}
}