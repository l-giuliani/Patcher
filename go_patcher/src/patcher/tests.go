package main

import (
	"fmt"
	libs "patcher/libs"
	utils "patcher/utils"
)

func crc(){
	res, _ := libs.ComputeCrc32("C:/patchTests/fer-0.1.0_new.war")
	fmt.Println(res)
	utils.WriteStringToFile("C:/patchTests/crcfile.txt", res)
}

func main() {
	crc()
}

