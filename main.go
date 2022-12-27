package main

import (
    "database/sql"
    "fmt"
    "github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	relation_map := loadRelation()
    r := setupRelation(relation_map)
    r.Run(":8080")
}

func setupRelation(relation_map map[string]string) *gin.Engine {
	fmt.Println("comes")
    r := gin.Default()
	
	for strong, weak := range relation_map {
        fmt.Println("key: %s, value: %d\n", strong, weak)
    }
    
	r.GET("/janken/gu", func(c *gin.Context) {
        c.String(200, "choki")
    })
    r.GET("/janken/choki", func(c *gin.Context) {
        c.String(200, "pa")
    })
    r.GET("/janken/pa", func(c *gin.Context) {
        c.String(200, "gu")
    })
    r.GET("/janken", func(c *gin.Context) {
        c.String(200, "usage: /janken/[gu,choki,pa] to see relation")
    })
    return r
}

func loadRelation() map[string]string {

	//sql.Open("mysql", "user:password@host/dbname")
	db, err := sql.Open("mysql", "root:admin@tcp(localhost:63306)/janken")

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
