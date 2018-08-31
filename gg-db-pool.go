package gg-db-pool

import (
	"time"
	"database/sql"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func Database(connection string) gin.HandlerFunc {

	db, err := sql.Open("mysql", connection)
	if err != nil {
		panic(err)
	}

	db.SetConnMaxLifetime(time.Minute * 5)
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(50)

	return func(ctx *gin.Context) {

		ctx.Set("Database", db)
		ctx.Next()
	}
}
