package main

import (
"C"
"os"
"fmt"
"io"
"io/ioutil"
"bytes"
"path/filepath"
"github.com/icedream/go-bsdiff"
)

//export GeneratePatch
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

//export ApplyPatch
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

func visit(files *[]string, directory string) filepath.WalkFunc {
    return func(path string, info os.FileInfo, err error) error {
        if err != nil {
            fmt.Println(err)
		}
		if !info.IsDir() {
			*files = append(*files, path[len(directory):])
		}        
        return nil
    }
}

func GenerateFolderPatc(){
    new_directory:="iLogger_new"
	old_directory:="Ilogger_old"
	patch_directory_path:= "iLogger_patch"
	
	var new_files []string
	var old_files []string

    err := filepath.Walk(new_directory, visit(&new_files, new_directory))
	if err != nil {
		panic(err)
	}

	err2 := filepath.Walk(old_directory, visit(&old_files, old_directory))
	if err2 != nil {
		panic(err)
	}
	fmt.Println(new_files)

	for i:=0; i<len(new_files); i++ {
		handlePatch(new_directory, old_directory, new_files[i], old_files, patch_directory_path)
	}

}

func handlePatch(new_directory string, old_directory string, new_file string, old_files []string, patch_directory_path string) bool{
	newFile, err := os.Open(new_directory + "/" +new_file)
    if err != nil {
        return false
    }
	defer newFile.Close()
	
	patch_file_name := patch_directory_path + new_file + "_patch"

	founded, index := find(old_files, new_file)
	if !founded {
		pFile, err := os.Create(patch_directory_path + new_file)
		if err != nil {
			return false
		}
		defer pFile.Close()
		io.Copy(pFile, newFile)
	} else {
		if !filesEqual(old_directory + new_file, new_directory + new_file) {
			ensureDir(patch_file_name)
			GeneratePatch(old_directory + new_file, new_directory + new_file, patch_file_name)
		}		
		remove(&old_files, index)
	}
	
	return true
}

func find(slice []string, element string) (bool, int){
	for i:=0;i<len(slice);i++ {
		if slice[i] == element {
			return true, i
		}
	}
	return false, 0
}

func remove(slice *[]string, index int) {
	/*copy(*slice[index:], *slice[index+1:]) 
	*slice[len(slice)-1] = ""     
	*slice = *slice[:len(*slice)-1]*/ 

	(*slice)[index] = (*slice)[len(*slice)-1] 
	(*slice)[len(*slice)-1] = ""   
	*slice = (*slice)[:len(*slice)-1]   
} 

func ensureDir(fileName string) {
	dirName := filepath.Dir(fileName)	
	merr := os.MkdirAll(dirName, os.ModePerm)
	if merr != nil {
		panic(merr)
	}	
}

func filesEqual(file1 string,file2 string) bool {	
	f1, err := ioutil.ReadFile(file1)
    if err != nil {
        return false
    }

    f2, err := ioutil.ReadFile(file2)
    if err != nil {
        return false
    }

    return (bytes.Equal(f1, f2)) 
}

func mainIl() {
    var files []string

    root := "/some/folder/to/scan"
    err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
        files = append(files, path)
        return nil
    })
    if err != nil {
        panic(err)
    }
    for _, file := range files {
        fmt.Println(file)
    }
}

func main()  {
	/*fmt.Println("Test")

	generatePatch()
    err := applyPatch()
    if err != nil{
        fmt.Println(err)
    }
    */

    GenerateFolderPatc()
}