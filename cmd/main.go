package main

import (
	"github.com/lautarok/yorcom/internal/infra"
)

func main() {
	infra.InitConfig()
	db := infra.InitDatabase()
	infra.InitApi(db)
}
