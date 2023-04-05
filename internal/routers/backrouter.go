package routers

import (
	"github.com/dsxg666/notebook/internal/controllers"
	"github.com/gin-gonic/gin"
)

func BackRouterInit(r *gin.Engine) {
	backRouter := r.Group("/")
	{
		backRouter.GET("/login", controllers.BackController{}.Login)
		backRouter.GET("/admin", controllers.BackController{}.Admin)
		backRouter.POST("/login", controllers.BackController{}.LoginP)
		backRouter.POST("/editor", controllers.BackController{}.Editor)
		backRouter.POST("/addArticle", controllers.BackController{}.AddArticle)
		backRouter.POST("/modifyArticle", controllers.BackController{}.ModifyArticle)
		backRouter.POST("/sendPicture", controllers.BackController{}.SendPicture)
	}
}
