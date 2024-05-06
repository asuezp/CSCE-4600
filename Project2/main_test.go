package main

import (
	"bytes"
	"errors"
	"io"
	"strings"
	"testing"
	"testing/iotest"
	"time"

	"github.com/stretchr/testify/require"
)

func Test_runLoop(t *testing.T) {
    t.Parallel()

    // Prepare scenarios for testing
    tests := []struct {
        name      string
        reader    io.Reader
        wantError string
        wantOutput bool
    }{
        {
            name:      "no error",
            reader:    strings.NewReader("exit\n"),
            wantError: "",
            wantOutput: true,  // Expect some output indicating successful command handling
        },
        {
            name:      "read error should have no effect",
            reader:    iotest.ErrReader(errors.New("EOF")), // Simulating an EOF error
            wantError: "EOF",
            wantOutput: true,
        },
    }

    for _, tt := range tests {
        tt := tt // Capture range variable
        t.Run(tt.name, func(t *testing.T) {
            t.Parallel()

            outputBuffer := &bytes.Buffer{}
            errorBuffer := &bytes.Buffer{}
            exitChannel := make(chan struct{}, 2) // Buffer to prevent deadlocks

            // Run the function in a goroutine to allow asynchronous operation
            go runLoop(tt.reader, outputBuffer, errorBuffer, exitChannel)

            // Give the function time to process the input
            time.Sleep(50 * time.Millisecond)  // Adjusted time for better reliability
            exitChannel <- struct{}{}  // Signal to exit the run loop

            // Validate error output if expected
            if tt.wantError != "" {
                require.Contains(t, errorBuffer.String(), tt.wantError,
                    "Error output should contain '%s'. Actual output: '%s'", tt.wantError, errorBuffer.String())
            } else {
                require.Empty(t, errorBuffer.String(),
                    "Error output should be empty. Actual output: '%s'", errorBuffer.String())
            }

            // Validate standard output if expected
            if tt.wantOutput {
                require.NotEmpty(t, outputBuffer.String(),
                    "Expected non-empty output, but got empty. Output: '%s'", outputBuffer.String())
            } else {
                require.Empty(t, outputBuffer.String(),
                    "Expected empty output, but got non-empty. Output: '%s'", outputBuffer.String())
            }
        })
    }
}
