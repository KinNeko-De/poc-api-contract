#!/bin/bash

cd ..

export PATH="$PATH:$(go env GOPATH)/bin"

protobase=protos
protopath=${protobase}/kinnekode
projects=( $(ls ${protopath}) )

for i in "${projects[@]}"
do
    find ${protopath} -name *.proto \
	-exec tools/protoc/win64/bin/protoc.exe --proto_path=${protobase} --go_out=go --go_opt=paths=source_relative --go-grpc_out=go --go-grpc_opt=paths=source_relative {} +
	
	echo "Generated protos for: $i"
done

read -p "Press [Enter] to exit."