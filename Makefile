MIGRATION_NAME = CreateUserTable
VERSION = 1;
run: 
	go run .\main.go 

migration_up: 
	migrate -path database/migrations/ -database "postgresql://postgres:root@localhost:5432/finance?sslmode=disable" -verbose up

migration_down: 
	migrate -path database/migrations/ -database "postgresql://postgres:root@localhost:5432/finance?sslmode=disable" -verbose down

migration_fix: 
	migrate -path database/migrations/ -database "postgresql://postgres:root@localhost:5432/finance?sslmode=disable" force 1

migration_create: 
	migrate create -ext sql -dir database/migrations/ -seq $(MIGRATION_NAME)