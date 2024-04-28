package builtins_test

import (
    "errors"
    "github.com/asuezp/CSCE-4600/Project2/builtins"
    "testing"
)

func TestMakeDirectory(t *testing.T) {
    tmpDir := t.TempDir()

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
            name: "create directory",
            args: args{
                args: []string{tmpDir + "/testdir"},
            },
        },
        // Add more test cases as needed
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            err := builtins.MakeDirectory(tt.args.args...)
            if tt.wantErr != nil {
                if !errors.Is(err, tt.wantErr) {
                    t.Fatalf("MakeDirectory() error = %v, wantErr %v", err, tt.wantErr)
                }
                return
            } else if err != nil {
                t.Fatalf("MakeDirectory() unexpected error: %v", err)
            }
        })
    }
}