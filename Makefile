run-gateway:
	@go build -C ./gateway -o ./bin/main.exe

	@./gateway/bin/main.exe

run-orders:
	@go build -C ./orders -o ./bin/main.exe

	@./orders/bin/main.exe

run:
	@go build -C ./orders -o ./bin/main.exe
	@go build -C ./gateway -o ./bin/main.exe
	@./gateway/bin/main.exe
	@./orders/bin/main.exe

gen-proto:
	@cd ./common && protoc --go_out=. --go_opt=paths=source_relative \
         --go-grpc_out=. --go-grpc_opt=paths=source_relative \
         api/order-proccesing.proto
