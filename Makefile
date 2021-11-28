migrate:
	migrate -path migrations -database "postgres://postgres@localhost:5432/tracking?sslmode=disable" up