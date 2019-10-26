package templates

var RoutesTemplate = `package routes

import (
	"encoding/json"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/{{.Owner}}/{{.Project}}/internal/pkg/{{.Model}}"
	"log"
)

func (h *Routes) Attach{{.Model}}Routes(r *gin.Engine) {
	v1 := r.Group("/{{.PluralLowerModel}}")
	v1.GET("/", h.All{{.PluralModel}})
	v1.GET("/new", h.New{{.Model}})
	v1.PUT("/:id", h.Edit{{.Model}})
	v1.DELETE("/:id", h.Delete{{.Model}})
}`
