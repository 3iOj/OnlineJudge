package main

import (
	"database/sql"
	"fmt"

	"github.com/3iOj/OnlineJudge/api"
	db "github.com/3iOj/OnlineJudge/db/sqlc"
	util "github.com/3iOj/OnlineJudge/utils"
	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".")
	logger := util.GetLogger()
	if err != nil {
		logger.Fatal().Err(err).Msg("cannot load config")
	}
	conn, err := sql.Open(config.DBDriver, fmt.Sprintf("postgresql://%s:%s@%s:5432/%s?sslmode=disable", config.DBUser, config.DBPassword, config.DBHost, config.DBName))
	if err != nil {
		logger.Fatal().Err(err).Msg("cannot connect to db")
	} else {
		logger.Info().Msg("Connected to DB")
	}
	store := db.NewStore(conn)

	server, err := api.NewServer(config, store)
	if err != nil {
		logger.Fatal().Err(err).Msg("cannot create server")
	}
	server.Start(config.ServerPort)

	if err != nil {
		logger.Fatal().Err(err).Msg("cannot start server")
	}

}
