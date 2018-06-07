package main

import (
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/luizalabs/go-training/services/routes"
	"github.com/luizalabs/go-training/util"
)

func main() {
	logger := util.GetLogger().WithFields(map[string]interface{}{
		"module":        "main",
		"function_name": "main",
	})

	dsn := "root:root@tcp(localhost:3306)/optimusprime"

	db, err := util.OpenMySQLConnection(dsn)
	if err != nil {
		logger.WithField("error", err.Error()).Error()
		os.Exit(1)
	}

	service := routes.New(db)
	routes, err := service.List(50, 0)
	if err != nil {
		logger.WithField("error", err.Error()).Error()
		os.Exit(1)
	}

	for i := range routes {
		fmt.Println(routes[i].ID)
	}
}
