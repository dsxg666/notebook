package routers

import (
	"github.com/dsxg666/notebook/internal/controllers"
	"github.com/gin-gonic/gin"
)

func ShowRouterInit(r *gin.Engine) {
	showRouter := r.Group("/")
	{
		showRouter.GET("/", controllers.ShowController{}.Home)
		showRouter.GET("/contact", controllers.ShowController{}.Contact)
		showRouter.GET("/intro", controllers.ShowController{}.Intro)
		showRouter.GET("/category/:id", controllers.ShowController{}.Category)
		showRouter.GET("/article/:id", controllers.ShowController{}.Article)
		showRouter.POST("/sendMsg", controllers.ShowController{}.SendMsg)
		showRouter.POST("/search", controllers.ShowController{}.Search)
	}
}
