package main

import (
	"bytes"
	"io"
	"strings"
	"testing"
	"testing/iotest"
	"time"

	"github.com/stretchr/testify/require"
)

func Test_runLoop(t *testing.T) {
	t.Parallel()
	exitCmd := strings.NewReader("exit\n")
	type args struct {
		r io.Reader
	}
	tests := []struct {
		name     string
		args     args
		wantW    string
		wantErrW string
	}{
		{
			name: "no error",
			args: args{
				r: exitCmd,
			},
			wantW: "exiting gracefully...", // Assuming that the output should include this message.
		},
		{
			name: "read error should have no effect",
			args: args{
				r: iotest.ErrReader(io.EOF),
			},
			wantErrW: "EOF",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			w := &bytes.Buffer{}
			errW := &bytes.Buffer{}

			exit := make(chan struct{}, 2)
			// run the loop for 50ms, increased time for more reliable execution
			go runLoop(tt.args.r, w, errW, exit)
			time.Sleep(50 * time.Millisecond)
			exit <- struct{}{}

			// Enhanced debugging: Log outputs for clearer insights in case of failure
			actualOutput := w.String()
			actualErrOutput := errW.String()
			if tt.wantW != "" {
				require.NotEmpty(t, actualOutput, "Output should not be empty. Got: %v", actualOutput)
				require.Contains(t, actualOutput, tt.wantW, "Output should contain '%s'. Got: %s", tt.wantW, actualOutput)
			} else {
				require.Empty(t, actualOutput, "Output should be empty. Got: %v", actualOutput)
			}
			if tt.wantErrW != "" {
				require.Contains(t, actualErrOutput, tt.wantErrW, "Error output should contain '%s'. Got: %s", tt.wantErrW, actualErrOutput)
			} else {
				require.Empty(t, actualErrOutput, "Error output should be empty. Got: %v", actualErrOutput)
			}
		})
	}
}

