package main

import (
	"database/sql"
	"log"

	"github.com/Kazukite12/StruckShare/api"
	db "github.com/Kazukite12/StruckShare/db/sqlc"
	"github.com/Kazukite12/StruckShare/util"
	_ "github.com/lib/pq"
)

func main() {

	config, err := util.LoadConfig(".")

	conn, err := sql.Open(config.DBDriver, config.DBSource)

	if err != nil {
		log.Fatal("cannot connect to dabase", err)

	}

	store := db.NewStore(conn)

	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot start the server", err)
	}

	err = server.Start("0.0.0.0:8080")
	if err != nil {
		log.Fatal("cannot start the server", err)
	}

}
