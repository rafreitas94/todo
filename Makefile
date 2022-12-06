# sintaxe
# target:
#	comandos

build:
	go build -o todoserver cmd/todoserver/main.go

run:
	go run cmd/todoserver/main.go

migration-status:
	goose -dir ./migrations postgres "host=localhost user=user password=password dbname=todo sslmode=disable" status

migration-up:
	goose -dir ./migrations postgres "host=localhost user=user password=password dbname=todo sslmode=disable" up

migration-down:
	goose -dir ./migrations postgres "host=localhost user=user password=password dbname=todo sslmode=disable" down

migration-create:
	goose -dir ./migrations create ${NAME} sql
