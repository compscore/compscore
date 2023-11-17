package data

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/compscore/compscore/pkg/config"
	"github.com/compscore/compscore/pkg/ent"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

var (
	Client *ent.Client
	Ctx    context.Context = context.Background()
)

func Init() {
	var (
		err error
		c   *ent.Client
	)
	if config.Production {
		c, err = ent.Open(
			"postgres",
			fmt.Sprintf(
				"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
				os.Getenv("POSTGRES_HOST"),
				os.Getenv("POSTGRES_PORT"),
				os.Getenv("POSTGRES_USER"),
				os.Getenv("POSTGRES_PASSWORD"),
				os.Getenv("POSTGRES_DB"),
			),
		)
	} else {
		c, err = ent.Open("sqlite3", "file:database.sqlite?_loc=auto&cache=shared&_fk=1")
	}

	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}

	Client = c

	// Run the auto migration tool.
	if err := Client.Schema.Create(Ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}
