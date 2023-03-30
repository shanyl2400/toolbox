package api

import (
	"fmt"
	"toolbox/internal/config"

	"github.com/gin-gonic/gin"
)

func Start() {
	r := gin.Default()
	route(r)
	r.Run(fmt.Sprintf(":%d", config.Get().HttpPort)) // listen and serve on 0.0.0.0:8080
}
