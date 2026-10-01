package main

import (
    "database/sql"
    "fmt"
    "os"
    "github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	relation_map := loadRelation()
    r := setupRelation(relation_map)
    r.Run(":8080")
}

func setupRelation(relation_map map[string]string) *gin.Engine {
    r := gin.Default()

	// Serve each hand from the loaded table, so the routes follow the DB
	// instead of a hard-coded copy of it.
	for strong, weak := range relation_map {
        fmt.Printf("key: %s, value: %s\n", strong, weak)
        weak := weak
        r.GET("/janken/"+strong, func(c *gin.Context) {
            c.String(200, weak)
        })
    }
    r.GET("/janken", func(c *gin.Context) {
        c.String(200, "usage: /janken/[gu,choki,pa] to see relation")
    })
    return r
}

func loadRelation() map[string]string {

	//sql.Open("mysql", "user:password@host/dbname")
	dsn := os.Getenv("JANKEN_DSN")
	if dsn == "" {
		dsn = "root:admin@tcp(localhost:63306)/janken"
	}
	db, err := sql.Open("mysql", dsn)

	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM janken_relation")

	if err != nil {
		fmt.Println("データベース接続失敗")
		panic(err.Error())
	} else {
		fmt.Println("データベース接続成功")
	}

	defer rows.Close()

    relation_map := make(map[string]string)
	for rows.Next() {

		var strong	string
		var weak	string
		
		err := rows.Scan(&strong, &weak)
		if err != nil {
			panic(err.Error())
		}

		fmt.Println(strong, weak)

		relation_map[strong] = weak
	}

	err = rows.Err()
	if err != nil {
		panic(err.Error())
	}

	return relation_map
}


func setupRouter() *gin.Engine {
    r := gin.Default()
    r.GET("/ping", func(c *gin.Context) {
        c.String(200, "pong")
    })
    return r
}
