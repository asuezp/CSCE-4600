package builtins

import (
    "fmt"
    "os"
)

func PrintWorkingDirectory() error {
    wd, err := os.Getwd()
    if err != nil {
        return err
    }

    fmt.Println(wd)
    return nil
}