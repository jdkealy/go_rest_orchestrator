package templates

var RegisterRoutesTemplate = `package routes

import (
    "github.com/gin-gonic/gin"
)

func (router *Router) registerRoutes(c *gin.Engine ){
	/*REGISTER_ROUTES_GEN*/
}`
