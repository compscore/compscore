package data

import (
	"context"

	"github.com/compscore/compscore/ent"
	"github.com/sirupsen/logrus"

	_ "github.com/mattn/go-sqlite3"
)

var (
	Client *ent.Client
	Ctx    context.Context = context.Background()

	Check    check_s    = check_s{}
	CheckLog checklog_s = checklog_s{}
	Round    round_s    = round_s{}
	Team     team_s     = team_s{}
)

func init() {
	client, err := ent.Open("sqlite3", "file:data.sqlite?_loc=auto&cache=shared&_fk=1")
	if err != nil {
		logrus.WithError(err).Fatal("Failed to open database connection")
	}
	Client = client

	if err := client.Schema.Create(Ctx); err != nil {
		logrus.WithError(err).Fatal("Failed to create schema")
	}

}
