package controllers

import (
    "github.com/gin-gonic/gin"
    "go-memo-api/models"
    "gorm.io/gorm"
    "net/http"
)

func CreateNoteHandler(c *gin.Context, db *gorm.DB) {
    var note models.Note
    if err := c.ShouldBindJSON(&note); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    db.Create(&note)
    c.JSON(http.StatusCreated, note)
}

func GetNotesHandler(c *gin.Context, db *gorm.DB) {
    var notes []models.Note
    db.Find(&notes)
    c.JSON(http.StatusOK, notes)
}

func GetNoteHandler(c *gin.Context, db *gorm.DB) {
    var note models.Note
    id := c.Param("id")
    if err := db.First(&note, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Note not found"})
        return
    }
    c.JSON(http.StatusOK, note)
}

func UpdateNoteHandler(c *gin.Context, db *gorm.DB) {
    var note models.Note
    id := c.Param("id")
    if err := db.First(&note, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Note not found"})
        return
    }
    if err := c.ShouldBindJSON(&note); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    db.Save(&note)
    c.JSON(http.StatusOK, note)
}

func DeleteNoteHandler(c *gin.Context, db *gorm.DB) {
    var note models.Note
    id := c.Param("id")
    if err := db.First(&note, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Note not found"})
        return
    }
    db.Delete(&note)
    c.JSON(http.StatusOK, gin.H{"message": "Note deleted successfully"})
}
