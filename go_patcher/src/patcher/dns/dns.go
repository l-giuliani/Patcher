package dns

import (
	services "patcher/services"
)

var ServiceMap = map[string](func([]string)bool){
    "generate": services.GeneratePatch,
    "apply": services.ApplyPatch,
    "generate-crc32": services.GeneratePatchWithCrc32,
    "apply-crc32": services.ApplyPatchWithCrc32,
    "print-crc32": services.PrintCrc32 }