package data

import (
	"context"
	"log"

	"github.com/compscore/compscore/pkg/config"
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

	// Create teams if they do not exist
	for _, team := range config.Teams {
		exists, err := Team.Exists(team.Number)
		if err != nil {
			log.Fatalf("failed checking for team %d: %v", team.Number, err)
		}

		if !exists {
			_, err := Team.Create(team.Number, team.Name, team.Password)
			if err != nil {
				log.Fatalf("failed creating team %d: %v", team.Number, err)
			}
		}
	}

	// Create checks if they do not exist
	for _, check := range config.Checks {
		exists, err := Check.Exists(check.Name)
		if err != nil {
			log.Fatalf("failed checking for check %s: %v", check.Name, err)
		}

		if !exists {
			_, err := Check.Create(check.Name)
			if err != nil {
				log.Fatalf("failed creating check %s: %v", check.Name, err)
			}
		}
	}

	// Create credentials if they do not exist
	for _, check := range config.Checks {
		for _, team := range config.Teams {
			exists, err := Credential.Exists(team.Number, check.Name)
			if err != nil {
				log.Fatalf("failed checking for credential %d:%s: %v", team.Number, check.Name, err)
			}

			if !exists {
				_, err := Credential.Create(team.Number, check.Name, team.Password)
				if err != nil {
					log.Fatalf("failed creating credential %d:%s: %v", team.Number, check.Name, err)
				}
			}
		}
	}
}
