package builtins

import (
    "fmt"
    "os"
)

func RemoveFile(args ...string) error {
    if len(args) == 0 {
        return fmt.Errorf("%w: expected at least one argument (file path)", ErrInvalidArgCount)
    }

    for _, path := range args {
        err := os.Remove(path)
        if err != nil {
            return err
        }
    }

    return nil
}