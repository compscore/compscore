// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"github.com/99designs/gqlgen/graphql"
	"github.com/compscore/compscore/pkg/ent/check"
	"github.com/compscore/compscore/pkg/ent/credential"
	"github.com/compscore/compscore/pkg/ent/round"
	"github.com/compscore/compscore/pkg/ent/score"
	"github.com/compscore/compscore/pkg/ent/status"
	"github.com/compscore/compscore/pkg/ent/user"
)

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (c *CheckQuery) CollectFields(ctx context.Context, satisfies ...string) (*CheckQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return c, nil
	}
	if err := c.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return c, nil
}

func (c *CheckQuery) collectField(ctx context.Context, opCtx *graphql.OperationContext, collected graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	var (
		unknownSeen    bool
		fieldSeen      = make(map[string]struct{}, len(check.Columns))
		selectedFields = []string{check.FieldID}
	)
	for _, field := range graphql.CollectFields(opCtx, collected.Selections, satisfies) {
		switch field.Name {
		case "credential":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&CredentialClient{config: c.config}).Query()
			)
			if err := query.collectField(ctx, opCtx, field, path, satisfies...); err != nil {
				return err
			}
			c.WithNamedCredential(alias, func(wq *CredentialQuery) {
				*wq = *query
			})
		case "status":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&StatusClient{config: c.config}).Query()
			)
			if err := query.collectField(ctx, opCtx, field, path, satisfies...); err != nil {
				return err
			}
			c.WithNamedStatus(alias, func(wq *StatusQuery) {
				*wq = *query
			})
		case "name":
			if _, ok := fieldSeen[check.FieldName]; !ok {
				selectedFields = append(selectedFields, check.FieldName)
				fieldSeen[check.FieldName] = struct{}{}
			}
		case "weight":
			if _, ok := fieldSeen[check.FieldWeight]; !ok {
				selectedFields = append(selectedFields, check.FieldWeight)
				fieldSeen[check.FieldWeight] = struct{}{}
			}
		case "id":
		case "__typename":
		default:
			unknownSeen = true
		}
	}
	if !unknownSeen {
		c.Select(selectedFields...)
	}
	return nil
}

type checkPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []CheckPaginateOption
}

func newCheckPaginateArgs(rv map[string]any) *checkPaginateArgs {
	args := &checkPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (c *CredentialQuery) CollectFields(ctx context.Context, satisfies ...string) (*CredentialQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return c, nil
	}
	if err := c.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return c, nil
}

func (c *CredentialQuery) collectField(ctx context.Context, opCtx *graphql.OperationContext, collected graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	var (
		unknownSeen    bool
		fieldSeen      = make(map[string]struct{}, len(credential.Columns))
		selectedFields = []string{credential.FieldID}
	)
	for _, field := range graphql.CollectFields(opCtx, collected.Selections, satisfies) {
		switch field.Name {
		case "user":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&UserClient{config: c.config}).Query()
			)
			if err := query.collectField(ctx, opCtx, field, path, satisfies...); err != nil {
				return err
			}
			c.withUser = query
		case "check":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&CheckClient{config: c.config}).Query()
			)
			if err := query.collectField(ctx, opCtx, field, path, satisfies...); err != nil {
				return err
			}
			c.withCheck = query
		case "password":
			if _, ok := fieldSeen[credential.FieldPassword]; !ok {
				selectedFields = append(selectedFields, credential.FieldPassword)
				fieldSeen[credential.FieldPassword] = struct{}{}
			}
		case "id":
		case "__typename":
		default:
			unknownSeen = true
		}
	}
	if !unknownSeen {
		c.Select(selectedFields...)
	}
	return nil
}

type credentialPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []CredentialPaginateOption
}

func newCredentialPaginateArgs(rv map[string]any) *credentialPaginateArgs {
	args := &credentialPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (r *RoundQuery) CollectFields(ctx context.Context, satisfies ...string) (*RoundQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return r, nil
	}
	if err := r.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return r, nil
}

func (r *RoundQuery) collectField(ctx context.Context, opCtx *graphql.OperationContext, collected graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	var (
		unknownSeen    bool
		fieldSeen      = make(map[string]struct{}, len(round.Columns))
		selectedFields = []string{round.FieldID}
	)
	for _, field := range graphql.CollectFields(opCtx, collected.Selections, satisfies) {
		switch field.Name {
		case "status":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&StatusClient{config: r.config}).Query()
			)
			if err := query.collectField(ctx, opCtx, field, path, satisfies...); err != nil {
				return err
			}
			r.WithNamedStatus(alias, func(wq *StatusQuery) {
				*wq = *query
			})
		case "scores":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&ScoreClient{config: r.config}).Query()
			)
			if err := query.collectField(ctx, opCtx, field, path, satisfies...); err != nil {
				return err
			}
			r.WithNamedScores(alias, func(wq *ScoreQuery) {
				*wq = *query
			})
		case "completed":
			if _, ok := fieldSeen[round.FieldCompleted]; !ok {
				selectedFields = append(selectedFields, round.FieldCompleted)
				fieldSeen[round.FieldCompleted] = struct{}{}
			}
		case "id":
		case "__typename":
		default:
			unknownSeen = true
		}
	}
	if !unknownSeen {
		r.Select(selectedFields...)
	}
	return nil
}

type roundPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []RoundPaginateOption
}

func newRoundPaginateArgs(rv map[string]any) *roundPaginateArgs {
	args := &roundPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (s *ScoreQuery) CollectFields(ctx context.Context, satisfies ...string) (*ScoreQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return s, nil
	}
	if err := s.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return s, nil
}

func (s *ScoreQuery) collectField(ctx context.Context, opCtx *graphql.OperationContext, collected graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	var (
		unknownSeen    bool
		fieldSeen      = make(map[string]struct{}, len(score.Columns))
		selectedFields = []string{score.FieldID}
	)
	for _, field := range graphql.CollectFields(opCtx, collected.Selections, satisfies) {
		switch field.Name {
		case "round":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&RoundClient{config: s.config}).Query()
			)
			if err := query.collectField(ctx, opCtx, field, path, satisfies...); err != nil {
				return err
			}
			s.withRound = query
		case "user":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&UserClient{config: s.config}).Query()
			)
			if err := query.collectField(ctx, opCtx, field, path, satisfies...); err != nil {
				return err
			}
			s.withUser = query
		case "score":
			if _, ok := fieldSeen[score.FieldScore]; !ok {
				selectedFields = append(selectedFields, score.FieldScore)
				fieldSeen[score.FieldScore] = struct{}{}
			}
		case "id":
		case "__typename":
		default:
			unknownSeen = true
		}
	}
	if !unknownSeen {
		s.Select(selectedFields...)
	}
	return nil
}

type scorePaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []ScorePaginateOption
}

func newScorePaginateArgs(rv map[string]any) *scorePaginateArgs {
	args := &scorePaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (s *StatusQuery) CollectFields(ctx context.Context, satisfies ...string) (*StatusQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return s, nil
	}
	if err := s.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return s, nil
}

func (s *StatusQuery) collectField(ctx context.Context, opCtx *graphql.OperationContext, collected graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	var (
		unknownSeen    bool
		fieldSeen      = make(map[string]struct{}, len(status.Columns))
		selectedFields = []string{status.FieldID}
	)
	for _, field := range graphql.CollectFields(opCtx, collected.Selections, satisfies) {
		switch field.Name {
		case "round":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&RoundClient{config: s.config}).Query()
			)
			if err := query.collectField(ctx, opCtx, field, path, satisfies...); err != nil {
				return err
			}
			s.withRound = query
		case "check":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&CheckClient{config: s.config}).Query()
			)
			if err := query.collectField(ctx, opCtx, field, path, satisfies...); err != nil {
				return err
			}
			s.withCheck = query
		case "user":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&UserClient{config: s.config}).Query()
			)
			if err := query.collectField(ctx, opCtx, field, path, satisfies...); err != nil {
				return err
			}
			s.withUser = query
		case "status":
			if _, ok := fieldSeen[status.FieldStatus]; !ok {
				selectedFields = append(selectedFields, status.FieldStatus)
				fieldSeen[status.FieldStatus] = struct{}{}
			}
		case "message":
			if _, ok := fieldSeen[status.FieldMessage]; !ok {
				selectedFields = append(selectedFields, status.FieldMessage)
				fieldSeen[status.FieldMessage] = struct{}{}
			}
		case "timestamp":
			if _, ok := fieldSeen[status.FieldTimestamp]; !ok {
				selectedFields = append(selectedFields, status.FieldTimestamp)
				fieldSeen[status.FieldTimestamp] = struct{}{}
			}
		case "points":
			if _, ok := fieldSeen[status.FieldPoints]; !ok {
				selectedFields = append(selectedFields, status.FieldPoints)
				fieldSeen[status.FieldPoints] = struct{}{}
			}
		case "id":
		case "__typename":
		default:
			unknownSeen = true
		}
	}
	if !unknownSeen {
		s.Select(selectedFields...)
	}
	return nil
}

type statusPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []StatusPaginateOption
}

func newStatusPaginateArgs(rv map[string]any) *statusPaginateArgs {
	args := &statusPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (u *UserQuery) CollectFields(ctx context.Context, satisfies ...string) (*UserQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return u, nil
	}
	if err := u.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return u, nil
}

func (u *UserQuery) collectField(ctx context.Context, opCtx *graphql.OperationContext, collected graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	var (
		unknownSeen    bool
		fieldSeen      = make(map[string]struct{}, len(user.Columns))
		selectedFields = []string{user.FieldID}
	)
	for _, field := range graphql.CollectFields(opCtx, collected.Selections, satisfies) {
		switch field.Name {
		case "credential":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&CredentialClient{config: u.config}).Query()
			)
			if err := query.collectField(ctx, opCtx, field, path, satisfies...); err != nil {
				return err
			}
			u.WithNamedCredential(alias, func(wq *CredentialQuery) {
				*wq = *query
			})
		case "status":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&StatusClient{config: u.config}).Query()
			)
			if err := query.collectField(ctx, opCtx, field, path, satisfies...); err != nil {
				return err
			}
			u.WithNamedStatus(alias, func(wq *StatusQuery) {
				*wq = *query
			})
		case "scores":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&ScoreClient{config: u.config}).Query()
			)
			if err := query.collectField(ctx, opCtx, field, path, satisfies...); err != nil {
				return err
			}
			u.WithNamedScores(alias, func(wq *ScoreQuery) {
				*wq = *query
			})
		case "name":
			if _, ok := fieldSeen[user.FieldName]; !ok {
				selectedFields = append(selectedFields, user.FieldName)
				fieldSeen[user.FieldName] = struct{}{}
			}
		case "teamNumber":
			if _, ok := fieldSeen[user.FieldTeamNumber]; !ok {
				selectedFields = append(selectedFields, user.FieldTeamNumber)
				fieldSeen[user.FieldTeamNumber] = struct{}{}
			}
		case "role":
			if _, ok := fieldSeen[user.FieldRole]; !ok {
				selectedFields = append(selectedFields, user.FieldRole)
				fieldSeen[user.FieldRole] = struct{}{}
			}
		case "id":
		case "__typename":
		default:
			unknownSeen = true
		}
	}
	if !unknownSeen {
		u.Select(selectedFields...)
	}
	return nil
}

type userPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []UserPaginateOption
}

func newUserPaginateArgs(rv map[string]any) *userPaginateArgs {
	args := &userPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	return args
}

const (
	afterField     = "after"
	firstField     = "first"
	beforeField    = "before"
	lastField      = "last"
	orderByField   = "orderBy"
	directionField = "direction"
	fieldField     = "field"
	whereField     = "where"
)

func fieldArgs(ctx context.Context, whereInput any, path ...string) map[string]any {
	field := collectedField(ctx, path...)
	if field == nil || field.Arguments == nil {
		return nil
	}
	oc := graphql.GetOperationContext(ctx)
	args := field.ArgumentMap(oc.Variables)
	return unmarshalArgs(ctx, whereInput, args)
}

// unmarshalArgs allows extracting the field arguments from their raw representation.
func unmarshalArgs(ctx context.Context, whereInput any, args map[string]any) map[string]any {
	for _, k := range []string{firstField, lastField} {
		v, ok := args[k]
		if !ok {
			continue
		}
		i, err := graphql.UnmarshalInt(v)
		if err == nil {
			args[k] = &i
		}
	}
	for _, k := range []string{beforeField, afterField} {
		v, ok := args[k]
		if !ok {
			continue
		}
		c := &Cursor{}
		if c.UnmarshalGQL(v) == nil {
			args[k] = c
		}
	}
	if v, ok := args[whereField]; ok && whereInput != nil {
		if err := graphql.UnmarshalInputFromContext(ctx, v, whereInput); err == nil {
			args[whereField] = whereInput
		}
	}

	return args
}

func limitRows(partitionBy string, limit int, orderBy ...sql.Querier) func(s *sql.Selector) {
	return func(s *sql.Selector) {
		d := sql.Dialect(s.Dialect())
		s.SetDistinct(false)
		with := d.With("src_query").
			As(s.Clone()).
			With("limited_query").
			As(
				d.Select("*").
					AppendSelectExprAs(
						sql.RowNumber().PartitionBy(partitionBy).OrderExpr(orderBy...),
						"row_number",
					).
					From(d.Table("src_query")),
			)
		t := d.Table("limited_query").As(s.TableName())
		*s = *d.Select(s.UnqualifiedColumns()...).
			From(t).
			Where(sql.LTE(t.C("row_number"), limit)).
			Prefix(with)
	}
}

// mayAddCondition appends another type condition to the satisfies list
// if condition is enabled (Node/Nodes) and it does not exist in the list.
func mayAddCondition(satisfies []string, typeCond string) []string {
	if len(satisfies) == 0 {
		return satisfies
	}
	for _, s := range satisfies {
		if typeCond == s {
			return satisfies
		}
	}
	return append(satisfies, typeCond)
}
