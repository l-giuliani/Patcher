package utils

import (
	"fmt"
	"os"
	"io/ioutil"
)

func WriteStringToFile(fileName, data string) bool{
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, 0644)
	defer file.Close()
	if err != nil {
		fmt.Println(err)
		return false
	}
	_, errw := file.WriteString(data)
	if (errw != nil) {
		fmt.Println(errw)
		return false
	}
	return true
}

func ReadStringFromFile(fileName string) (string ,bool) {
	dat, err := ioutil.ReadFile(fileName)
	if (err != nil) {
		fmt.Println(err)
		return "", false
	}
	return string(dat), true
}