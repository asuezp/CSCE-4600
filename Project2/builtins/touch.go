package builtins

import (
    "fmt"
    "os"
    "time"
)

func TouchFile(args ...string) error {
    if len(args) == 0 {
        return fmt.Errorf("%w: expected at least one argument (file path)", ErrInvalidArgCount)
    }

    for _, path := range args {
        file, err := os.OpenFile(path, os.O_CREATE|os.O_APPEND, 0644)
        if err != nil {
            return err
        }
        file.Close()

        err = os.Chtimes(path, time.Now(), time.Now())
        if err != nil {
            return err
        }
    }

    return nil
}