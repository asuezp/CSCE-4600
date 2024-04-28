package builtins

import (
    "fmt"
    "os"
)

func MakeDirectory(args ...string) error {
    if len(args) == 0 {
        return fmt.Errorf("%w: expected at least one argument (directory name)", ErrInvalidArgCount)
    }

    for _, dir := range args {
        err := os.Mkdir(dir, 0755)
        if err != nil {
            return err
        }
    }

    return nil
}