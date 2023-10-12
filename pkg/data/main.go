package data

import (
	"bytes"
	"context"
	"log"
	"sync"
	"text/template"

	"github.com/compscore/compscore/pkg/config"
	"github.com/compscore/compscore/pkg/ent"
	_ "github.com/mattn/go-sqlite3"
)

var (
	client *ent.Client
	ctx    context.Context = context.Background()
	mutex  *sync.Mutex     = &sync.Mutex{}
)

func Init() {
	c, err := ent.Open("sqlite3", "file:database.sqlite?_loc=auto&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}

	client = c

	// Run the auto migration tool.
	if err := c.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	teamNameTemplate, err := template.New("Name Template").Parse(config.Teams.NameFormat)
	if err != nil {
		log.Fatalf("failed parsing team name template: %v", err)
	}

	// Create teams if they do not exist
	for i := 1; i <= config.Teams.Amount; i++ {
		exists, err := Team.Exists(int8(i))
		if err != nil {
			log.Fatalf("failed checking for team %d: %v", i, err)
		}
		output := bytes.NewBuffer([]byte{})
		teamNameTemplate.Execute(output, struct{ Team int }{Team: i})
		if !exists {
			_, err := Team.Create(int8(i), output.String(), config.Teams.Password)
			if err != nil {
				log.Fatalf("failed creating team %d: %v", i, err)
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
		for i := 1; i <= config.Teams.Amount; i++ {
			exists, err := Credential.Exists(int8(i), check.Name)
			if err != nil {
				log.Fatalf("failed checking for credential %d:%s: %v", i, check.Name, err)
			}

			if !exists {
				_, err := Credential.Create(int8(i), check.Name, check.Credentials.Password)
				if err != nil {
					log.Fatalf("failed creating credential %d:%s: %v", i, check.Name, err)
				}
			}
		}
	}
}

func Client(function func(*ent.Client) (interface{}, error)) (interface{}, error) {
	mutex.Lock()
	defer mutex.Unlock()

	return function(client)
}
