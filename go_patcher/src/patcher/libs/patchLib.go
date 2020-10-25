package libs

import (
	"os"
	"io"
	"encoding/hex"
	"hash/crc32"
	patch_lib "patch_lib/lib_internal"
)

func Generate (old, new, patch string) bool {
	return patch_lib.GeneratePatch(old, new, patch)
}

func Apply (old, patch, patched string) bool {
	return patch_lib.ApplyPatch(old, patched, patch)
}

func ComputeCrc32(filePath string) (string, error) {
	var returnCRC32String string
	
	var polynomial uint32 = 0xedb88320;
	
	file, err := os.Open(filePath)
	if err != nil {
		return returnCRC32String, err
	}
	defer file.Close()
	tablePolynomial := crc32.MakeTable(polynomial)
	hash := crc32.New(tablePolynomial)
	if _, err := io.Copy(hash, file); err != nil {
		return returnCRC32String, err
	}
	hashInBytes := hash.Sum(nil)[:]
	returnCRC32String = hex.EncodeToString(hashInBytes)
	return returnCRC32String, nil
}