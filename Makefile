# Makefile для создания миграций

# Переменные которые будут использоваться в командах (Таргетах)
DB_DSN := "postgres://postgres:spice515463@localhost:5432/postgres?sslmode=disable"
MIGRATE := migrate -path ./migrations -database $(DB_DSN)

# Таргет для создания новой миграции
migrate-new:
	migrate create -ext sql -dir ./migrations ${NAME}

# Применение миграций
migrate:
	$(MIGRATE) up

# Откат миграций
migrate-down:
	$(MIGRATE) down

# команда run, которая будет запускать приложение
run:
	go run cmd/app/main.go # Теперь при вызове make run мы запустим наш сервер