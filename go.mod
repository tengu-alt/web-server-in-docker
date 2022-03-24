module web-server-in-docker

go 1.17

require (
	github.com/golang-jwt/jwt v3.2.2+incompatible
	github.com/jackc/pgx/v4 v4.14.1
	github.com/jmoiron/sqlx v1.3.4
	github.com/lib/pq v1.10.4
	golang.org/x/crypto v0.0.0-20210711020723-a769d52b0f97
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b
)

require github.com/go-sql-driver/mysql v1.6.0 // indirect
