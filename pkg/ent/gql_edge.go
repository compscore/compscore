// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
)

func (c *Check) Credential(ctx context.Context) (result []*Credential, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = c.NamedCredential(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = c.Edges.CredentialOrErr()
	}
	if IsNotLoaded(err) {
		result, err = c.QueryCredential().All(ctx)
	}
	return result, err
}

func (c *Check) Status(ctx context.Context) (result []*Status, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = c.NamedStatus(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = c.Edges.StatusOrErr()
	}
	if IsNotLoaded(err) {
		result, err = c.QueryStatus().All(ctx)
	}
	return result, err
}

func (c *Credential) User(ctx context.Context) (*User, error) {
	result, err := c.Edges.UserOrErr()
	if IsNotLoaded(err) {
		result, err = c.QueryUser().Only(ctx)
	}
	return result, err
}

func (c *Credential) Check(ctx context.Context) (*Check, error) {
	result, err := c.Edges.CheckOrErr()
	if IsNotLoaded(err) {
		result, err = c.QueryCheck().Only(ctx)
	}
	return result, err
}

func (r *Round) Status(ctx context.Context) (result []*Status, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = r.NamedStatus(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = r.Edges.StatusOrErr()
	}
	if IsNotLoaded(err) {
		result, err = r.QueryStatus().All(ctx)
	}
	return result, err
}

func (s *Status) Round(ctx context.Context) (*Round, error) {
	result, err := s.Edges.RoundOrErr()
	if IsNotLoaded(err) {
		result, err = s.QueryRound().Only(ctx)
	}
	return result, err
}

func (s *Status) Check(ctx context.Context) (*Check, error) {
	result, err := s.Edges.CheckOrErr()
	if IsNotLoaded(err) {
		result, err = s.QueryCheck().Only(ctx)
	}
	return result, err
}

func (s *Status) User(ctx context.Context) (*User, error) {
	result, err := s.Edges.UserOrErr()
	if IsNotLoaded(err) {
		result, err = s.QueryUser().Only(ctx)
	}
	return result, err
}

func (u *User) Credential(ctx context.Context) (result []*Credential, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = u.NamedCredential(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = u.Edges.CredentialOrErr()
	}
	if IsNotLoaded(err) {
		result, err = u.QueryCredential().All(ctx)
	}
	return result, err
}

func (u *User) Status(ctx context.Context) (result []*Status, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = u.NamedStatus(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = u.Edges.StatusOrErr()
	}
	if IsNotLoaded(err) {
		result, err = u.QueryStatus().All(ctx)
	}
	return result, err
}
