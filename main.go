package main

import (
    "errors"
    "fmt"
    "io/fs"
    "os"
)

func main () {
    if _, err:= os.Stat("/workspace/go.mod"); errors.Is(err, fs.ErrNotExist){
        fmt.Println(err.Error())
    } else {
        fmt.Println("file exists!")
    }
}