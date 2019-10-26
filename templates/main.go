package templates

var MainTemplate = `
package main

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	rsess "github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	redis "github.com/go-redis/redis/v7"
	"github.com/jdkealy/fbfa/db"
	"github.com/jdkealy/fbfa/internal/pkg/users"
	myRoutes "github.com/jdkealy/fbfa/routes"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)`

