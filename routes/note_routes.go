package routes

import (
    "github.com/gin-gonic/gin"
    "go-memo-api/controllers"
    "gorm.io/gorm"
)

func SetupRoutes(router *gin.Engine, db *gorm.DB) {
    router.POST("/notes", func(c *gin.Context) { controllers.CreateNoteHandler(c, db) })
    router.GET("/notes", func(c *gin.Context) { controllers.GetNotesHandler(c, db) })
    router.GET("/notes/:id", func(c *gin.Context) { controllers.GetNoteHandler(c, db) })
    router.PUT("/notes/:id", func(c *gin.Context) { controllers.UpdateNoteHandler(c, db) })
    router.DELETE("/notes/:id", func(c *gin.Context) { controllers.DeleteNoteHandler(c, db) })
}
