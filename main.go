package main

import (
        // book adalah directory root project go yang kita buat
    "simpleCRUD/models" // memanggil package models pada directory models
    "simpleCRUD/routes"
    "github.com/labstack/echo"
    "github.com/labstack/echo/middleware"
)

func main() {
    e := echo.New()
    e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
        AllowOrigins: []string{"*"},
        AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
      }))
      
    db := models.SetupDB()
    db.AutoMigrate(&models.Task{})

    r := routes.SetupRoutes(db)
    r.Run()
}