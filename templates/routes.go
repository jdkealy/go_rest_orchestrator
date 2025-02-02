package templates

var RoutesTemplate = `package routes

import (
	"github.com/gin-gonic/gin"
	"{{.GitModelsPath}}"
)

func (h *Router) Attach{{.Model}}Routes(r *gin.Engine) {
	v1 := r.Group("/{{.PluralLowerModel}}")
	v1.GET("/", h.All{{.PluralModel}})
	v1.POST("/", h.New{{.Model}})
	v1.PUT("/:id", h.Edit{{.Model}})
	v1.DELETE("/:id", h.Delete{{.Model}})
}

func (h *Router) Delete{{.Model}}(c *gin.Context) {
	res, err := h.Model.All{{.PluralModel}}()
	if err != nil {
		c.JSON(500, err)
	}
	c.JSON(200, res)
}

func (h *Router) Edit{{.Model}}(c *gin.Context) {
	res, err := h.Model.All{{.PluralModel}}()
	if err != nil {
		c.JSON(500, err)
	}
	c.JSON(200, res)
}

func (h *Router) All{{.PluralModel}}(c *gin.Context) {
	res, err := h.Model.All{{.PluralModel}}()
	if err != nil {
		c.JSON(500, err)
	}
	c.JSON(200, res)
}

func (h *Router) New{{.Model}}(c *gin.Context) {
	var {{.LowerModel}} models.{{.Model}}
	c.Bind(&{{.LowerModel}})
	res, err := h.Model.Create{{.Model}}({{.LowerModel}})
	if err != nil {
		c.JSON(500, err)
	}
	c.JSON(200, res)
}
`
