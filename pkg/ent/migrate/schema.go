// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// ChecksColumns holds the columns for the "checks" table.
	ChecksColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString, Unique: true},
	}
	// ChecksTable holds the schema information for the "checks" table.
	ChecksTable = &schema.Table{
		Name:       "checks",
		Columns:    ChecksColumns,
		PrimaryKey: []*schema.Column{ChecksColumns[0]},
	}
	// RoundsColumns holds the columns for the "rounds" table.
	RoundsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "round", Type: field.TypeInt, Unique: true},
	}
	// RoundsTable holds the schema information for the "rounds" table.
	RoundsTable = &schema.Table{
		Name:       "rounds",
		Columns:    RoundsColumns,
		PrimaryKey: []*schema.Column{RoundsColumns[0]},
	}
	// StatusColumns holds the columns for the "status" table.
	StatusColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "error", Type: field.TypeString, Nullable: true},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"up", "down", "unknown"}, Default: "unknown"},
		{Name: "time", Type: field.TypeTime},
		{Name: "status_check", Type: field.TypeInt},
		{Name: "status_team", Type: field.TypeInt},
		{Name: "status_round", Type: field.TypeInt},
	}
	// StatusTable holds the schema information for the "status" table.
	StatusTable = &schema.Table{
		Name:       "status",
		Columns:    StatusColumns,
		PrimaryKey: []*schema.Column{StatusColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "status_checks_check",
				Columns:    []*schema.Column{StatusColumns[4]},
				RefColumns: []*schema.Column{ChecksColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "status_teams_team",
				Columns:    []*schema.Column{StatusColumns[5]},
				RefColumns: []*schema.Column{TeamsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "status_rounds_round",
				Columns:    []*schema.Column{StatusColumns[6]},
				RefColumns: []*schema.Column{RoundsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// TeamsColumns holds the columns for the "teams" table.
	TeamsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "number", Type: field.TypeInt8, Unique: true},
		{Name: "name", Type: field.TypeString, Unique: true},
	}
	// TeamsTable holds the schema information for the "teams" table.
	TeamsTable = &schema.Table{
		Name:       "teams",
		Columns:    TeamsColumns,
		PrimaryKey: []*schema.Column{TeamsColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		ChecksTable,
		RoundsTable,
		StatusTable,
		TeamsTable,
	}
)

func init() {
	StatusTable.ForeignKeys[0].RefTable = ChecksTable
	StatusTable.ForeignKeys[1].RefTable = TeamsTable
	StatusTable.ForeignKeys[2].RefTable = RoundsTable
}
