package main

import (
"C"
"os"
"fmt"
"github.com/icedream/go-bsdiff"
)

//export GeneratePatch
func GeneratePatch(oldFilePath *C.char, newFilePath *C.char, patchFilePath *C.char) C.char {
    oldFile, err := os.Open(C.GoString(oldFilePath))
    if err != nil {
        fmt.Println(err)
        return -1
    }
    defer oldFile.Close()
    newFile, err := os.Open(C.GoString(newFilePath))
    if err != nil {
        fmt.Println(err)
        return -2
    }
    defer newFile.Close()
    patchFile, err := os.Create(C.GoString(patchFilePath))
    if err != nil {
        fmt.Println(err)
        return -3
    }
    defer patchFile.Close()

    diffErr := bsdiff.Diff(oldFile, newFile, patchFile)
    if diffErr != nil{
        fmt.Println(err)
        return -4
    }
    return 0
}

//export ApplyPatch
func ApplyPatch(oldFilePath *C.char, patchedFilePath *C.char, patchFilePath *C.char) C.char {
    oldFile, err := os.Open(C.GoString(oldFilePath))
    if err != nil {
        fmt.Println(err)
        return -1
    }
    defer oldFile.Close()
    newFile, err := os.Create(C.GoString(patchedFilePath))
    if err != nil {
        fmt.Println(err)
        return -2
    }
    defer newFile.Close()
    patchFile, err := os.Open(C.GoString(patchFilePath))
    if err != nil {
        fmt.Println(err)
        return -3
    }
    defer patchFile.Close()

    diffErr := bsdiff.Patch(oldFile, newFile, patchFile)
    if diffErr != nil{
        fmt.Println(err)
        return -4
    }
    return 0
}


func main()  {

}