proto:
	protoc pkg/pb/*.proto --go_out=plugins=grpc:.

server:
	go run cmd/main.go
	
start:
	sudo docker compose down --remove-orphans -v && sudo docker compose up --build -d && lazydocker 
