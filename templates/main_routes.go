package templates

var MainRoutes = `package routes

import (
    "github.com/gin-contrib/sessions"
    rsess "github.com/gin-contrib/sessions/redis"
    "github.com/gin-gonic/contrib/static"
    "github.com/gin-gonic/gin"
    redis "github.com/go-redis/redis/v7"
    "{{.GitModelsPath}}"
)

type Router struct {
    Model *models.Model
    Cache *redis.Client
}

func newRoutes(model *models.Model, cache *redis.Client) Router {
    return Router {
        Model: model,
        Cache: cache,
    }
}

func Routes(model *models.Model, client *redis.Client){
    r := gin.Default()
    // SESSIONS
    store, _ := rsess.NewStore(10, "tcp", "localhost:6379", "", []byte("secret"))
    sessionNames := []string{"A"}
    r.Use(sessions.SessionsMany(sessionNames, store))
    routes := newRoutes(model, client )
    routes.registerRoutes(r)
    r.GET("/footest", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "FOO",
        })
    })

    r.Use(static.Serve("/bar", static.LocalFile("./fbfa/public/", true)))
    r.Run(":3001")
}

`