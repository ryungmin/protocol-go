# Why this proejct?
Protocol is a command-line tool to display ASCII RFC-like protocol header diagrams for both existing network protocols or user-defined ones (using a very simple syntax). Protocol is written in Python and it's open-source software, licensed under the GPLv3 license.

[Protocol](https://www.luismg.com/protocol/)

[A Simple ASCII Header Generator for Network Protocols](https://github.com/luismartingarcia/protocol)

This project re-implements protocol in go-lang. Compiled protocol is very suitable for environments where python is not installed.

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
> SET GOOS=windows & SET GOARCH=amd64 & go build -o protocol.exe .\cmd\protocol\. 
```

## macos
```
# build for Intel MAC(x64)
$ GOOS=darwin GOARCH=amd64 $(GO) build -o protocol_amd64 ./cmd/protocol/.

# build for Apple Silicon(arm64)
$ GOOS=darwin GOARCH=arm64 $(GO) build -o protocol_arm64 ./cmd/protocol/.

# Create a universal file
$ lipo -create -output protocol protocol_amd64 protocol_arm64
```

