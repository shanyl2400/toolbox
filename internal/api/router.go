package api

import "github.com/gin-gonic/gin"

func route(e *gin.Engine) {
	_api := newAPI()

	e.Use(_api.corsMiddleware())
	users := e.Group("/users")
	{
		//users.POST("/entities", _api.register)
		users.POST("/token", _api.login)
		users.PUT("/entities/:name/password", _api.checkAuth, _api.updatePassword)
		users.GET("/entities/:name", _api.checkAuth, _api.getUser)
		users.GET("/entities", _api.checkAuth, _api.listUsers)
	}

	tools := e.Group("/tools")
	{
		tools.POST("", _api.checkAuth, _api.createTool)
		tools.PUT("/:id", _api.checkAuth, _api.updateTool)
		tools.DELETE("/:id", _api.checkAuth, _api.deleteTool)
		tools.GET("/:id", _api.getTool)
		tools.GET("", _api.listTools)
	}

	columns := e.Group("/columns")
	{
		columns.PUT("/:name", _api.checkAuth, _api.putColumn)
		columns.DELETE("/:name", _api.checkAuth, _api.deleteColumn)
		columns.GET("", _api.listColumns)
	}

	environments := e.Group("/environments")
	{
		environments.GET("", _api.listEnvironments)
	}

	profiles := e.Group("/profiles")
	{
		profiles.POST("", _api.checkAuth, _api.uploadProfile)
		profiles.GET("/:name", _api.getProfile)
	}
}
