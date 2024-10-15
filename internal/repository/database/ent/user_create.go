// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"samsamoohooh-go-api/internal/repository/database/ent/comment"
	"samsamoohooh-go-api/internal/repository/database/ent/group"
	"samsamoohooh-go-api/internal/repository/database/ent/post"
	"samsamoohooh-go-api/internal/repository/database/ent/topic"
	"samsamoohooh-go-api/internal/repository/database/ent/user"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserCreate is the builder for creating a User entity.
type UserCreate struct {
	config
	mutation *UserMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (uc *UserCreate) SetCreatedAt(t time.Time) *UserCreate {
	uc.mutation.SetCreatedAt(t)
	return uc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (uc *UserCreate) SetNillableCreatedAt(t *time.Time) *UserCreate {
	if t != nil {
		uc.SetCreatedAt(*t)
	}
	return uc
}

// SetUpdatedAt sets the "updated_at" field.
func (uc *UserCreate) SetUpdatedAt(t time.Time) *UserCreate {
	uc.mutation.SetUpdatedAt(t)
	return uc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (uc *UserCreate) SetNillableUpdatedAt(t *time.Time) *UserCreate {
	if t != nil {
		uc.SetUpdatedAt(*t)
	}
	return uc
}

// SetDeleteAt sets the "delete_at" field.
func (uc *UserCreate) SetDeleteAt(t time.Time) *UserCreate {
	uc.mutation.SetDeleteAt(t)
	return uc
}

// SetNillableDeleteAt sets the "delete_at" field if the given value is not nil.
func (uc *UserCreate) SetNillableDeleteAt(t *time.Time) *UserCreate {
	if t != nil {
		uc.SetDeleteAt(*t)
	}
	return uc
}

// SetName sets the "name" field.
func (uc *UserCreate) SetName(s string) *UserCreate {
	uc.mutation.SetName(s)
	return uc
}

// SetResolution sets the "resolution" field.
func (uc *UserCreate) SetResolution(s string) *UserCreate {
	uc.mutation.SetResolution(s)
	return uc
}

// SetRole sets the "role" field.
func (uc *UserCreate) SetRole(u user.Role) *UserCreate {
	uc.mutation.SetRole(u)
	return uc
}

// SetSocial sets the "social" field.
func (uc *UserCreate) SetSocial(u user.Social) *UserCreate {
	uc.mutation.SetSocial(u)
	return uc
}

// SetSocialSub sets the "socialSub" field.
func (uc *UserCreate) SetSocialSub(s string) *UserCreate {
	uc.mutation.SetSocialSub(s)
	return uc
}

// AddGroupIDs adds the "groups" edge to the Group entity by IDs.
func (uc *UserCreate) AddGroupIDs(ids ...int) *UserCreate {
	uc.mutation.AddGroupIDs(ids...)
	return uc
}

// AddGroups adds the "groups" edges to the Group entity.
func (uc *UserCreate) AddGroups(g ...*Group) *UserCreate {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return uc.AddGroupIDs(ids...)
}

// AddCommentIDs adds the "comments" edge to the Comment entity by IDs.
func (uc *UserCreate) AddCommentIDs(ids ...int) *UserCreate {
	uc.mutation.AddCommentIDs(ids...)
	return uc
}

// AddComments adds the "comments" edges to the Comment entity.
func (uc *UserCreate) AddComments(c ...*Comment) *UserCreate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return uc.AddCommentIDs(ids...)
}

// AddPostIDs adds the "posts" edge to the Post entity by IDs.
func (uc *UserCreate) AddPostIDs(ids ...int) *UserCreate {
	uc.mutation.AddPostIDs(ids...)
	return uc
}

// AddPosts adds the "posts" edges to the Post entity.
func (uc *UserCreate) AddPosts(p ...*Post) *UserCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return uc.AddPostIDs(ids...)
}

// AddTopicIDs adds the "topics" edge to the Topic entity by IDs.
func (uc *UserCreate) AddTopicIDs(ids ...int) *UserCreate {
	uc.mutation.AddTopicIDs(ids...)
	return uc
}

// AddTopics adds the "topics" edges to the Topic entity.
func (uc *UserCreate) AddTopics(t ...*Topic) *UserCreate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return uc.AddTopicIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uc *UserCreate) Mutation() *UserMutation {
	return uc.mutation
}

// Save creates the User in the database.
func (uc *UserCreate) Save(ctx context.Context) (*User, error) {
	if err := uc.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, uc.sqlSave, uc.mutation, uc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (uc *UserCreate) SaveX(ctx context.Context) *User {
	v, err := uc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (uc *UserCreate) Exec(ctx context.Context) error {
	_, err := uc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uc *UserCreate) ExecX(ctx context.Context) {
	if err := uc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uc *UserCreate) defaults() error {
	if _, ok := uc.mutation.CreatedAt(); !ok {
		if user.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized user.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := user.DefaultCreatedAt()
		uc.mutation.SetCreatedAt(v)
	}
	if _, ok := uc.mutation.UpdatedAt(); !ok {
		if user.DefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized user.DefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := user.DefaultUpdatedAt()
		uc.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (uc *UserCreate) check() error {
	if _, ok := uc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "User.created_at"`)}
	}
	if _, ok := uc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "User.updated_at"`)}
	}
	if _, ok := uc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "User.name"`)}
	}
	if _, ok := uc.mutation.Resolution(); !ok {
		return &ValidationError{Name: "resolution", err: errors.New(`ent: missing required field "User.resolution"`)}
	}
	if _, ok := uc.mutation.Role(); !ok {
		return &ValidationError{Name: "role", err: errors.New(`ent: missing required field "User.role"`)}
	}
	if v, ok := uc.mutation.Role(); ok {
		if err := user.RoleValidator(v); err != nil {
			return &ValidationError{Name: "role", err: fmt.Errorf(`ent: validator failed for field "User.role": %w`, err)}
		}
	}
	if _, ok := uc.mutation.Social(); !ok {
		return &ValidationError{Name: "social", err: errors.New(`ent: missing required field "User.social"`)}
	}
	if v, ok := uc.mutation.Social(); ok {
		if err := user.SocialValidator(v); err != nil {
			return &ValidationError{Name: "social", err: fmt.Errorf(`ent: validator failed for field "User.social": %w`, err)}
		}
	}
	if _, ok := uc.mutation.SocialSub(); !ok {
		return &ValidationError{Name: "socialSub", err: errors.New(`ent: missing required field "User.socialSub"`)}
	}
	return nil
}

func (uc *UserCreate) sqlSave(ctx context.Context) (*User, error) {
	if err := uc.check(); err != nil {
		return nil, err
	}
	_node, _spec := uc.createSpec()
	if err := sqlgraph.CreateNode(ctx, uc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	uc.mutation.id = &_node.ID
	uc.mutation.done = true
	return _node, nil
}

func (uc *UserCreate) createSpec() (*User, *sqlgraph.CreateSpec) {
	var (
		_node = &User{config: uc.config}
		_spec = sqlgraph.NewCreateSpec(user.Table, sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt))
	)
	if value, ok := uc.mutation.CreatedAt(); ok {
		_spec.SetField(user.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := uc.mutation.UpdatedAt(); ok {
		_spec.SetField(user.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := uc.mutation.DeleteAt(); ok {
		_spec.SetField(user.FieldDeleteAt, field.TypeTime, value)
		_node.DeleteAt = value
	}
	if value, ok := uc.mutation.Name(); ok {
		_spec.SetField(user.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := uc.mutation.Resolution(); ok {
		_spec.SetField(user.FieldResolution, field.TypeString, value)
		_node.Resolution = value
	}
	if value, ok := uc.mutation.Role(); ok {
		_spec.SetField(user.FieldRole, field.TypeEnum, value)
		_node.Role = value
	}
	if value, ok := uc.mutation.Social(); ok {
		_spec.SetField(user.FieldSocial, field.TypeEnum, value)
		_node.Social = value
	}
	if value, ok := uc.mutation.SocialSub(); ok {
		_spec.SetField(user.FieldSocialSub, field.TypeString, value)
		_node.SocialSub = value
	}
	if nodes := uc.mutation.GroupsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.GroupsTable,
			Columns: user.GroupsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(group.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.CommentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.CommentsTable,
			Columns: []string{user.CommentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(comment.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.PostsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.PostsTable,
			Columns: []string{user.PostsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(post.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.TopicsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.TopicsTable,
			Columns: []string{user.TopicsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(topic.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// UserCreateBulk is the builder for creating many User entities in bulk.
type UserCreateBulk struct {
	config
	err      error
	builders []*UserCreate
}

// Save creates the User entities in the database.
func (ucb *UserCreateBulk) Save(ctx context.Context) ([]*User, error) {
	if ucb.err != nil {
		return nil, ucb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(ucb.builders))
	nodes := make([]*User, len(ucb.builders))
	mutators := make([]Mutator, len(ucb.builders))
	for i := range ucb.builders {
		func(i int, root context.Context) {
			builder := ucb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UserMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ucb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ucb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ucb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ucb *UserCreateBulk) SaveX(ctx context.Context) []*User {
	v, err := ucb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ucb *UserCreateBulk) Exec(ctx context.Context) error {
	_, err := ucb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ucb *UserCreateBulk) ExecX(ctx context.Context) {
	if err := ucb.Exec(ctx); err != nil {
		panic(err)
	}
}
