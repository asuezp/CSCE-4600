package builtins

import (
    "fmt"
    "io/ioutil"

)

func ListFiles(args ...string) error {
    var path string
    if len(args) == 0 {
        path = "."
    } else {
        path = args[0]
    }

    files, err := ioutil.ReadDir(path)
    if err != nil {
        return err
    }

    for _, file := range files {
        fmt.Println(file.Name())
    }

    return nil
}