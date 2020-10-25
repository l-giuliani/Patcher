@echo off
set GOOS=windows
set GOARCH=amd64

set distPath=dist

set libs=patch_library/src/patch_lib
set bin=patch_lib

set param=%1
if "%1"=="" set param=all
goto :%param%

:all

:libs
echo building libs...
(for %%a in (%libs%) do (
   call %%a/build.bat %distPath%/lib
))
if NOT "%param%"=="all" goto :end

:exe
echo building exe...
XCOPY /E /I patcher dist\patcher
copy patcher\patcher.bat dist

:end
echo Done


