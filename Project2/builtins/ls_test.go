package builtins_test

import (
	"os"
	"testing"

	"github.com/asuezp/CSCE-4600/Project2/builtins"
)

func TestListFiles(t *testing.T) {
	tmpDir := t.TempDir()
	tmpFile1 := tmpDir + "/file1.txt"
	tmpFile2 := tmpDir + "/file2.txt"
	f1, _ := os.Create(tmpFile1)
	f1.Close()
	f2, _ := os.Create(tmpFile2)
	f2.Close()

	type args struct {
		args []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "list files in current directory",
			args: args{
				args: []string{tmpDir},
			},
		},
		{
			name:    "list files in non-existent directory",
			args:    args{args: []string{"/non-existent"}},
			wantErr: true,
		},
		// Add more test cases as needed
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := builtins.ListFiles(tt.args.args...)
			if (err != nil) != tt.wantErr {
				t.Fatalf("ListFiles() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
