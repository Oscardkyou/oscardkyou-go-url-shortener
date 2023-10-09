
package routes

import (
    "github.com/gin-gonic/gin"
    "../handlers"
)

func SetUpURLRoutes(router *gin.Engine) {
    urlGroup := router.Group("/urls")
    {
        urlGroup.GET("/", handlers.GetAllURLs)
        // Further routes like POST, PUT, DELETE can be added here
    }
}
