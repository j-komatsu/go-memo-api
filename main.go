package main

import (
    "github.com/gin-gonic/gin"
    "go-memo-api/config"
    "go-memo-api/routes"
    "github.com/gin-contrib/cors"
    "fmt"
    "time"
)

func main() {
    db := config.InitDB()
    router := gin.Default()

    // CORSのミドルウェアを追加
    router.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"*"}, // 必要に応じて変更
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
        AllowCredentials: true,
        MaxAge:           12 * time.Hour,
    }))

    routes.SetupRoutes(router, db)
    fmt.Println("Server is running on http://localhost:8080")
    router.Run(":8080")
}
