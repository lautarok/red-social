package infra

import (
	"context"
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/lautarok/yorcom/internal/domain"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mysqldialect"
)

func InitDatabase() *bun.DB {
	sqldb, err := sql.Open("mysql", os.Getenv("MYSQL_DSN"))
	if err != nil {
		log.Fatal(err)
	}

	db := bun.NewDB(sqldb, mysqldialect.New())
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	db.RegisterModel((*domain.UserConversation)(nil))
	db.NewCreateTable().IfNotExists().Model((*domain.User)(nil)).Exec(context.Background())
	db.NewCreateTable().IfNotExists().Model((*domain.Conversation)(nil)).Exec(context.Background())
	db.NewCreateTable().IfNotExists().Model((*domain.Message)(nil)).Exec(context.Background())

	return db
}
