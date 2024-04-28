package builtins_test

import (
	"github.com/asuezp/CSCE-4600/Project2/builtins"
	"os"
	"testing"
)

func TestPrintWorkingDirectory(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "print working directory",
		},
		// Add more test cases if needed
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := builtins.PrintWorkingDirectory()
			if err != nil {
				t.Fatalf("PrintWorkingDirectory() unexpected error: %v", err)
			}

			wd, _ := os.Getwd()
			t.Logf("Current working directory: %s", wd)
		})
	}
}