package dbrepo

import (
	"database/sql"

	"github.com/cale-i/building-modern-web-applications-with-go-bookings-project/internal/config"
	"github.com/cale-i/building-modern-web-applications-with-go-bookings-project/internal/repository"
)

type postgresDBRepo struct {
	App *config.AppConfig
	DB  *sql.DB
}

func NewPostgresRepo(conn *sql.DB, a *config.AppConfig) repository.DatabaseRepo {
	return &postgresDBRepo{
		App: a,
		DB:  conn,
	}
}

type testDBRepo struct {
	App *config.AppConfig
	DB  *sql.DB
}

func NewTestingRepo(a *config.AppConfig) repository.DatabaseRepo {
	return &testDBRepo{
		App: a,
	}
}

// sample for MariaDB
// type mariaDBRepo struct {
// 	App *config.AppConfig
// 	DB  *sql.DB
// }

// func NewMariaRepo(conn *sql.DB, a *config.AppConfig) repository.DatabaseRepo {
// 	return &mariaDBRepo{
// 		App: a,
// 		DB:  conn,
// 	}
// }
