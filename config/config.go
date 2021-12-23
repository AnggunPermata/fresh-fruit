package config


import (
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"os"
	"strconv"
)

var PORT int

func LoadEnv(key string) string{
	err := godotenv.Load(".env")
	if err != nil {
		return "error fetching environment variable"
	}
	return os.Getenv(key)
}

func CORSWithConfig(e *echo.Echo){
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowCredentials: true,
		AllowMethods:     []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
		AllowOrigins:     []string{"*"},
	}))
}

func InitPort() {
	PORT, _ = strconv.Atoi(LoadEnv("PORT"))
}