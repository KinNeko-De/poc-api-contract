#!/bin/bash

cd ..

export PATH="$PATH:$(go env GOPATH)/bin"

protobase=protos
protopath=${protobase}/kinnekode
projects=( $(command ls ${protopath}) )
outputpath=docs

for i in "${projects[@]}"
do
    projectoutputpath=${outputpath}/${i}
    mkdir ${projectoutputpath}

    find ${protopath}/${i} ${protopath}/protobuf ${protobase}/google -name *.proto \
    -exec tools/protoc/win64/bin/protoc.exe --proto_path=${protobase} --doc_out=${projectoutputpath} --doc_opt=markdown,index.md {} +

    echo "Generated doc for: $i"
done

read -p "Press [Enter] to exit."