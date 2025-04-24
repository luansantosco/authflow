package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/luansantosco/authflow/config"
)

func main() {
	gin.SetMode(gin.Mode())
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erro ao carrgar o arquivo .env")
	}

	cfg := config.Load()

	router := gin.Default()
	router.SetTrustedProxies([]string{"127.0.01"})

	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	err = router.Run(":" + cfg.Port)
	if err != nil {
		log.Fatal("Error ao iniciar servidor: %v", err)
	}

	fmt.Println("Servidor rodando na porta:", cfg.Port)
}
