set GOOS=windows
set GOARCH=amd64

set GOPATH=C:\Users\lgiuliani\Projects\Patcher\patch_library

go build -o %1/patchLib.dll -buildmode=c-shared %GOPATH%/src/patch_lib/main/patchLib.go