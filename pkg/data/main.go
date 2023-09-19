package data

import (
	"context"
	"log"

	"github.com/compscore/compscore/pkg/ent"
	_ "github.com/mattn/go-sqlite3"
)

var (
	Client *ent.Client
	Ctx    context.Context = context.Background()
)

func Init() {
	client, err := ent.Open("sqlite3", "file:database.sqlite?_loc=auto&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}

	Client = client

	// Run the auto migration tool.
	if err := Client.Schema.Create(Ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}
