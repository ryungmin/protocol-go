@echo off

SET VERSION=0.1.2
SET OUT=protocol

IF "%1" == "help" GOTO HELP
IF "%1" == "clean" GOTO CLEAN
IF "%1" == "dist" GOTO DIST

IF "%1" == "386" GOTO X86

IF "%1" == "" GOTO AMD64
IF "%1" == "amd64" GOTO AMD64


:GOBUILD
SET GOOS=windows
SET GOARCH=%~1
go.exe build -o %OUT%.exe -ldflags "-s -w -X 'main.APPLICATION_VERSION=%VERSION%'" .\cmd\protocol\.
EXIT /B 0

:X86
CALL :GOBUILD 386
GOTO EXIT

:AMD64
CALL :GOBUILD amd64
GOTO EXIT

:DIST
del %OUT%.exe
CALL :GOBUILD 386
tar.exe cvf %OUT%-%VERSION%-win-x86.zip %OUT%.exe

CALL :GOBUILD amd64
tar.exe cvf %OUT%-%VERSION%-win-amd64.zip %OUT%.exe

GOTO EXIT

:CLEAN
del %OUT%.exe
GOTO EXIT

:HELP

SET PROG=build.bat

@echo %PROG%
@echo %PROG% amd64  

@echo %PROG% 386

@echo %PROG% clean

:EXIT


