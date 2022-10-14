# Why this project? 
Protocol is a command-line tool to display ASCII RFC-like protocol header diagrams for both existing network protocols or user-defined ones (using a very simple syntax). Protocol is written in Python and it's open-source software, licensed under the GPLv3 license. 
 
This project re-implements protocol in go-lang. Compiled protocol is very suitable for environments where python is not installed. 
 
## Original Protocol (is written in Python) 
 
[Luis MartinGarcia's Protocol](https://www.luismg.com/protocol/) 
 
github: [A Simple ASCII Header Generator for Network Protocols](https://github.com/luismartingarcia/protocol) 
 
# Building the protocol-go 
``` 
$ make  
``` 
 
## linux 
``` 
$ GOOS=linux go build -o protocol ./cmd/protocol/. 
``` 
 
## windows 
``` 
> SET GOOS=windows & SET GOARCH=amd64 & go.exe build -o protocol.exe .\cmd\protocol\.  
``` 
 
## macos 
``` 
# build for Intel MAC(x64) 
$ GOOS=darwin GOARCH=amd64 go build -o protocol_amd64 ./cmd/protocol/. 
 
# build for Apple Silicon(arm64) 
$ GOOS=darwin GOARCH=arm64 go build -o protocol_arm64 ./cmd/protocol/. 
 
# Create an universal file 
$ lipo -create -output protocol protocol_amd64 protocol_arm64 
``` 
 
# How to use 
[Luis MartinGarcia's Examples](https://www.luismg.com/protocol/#05) 
 
``` 
$ protocol "Source:16,TTL:8,Reserved:40" 
 0                   1                   2                   3  
 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|            Source             |      TTL      |               |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+               +
|                           Reserved                            |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+

* Source (16 bytes)
* TTL (8 bytes)
* Reserved (40 bytes)
total 64 bytes
``` 
 
``` 
$ protocol "Source:16,Reserved:40,TTL:8" 
 0                   1                   2                   3  
 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|            Source             |                               |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+               +-+-+-+-+-+-+-+-+
|                   Reserved                    |      TTL      |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+

* Source (16 bytes)
* Reserved (40 bytes)
* TTL (8 bytes)
total 64 bytes
``` 
 
``` 
$ protocol "Reserved:32,Target Address:128" 
 0                   1                   2                   3  
 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|                           Reserved                            |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|                                                               |
+                                                               +
|                                                               |
+                        Target Address                         +
|                                                               |
+                                                               +
|                                                               |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+

* Reserved (32 bytes)
* Target Address (128 bytes)
total 160 bytes
``` 
 
