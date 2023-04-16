https://developers.google.com/protocol-buffers/docs/gotutorial

https://github.com/pseudomuto/protoc-gen-doc

https://github.com/envoyproxy/protoc-gen-validate

```` go
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go get -u github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc
go get -u github.com/envoyproxy/protoc-gen-validate
````
> 如果上面应用包不生效， 可以直接去github上下载对应的release可以执行文件，将对应文件（.exe）路径设置环境变量，后可正常运行 gen.bat