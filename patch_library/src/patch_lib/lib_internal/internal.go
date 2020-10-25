package lib_internal

import (
	"fmt"
	"os"
	"github.com/icedream/go-bsdiff"
)

func GeneratePatch(oldFilePath string, newFilePath string, patchFilePath string) bool {
    oldFile, err := os.Open(oldFilePath)
    if err != nil {
		fmt.Println(err)
        return false
    }
    defer oldFile.Close()
    newFile, err := os.Open(newFilePath)
    if err != nil {
		fmt.Println(err)
        return false
    }
    defer newFile.Close()
    patchFile, err := os.Create(patchFilePath)
    if err != nil {
		fmt.Println("Err create")
		fmt.Println(err)
        return false
    }
    defer patchFile.Close()

    diffErr := bsdiff.Diff(oldFile, newFile, patchFile)
    if diffErr != nil{
		fmt.Println(diffErr)
        return false
    }
    return true
}

func ApplyPatch(oldFilePath string, patchedFilePath string, patchFilePath string) bool {
    oldFile, err := os.Open(oldFilePath)
    if err != nil {
        return false
    }
    defer oldFile.Close()
    newFile, err := os.Create(patchedFilePath)
    if err != nil {
        return false
    }
    defer newFile.Close()
    patchFile, err := os.Open(patchFilePath)
    if err != nil {
        return false
    }
    defer patchFile.Close()

    diffErr := bsdiff.Patch(oldFile, newFile, patchFile)
    if diffErr != nil{
        return false
    }
    return true
}