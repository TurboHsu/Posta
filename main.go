package main

import (
	"fmt"

	"github.com/TurboHsu/Posta/config"
	"github.com/TurboHsu/Posta/sql"
)

func main() {
	config.Read("config.toml")
	fmt.Println(config.C)
	sql.Init()
	r := setupRouter()
	r.Run(config.C.WebEngine.ListenAddr)
}
