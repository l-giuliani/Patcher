package services
import (
	"fmt"
	libs "patcher/libs"
	utils "patcher/utils"
)

func GeneratePatch(params []string) bool{
	res := libs.Generate(params[0], params[1], params[2])
	if (!res){
		fmt.Println("Error Generating Patch")
	}
	return res
}

func ApplyPatch(params []string) bool{
	res := libs.Apply(params[0], params[1], params[2])
	if (!res) {
		fmt.Println("Error Applying Patch")
	}
	return res
}

func GeneratePatchWithCrc32(params []string) bool{
	res := GeneratePatch(params[:3])
	if (!res) {
		return false
	}
	new := params[1]
	crc32, err := libs.ComputeCrc32(new)
	if (err != nil){
		fmt.Println("Error Generating crc32")
		return false
	}
	crc32FileName := params[3]
	res = utils.WriteStringToFile(crc32FileName, crc32)
	if (!res) {
		fmt.Println("Error Writing crc32 file")
		return false
	}
	return true
}

func ApplyPatchWithCrc32(params []string) bool{
	res := ApplyPatch(params[:3])
	if (!res) {
		return false
	}
	patched := params[2]
	crc32Patched, err := libs.ComputeCrc32(patched)
	if (err != nil){
		fmt.Println("Error Generating crc32")
		return false
	}
	crc32FileName := params[3]
	crc32NewFile, resr :=utils.ReadStringFromFile(crc32FileName) 
	if (!resr){
		fmt.Println("Error reading crc32 file")
		return false
	}
	if (crc32Patched == crc32NewFile){
		//fmt.Println("Success")
		return true
	}
	fmt.Println("crc32 not equals")
	return false
}

func PrintCrc32(params []string) bool{
	crc32, err := libs.ComputeCrc32(params[0])
	if (err != nil){
		fmt.Println("Error Computing crc32")
		return false
	}
	fmt.Println(crc32)
	return true
}