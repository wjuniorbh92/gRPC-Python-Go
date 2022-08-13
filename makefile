proto_create_python:
	rm -rf ./client/stock_pb2_grpc.py ./client/stock_pb2.py
	python -m grpc_tools.protoc -I ./protos --python_out=./client/ --grpc_python_out=./client/ ./protos/stock.proto
	export PATH="$PATH:$(go env GOPATH)/bin" 
	protoc --go_out=./server --go_opt=paths=source_relative \
   --go-grpc_out=./server --go-grpc_opt=paths=source_relative \
   protos/stock.proto