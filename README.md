# Gin-Gonic Database Pool Middleware

Simple database pool middleware for Gin-Gonic framework.

This middleware set on each request a mysql connection (new or pooled), and can be accessed with context `Get("Database")` or `MustGet("Database")` method.

```
package main

import (
	ggDBPool "github.com/valenciaj/gg-db-pool"
)

// Create server
srv := gin.Default()

dbConnStr =: "mysql://user:pwd@tcp(localhost)/database"

// Set middlewares
srv.Use(
	gin.Recovery(),
	ggDBPool.Database(dbConnStr),
)

// ... set other stuff ...

// Set routes
srv.GET("/", func(ctx *gin.Context) {

	// With Get method
	dbI, err := ctx.Get("Database")
	if err != nil {
		panic(err)
	}
	// Unbox database connection
	db := dbI.(*DB)

	// With MustGet method
	db := ctx.MustGet("Database").(*DB)

	// ... do some stuff with database connection ...

	// Response data
	ctx.HTML(200, "welcome/index.tmpl", gin.H{})
})

```
