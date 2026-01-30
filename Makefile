all:
	go install .


IO=$(shell find proto/io -name "*.proto")
CEL=$(shell find proto/cel -name "*.proto")
APIS=${IO} 

regenproto:
	rm -rf genproto
	make genproto

genproto: 
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	mkdir -p genproto
	protoc ${APIS} \
	--proto_path='proto' \
	--go_out='genproto'
