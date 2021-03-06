package main

import (
	"log"
	_ "os"

	"github.com/gin-gonic/gin"
	vertex "github.com/proxima-one/proxima-data-vertex/pkg/vertex"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	configFilePath := "./app-config.yml"
	dbConfigFilePath := "./database/db-config.yaml"

	applicationVertex, err := vertex.LoadDataVertex(configFilePath, dbConfigFilePath)
	if err != nil {
		log.Fatalf("Data vertex creation error: %v", err)
	}

	applicationVertex.StartVertexServer()
}
