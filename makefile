run:
	go run cmd/main/main.go

stan:
	docker run -d -p 4222:4222 -p 8222:8222 nats-streaming

psql:
	docker run --name=postgres -d -p 5432:5432 -e POSTGRES_PASSWORD=password postgres:latest

rmPsql:
	docker rm -f postgres

test:
	go test ./... -race

dbNext:
	migrate create -ext sql -dir ./migration -seq $(name)
dbUp:
	migrate -path ./migration -database "postgres://postgres:password@localhost:5432/postgres?sslmode=disable" -verbose up $(v)
dbDown:
	migrate -path ./migration -database "postgres://postgres:password@localhost:5432/postgres?sslmode=disable" -verbose down $(v)