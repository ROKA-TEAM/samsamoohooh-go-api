// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"samsamoohooh-go-api/internal/repository/database/ent/comment"
	"samsamoohooh-go-api/internal/repository/database/ent/post"
	"samsamoohooh-go-api/internal/repository/database/ent/predicate"
	"samsamoohooh-go-api/internal/repository/database/ent/user"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CommentUpdate is the builder for updating Comment entities.
type CommentUpdate struct {
	config
	hooks    []Hook
	mutation *CommentMutation
}

// Where appends a list predicates to the CommentUpdate builder.
func (cu *CommentUpdate) Where(ps ...predicate.Comment) *CommentUpdate {
	cu.mutation.Where(ps...)
	return cu
}

// SetUpdatedAt sets the "updated_at" field.
func (cu *CommentUpdate) SetUpdatedAt(t time.Time) *CommentUpdate {
	cu.mutation.SetUpdatedAt(t)
	return cu
}

// SetDeleteAt sets the "delete_at" field.
func (cu *CommentUpdate) SetDeleteAt(t time.Time) *CommentUpdate {
	cu.mutation.SetDeleteAt(t)
	return cu
}

// SetNillableDeleteAt sets the "delete_at" field if the given value is not nil.
func (cu *CommentUpdate) SetNillableDeleteAt(t *time.Time) *CommentUpdate {
	if t != nil {
		cu.SetDeleteAt(*t)
	}
	return cu
}

// ClearDeleteAt clears the value of the "delete_at" field.
func (cu *CommentUpdate) ClearDeleteAt() *CommentUpdate {
	cu.mutation.ClearDeleteAt()
	return cu
}

// SetContent sets the "content" field.
func (cu *CommentUpdate) SetContent(s string) *CommentUpdate {
	cu.mutation.SetContent(s)
	return cu
}

// SetNillableContent sets the "content" field if the given value is not nil.
func (cu *CommentUpdate) SetNillableContent(s *string) *CommentUpdate {
	if s != nil {
		cu.SetContent(*s)
	}
	return cu
}

// SetUserID sets the "user" edge to the User entity by ID.
func (cu *CommentUpdate) SetUserID(id int) *CommentUpdate {
	cu.mutation.SetUserID(id)
	return cu
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (cu *CommentUpdate) SetNillableUserID(id *int) *CommentUpdate {
	if id != nil {
		cu = cu.SetUserID(*id)
	}
	return cu
}

// SetUser sets the "user" edge to the User entity.
func (cu *CommentUpdate) SetUser(u *User) *CommentUpdate {
	return cu.SetUserID(u.ID)
}

// SetPostID sets the "post" edge to the Post entity by ID.
func (cu *CommentUpdate) SetPostID(id int) *CommentUpdate {
	cu.mutation.SetPostID(id)
	return cu
}

// SetNillablePostID sets the "post" edge to the Post entity by ID if the given value is not nil.
func (cu *CommentUpdate) SetNillablePostID(id *int) *CommentUpdate {
	if id != nil {
		cu = cu.SetPostID(*id)
	}
	return cu
}

// SetPost sets the "post" edge to the Post entity.
func (cu *CommentUpdate) SetPost(p *Post) *CommentUpdate {
	return cu.SetPostID(p.ID)
}

// Mutation returns the CommentMutation object of the builder.
func (cu *CommentUpdate) Mutation() *CommentMutation {
	return cu.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (cu *CommentUpdate) ClearUser() *CommentUpdate {
	cu.mutation.ClearUser()
	return cu
}

// ClearPost clears the "post" edge to the Post entity.
func (cu *CommentUpdate) ClearPost() *CommentUpdate {
	cu.mutation.ClearPost()
	return cu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *CommentUpdate) Save(ctx context.Context) (int, error) {
	if err := cu.defaults(); err != nil {
		return 0, err
	}
	return withHooks(ctx, cu.sqlSave, cu.mutation, cu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cu *CommentUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *CommentUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *CommentUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cu *CommentUpdate) defaults() error {
	if _, ok := cu.mutation.UpdatedAt(); !ok {
		if comment.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized comment.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := comment.UpdateDefaultUpdatedAt()
		cu.mutation.SetUpdatedAt(v)
	}
	return nil
}

func (cu *CommentUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(comment.Table, comment.Columns, sqlgraph.NewFieldSpec(comment.FieldID, field.TypeInt))
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.UpdatedAt(); ok {
		_spec.SetField(comment.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := cu.mutation.DeleteAt(); ok {
		_spec.SetField(comment.FieldDeleteAt, field.TypeTime, value)
	}
	if cu.mutation.DeleteAtCleared() {
		_spec.ClearField(comment.FieldDeleteAt, field.TypeTime)
	}
	if value, ok := cu.mutation.Content(); ok {
		_spec.SetField(comment.FieldContent, field.TypeString, value)
	}
	if cu.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   comment.UserTable,
			Columns: []string{comment.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   comment.UserTable,
			Columns: []string{comment.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cu.mutation.PostCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   comment.PostTable,
			Columns: []string{comment.PostColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(post.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.PostIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   comment.PostTable,
			Columns: []string{comment.PostColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(post.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{comment.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	cu.mutation.done = true
	return n, nil
}

// CommentUpdateOne is the builder for updating a single Comment entity.
type CommentUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CommentMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (cuo *CommentUpdateOne) SetUpdatedAt(t time.Time) *CommentUpdateOne {
	cuo.mutation.SetUpdatedAt(t)
	return cuo
}

// SetDeleteAt sets the "delete_at" field.
func (cuo *CommentUpdateOne) SetDeleteAt(t time.Time) *CommentUpdateOne {
	cuo.mutation.SetDeleteAt(t)
	return cuo
}

// SetNillableDeleteAt sets the "delete_at" field if the given value is not nil.
func (cuo *CommentUpdateOne) SetNillableDeleteAt(t *time.Time) *CommentUpdateOne {
	if t != nil {
		cuo.SetDeleteAt(*t)
	}
	return cuo
}

// ClearDeleteAt clears the value of the "delete_at" field.
func (cuo *CommentUpdateOne) ClearDeleteAt() *CommentUpdateOne {
	cuo.mutation.ClearDeleteAt()
	return cuo
}

// SetContent sets the "content" field.
func (cuo *CommentUpdateOne) SetContent(s string) *CommentUpdateOne {
	cuo.mutation.SetContent(s)
	return cuo
}

// SetNillableContent sets the "content" field if the given value is not nil.
func (cuo *CommentUpdateOne) SetNillableContent(s *string) *CommentUpdateOne {
	if s != nil {
		cuo.SetContent(*s)
	}
	return cuo
}

// SetUserID sets the "user" edge to the User entity by ID.
func (cuo *CommentUpdateOne) SetUserID(id int) *CommentUpdateOne {
	cuo.mutation.SetUserID(id)
	return cuo
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (cuo *CommentUpdateOne) SetNillableUserID(id *int) *CommentUpdateOne {
	if id != nil {
		cuo = cuo.SetUserID(*id)
	}
	return cuo
}

// SetUser sets the "user" edge to the User entity.
func (cuo *CommentUpdateOne) SetUser(u *User) *CommentUpdateOne {
	return cuo.SetUserID(u.ID)
}

// SetPostID sets the "post" edge to the Post entity by ID.
func (cuo *CommentUpdateOne) SetPostID(id int) *CommentUpdateOne {
	cuo.mutation.SetPostID(id)
	return cuo
}

// SetNillablePostID sets the "post" edge to the Post entity by ID if the given value is not nil.
func (cuo *CommentUpdateOne) SetNillablePostID(id *int) *CommentUpdateOne {
	if id != nil {
		cuo = cuo.SetPostID(*id)
	}
	return cuo
}

// SetPost sets the "post" edge to the Post entity.
func (cuo *CommentUpdateOne) SetPost(p *Post) *CommentUpdateOne {
	return cuo.SetPostID(p.ID)
}

// Mutation returns the CommentMutation object of the builder.
func (cuo *CommentUpdateOne) Mutation() *CommentMutation {
	return cuo.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (cuo *CommentUpdateOne) ClearUser() *CommentUpdateOne {
	cuo.mutation.ClearUser()
	return cuo
}

// ClearPost clears the "post" edge to the Post entity.
func (cuo *CommentUpdateOne) ClearPost() *CommentUpdateOne {
	cuo.mutation.ClearPost()
	return cuo
}

// Where appends a list predicates to the CommentUpdate builder.
func (cuo *CommentUpdateOne) Where(ps ...predicate.Comment) *CommentUpdateOne {
	cuo.mutation.Where(ps...)
	return cuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *CommentUpdateOne) Select(field string, fields ...string) *CommentUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated Comment entity.
func (cuo *CommentUpdateOne) Save(ctx context.Context) (*Comment, error) {
	if err := cuo.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, cuo.sqlSave, cuo.mutation, cuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *CommentUpdateOne) SaveX(ctx context.Context) *Comment {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *CommentUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *CommentUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cuo *CommentUpdateOne) defaults() error {
	if _, ok := cuo.mutation.UpdatedAt(); !ok {
		if comment.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized comment.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := comment.UpdateDefaultUpdatedAt()
		cuo.mutation.SetUpdatedAt(v)
	}
	return nil
}

func (cuo *CommentUpdateOne) sqlSave(ctx context.Context) (_node *Comment, err error) {
	_spec := sqlgraph.NewUpdateSpec(comment.Table, comment.Columns, sqlgraph.NewFieldSpec(comment.FieldID, field.TypeInt))
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Comment.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, comment.FieldID)
		for _, f := range fields {
			if !comment.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != comment.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cuo.mutation.UpdatedAt(); ok {
		_spec.SetField(comment.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := cuo.mutation.DeleteAt(); ok {
		_spec.SetField(comment.FieldDeleteAt, field.TypeTime, value)
	}
	if cuo.mutation.DeleteAtCleared() {
		_spec.ClearField(comment.FieldDeleteAt, field.TypeTime)
	}
	if value, ok := cuo.mutation.Content(); ok {
		_spec.SetField(comment.FieldContent, field.TypeString, value)
	}
	if cuo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   comment.UserTable,
			Columns: []string{comment.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   comment.UserTable,
			Columns: []string{comment.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cuo.mutation.PostCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   comment.PostTable,
			Columns: []string{comment.PostColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(post.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.PostIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   comment.PostTable,
			Columns: []string{comment.PostColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(post.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Comment{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{comment.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	cuo.mutation.done = true
	return _node, nil
}
