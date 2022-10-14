@echo off

SET VERSION=0.1.0
SET OUT=protocol

IF "%1" == "help" GOTO HELP
IF "%1" == "clean" GOTO CLEAN

IF "%1" == "386" GOTO X86

IF "%1" == "" GOTO AMD64
IF "%1" == "amd64" GOTO AMD64

:X86
SET GOOS=windows
SET GOARCH=386
GOTO BUILD

:AMD64
SET GOOS=windows
SET GOARCH=amd64
GOTO BUILD

:BUILD
@echo Building %OUT%.exe (%GOOS% %GOARCH%)
go.exe build -o %OUT%.exe -ldflags "-X 'main.APPLICATION_VERSION=%VERSION%'" .\cmd\protocol\.
GOTO EXIT

:CLEAN
del %OUT%
GOTO EXIT

:HELP

SET PROG=build.bat

@echo %PROG%
@echo %PROG% amd64  

@echo %PROG% 386

@echo %PROG% clean

:EXIT
