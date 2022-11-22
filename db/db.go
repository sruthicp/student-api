package db

import (
	"context"
	"fmt"
	"log"
	"student-api/config"

	"github.com/go-pg/pg"
	_ "github.com/lib/pq"
)

type dbLogger struct{}

func (d dbLogger) BeforeQuery(c *pg.QueryEvent) {
}

func (d dbLogger) AfterQuery(q *pg.QueryEvent) {
	fmt.Println("-----------------------------------******-------------------------------------------------------------------------------")
	fmt.Println(q.FormattedQuery())
	fmt.Println("-----------------------------------******-------------------------------------------------------------------------------")
}

// NewDBConnection connects to the DB
func NewDBConnection(conf *config.ServiceConfig) (conn *pg.DB, err error) {
	pgConf := conf.Postgres
	client := pg.Connect(&pg.Options{
		User:     "postgres",
		Password: "postgres",
		Database: pgConf.DBName,
		Addr:     fmt.Sprintf("%v:%v", pgConf.Host, pgConf.Port),
	})

	ctx := context.Background()

	_, err = client.ExecContext(ctx, "SELECT 1")
	if err != nil {
		log.Println("did not connect to postgres: ", err)
	}

	client.AddQueryHook(dbLogger{})
	log.Println("connected to postgres server....................")

	return client, err

	// var db_host string
	// conn = new(DBConnection)
	// conn.dbType = "postgres"
	// if os.Getenv("DB_HOST") == "" {
	// 	db_host = "localhost"
	// } else {
	// 	db_host = os.Getenv("DB_HOST")
	// }

	// conn.DB, err = sql.Open("postgres", fmt.Sprintf("postgres://postgres:postgres@%s:5432/postgres?sslmode=disable", db_host))
	// if err != nil {
	// 	return
	// }
	// err = conn.DB.Ping()
	// if err != nil {
	// 	return
	// }
	// conn.IDMax = 9223372036854775807

	// conn.Builder = sqlbuilder.PostgreSQL
	// return
}
