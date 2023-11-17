package data

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/compscore/compscore/pkg/config"
	"github.com/compscore/compscore/pkg/ent"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

var (
	client *ent.Client
	ctx    context.Context = context.Background()
	mutex  *sync.Mutex     = &sync.Mutex{}
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

	client = c

	// Run the auto migration tool.
	if err := c.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}

func Client(function func(*ent.Client, context.Context) (interface{}, error)) (interface{}, error) {
	if !config.Production {
		mutex.Lock()
		defer mutex.Unlock()
	}

	return function(client, ctx)
}
