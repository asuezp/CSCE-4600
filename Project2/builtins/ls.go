package builtins

import (
    "fmt"
    "io/ioutil"
    "os"
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
        // Check if the error is because the directory does not exist
        if os.IsNotExist(err) {
            return fmt.Errorf("directory does not exist: %s", path)
        }
        return err
    }

    for _, file := range files {
        fmt.Println(file.Name())
    }

    return nil
}
