package templates

var MainTemplate = `package main

import (
	"fmt"
	redis "github.com/go-redis/redis/v7"
	"{{.GitDbPath}}"
	"{{.GitModelsPath}}"
	"{{.GitRoutesPath}}"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func main() {
	client := initCache()
	db := db.SetupDb()
	model := models.NewModel(db)
	routes.Routes(&model, client)
}

func initCache( ) *redis.Client{
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
	return client
}
`

