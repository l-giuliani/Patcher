set GOOS=windows
set GOARCH=amd64

set GOPATH=C:\Users\lgiuliani\Projects\Patcher\go_patcher;C:\Users\lgiuliani\Projects\Patcher\patch_library

::go build -ldflags -H=windowsgui -o dist/patcher.exe src/patcher/main.go
go build -o dist/patcher.exe src/patcher/main.go

robocopy src/patcher/gui/resources dist\resources /e