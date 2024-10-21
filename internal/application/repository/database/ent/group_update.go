// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"samsamoohooh-go-api/internal/application/repository/database/ent/group"
	"samsamoohooh-go-api/internal/application/repository/database/ent/post"
	"samsamoohooh-go-api/internal/application/repository/database/ent/predicate"
	"samsamoohooh-go-api/internal/application/repository/database/ent/task"
	"samsamoohooh-go-api/internal/application/repository/database/ent/user"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// GroupUpdate is the builder for updating Group entities.
type GroupUpdate struct {
	config
	hooks    []Hook
	mutation *GroupMutation
}

// Where appends a list predicates to the GroupUpdate builder.
func (gu *GroupUpdate) Where(ps ...predicate.Group) *GroupUpdate {
	gu.mutation.Where(ps...)
	return gu
}

// SetUpdatedAt sets the "updated_at" field.
func (gu *GroupUpdate) SetUpdatedAt(t time.Time) *GroupUpdate {
	gu.mutation.SetUpdatedAt(t)
	return gu
}

// SetDeleteAt sets the "delete_at" field.
func (gu *GroupUpdate) SetDeleteAt(t time.Time) *GroupUpdate {
	gu.mutation.SetDeleteAt(t)
	return gu
}

// SetNillableDeleteAt sets the "delete_at" field if the given value is not nil.
func (gu *GroupUpdate) SetNillableDeleteAt(t *time.Time) *GroupUpdate {
	if t != nil {
		gu.SetDeleteAt(*t)
	}
	return gu
}

// ClearDeleteAt clears the value of the "delete_at" field.
func (gu *GroupUpdate) ClearDeleteAt() *GroupUpdate {
	gu.mutation.ClearDeleteAt()
	return gu
}

// SetBookTitle sets the "book_title" field.
func (gu *GroupUpdate) SetBookTitle(s string) *GroupUpdate {
	gu.mutation.SetBookTitle(s)
	return gu
}

// SetNillableBookTitle sets the "book_title" field if the given value is not nil.
func (gu *GroupUpdate) SetNillableBookTitle(s *string) *GroupUpdate {
	if s != nil {
		gu.SetBookTitle(*s)
	}
	return gu
}

// SetAuthor sets the "author" field.
func (gu *GroupUpdate) SetAuthor(s string) *GroupUpdate {
	gu.mutation.SetAuthor(s)
	return gu
}

// SetNillableAuthor sets the "author" field if the given value is not nil.
func (gu *GroupUpdate) SetNillableAuthor(s *string) *GroupUpdate {
	if s != nil {
		gu.SetAuthor(*s)
	}
	return gu
}

// SetMaxPage sets the "max_page" field.
func (gu *GroupUpdate) SetMaxPage(i int) *GroupUpdate {
	gu.mutation.ResetMaxPage()
	gu.mutation.SetMaxPage(i)
	return gu
}

// SetNillableMaxPage sets the "max_page" field if the given value is not nil.
func (gu *GroupUpdate) SetNillableMaxPage(i *int) *GroupUpdate {
	if i != nil {
		gu.SetMaxPage(*i)
	}
	return gu
}

// AddMaxPage adds i to the "max_page" field.
func (gu *GroupUpdate) AddMaxPage(i int) *GroupUpdate {
	gu.mutation.AddMaxPage(i)
	return gu
}

// SetPublisher sets the "publisher" field.
func (gu *GroupUpdate) SetPublisher(s string) *GroupUpdate {
	gu.mutation.SetPublisher(s)
	return gu
}

// SetNillablePublisher sets the "publisher" field if the given value is not nil.
func (gu *GroupUpdate) SetNillablePublisher(s *string) *GroupUpdate {
	if s != nil {
		gu.SetPublisher(*s)
	}
	return gu
}

// SetDescription sets the "description" field.
func (gu *GroupUpdate) SetDescription(s string) *GroupUpdate {
	gu.mutation.SetDescription(s)
	return gu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (gu *GroupUpdate) SetNillableDescription(s *string) *GroupUpdate {
	if s != nil {
		gu.SetDescription(*s)
	}
	return gu
}

// SetBookMark sets the "book_mark" field.
func (gu *GroupUpdate) SetBookMark(i int) *GroupUpdate {
	gu.mutation.ResetBookMark()
	gu.mutation.SetBookMark(i)
	return gu
}

// SetNillableBookMark sets the "book_mark" field if the given value is not nil.
func (gu *GroupUpdate) SetNillableBookMark(i *int) *GroupUpdate {
	if i != nil {
		gu.SetBookMark(*i)
	}
	return gu
}

// AddBookMark adds i to the "book_mark" field.
func (gu *GroupUpdate) AddBookMark(i int) *GroupUpdate {
	gu.mutation.AddBookMark(i)
	return gu
}

// AddUserIDs adds the "users" edge to the User entity by IDs.
func (gu *GroupUpdate) AddUserIDs(ids ...int) *GroupUpdate {
	gu.mutation.AddUserIDs(ids...)
	return gu
}

// AddUsers adds the "users" edges to the User entity.
func (gu *GroupUpdate) AddUsers(u ...*User) *GroupUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return gu.AddUserIDs(ids...)
}

// AddPostIDs adds the "posts" edge to the Post entity by IDs.
func (gu *GroupUpdate) AddPostIDs(ids ...int) *GroupUpdate {
	gu.mutation.AddPostIDs(ids...)
	return gu
}

// AddPosts adds the "posts" edges to the Post entity.
func (gu *GroupUpdate) AddPosts(p ...*Post) *GroupUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return gu.AddPostIDs(ids...)
}

// AddTaskIDs adds the "tasks" edge to the Task entity by IDs.
func (gu *GroupUpdate) AddTaskIDs(ids ...int) *GroupUpdate {
	gu.mutation.AddTaskIDs(ids...)
	return gu
}

// AddTasks adds the "tasks" edges to the Task entity.
func (gu *GroupUpdate) AddTasks(t ...*Task) *GroupUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return gu.AddTaskIDs(ids...)
}

// Mutation returns the GroupMutation object of the builder.
func (gu *GroupUpdate) Mutation() *GroupMutation {
	return gu.mutation
}

// ClearUsers clears all "users" edges to the User entity.
func (gu *GroupUpdate) ClearUsers() *GroupUpdate {
	gu.mutation.ClearUsers()
	return gu
}

// RemoveUserIDs removes the "users" edge to User entities by IDs.
func (gu *GroupUpdate) RemoveUserIDs(ids ...int) *GroupUpdate {
	gu.mutation.RemoveUserIDs(ids...)
	return gu
}

// RemoveUsers removes "users" edges to User entities.
func (gu *GroupUpdate) RemoveUsers(u ...*User) *GroupUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return gu.RemoveUserIDs(ids...)
}

// ClearPosts clears all "posts" edges to the Post entity.
func (gu *GroupUpdate) ClearPosts() *GroupUpdate {
	gu.mutation.ClearPosts()
	return gu
}

// RemovePostIDs removes the "posts" edge to Post entities by IDs.
func (gu *GroupUpdate) RemovePostIDs(ids ...int) *GroupUpdate {
	gu.mutation.RemovePostIDs(ids...)
	return gu
}

// RemovePosts removes "posts" edges to Post entities.
func (gu *GroupUpdate) RemovePosts(p ...*Post) *GroupUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return gu.RemovePostIDs(ids...)
}

// ClearTasks clears all "tasks" edges to the Task entity.
func (gu *GroupUpdate) ClearTasks() *GroupUpdate {
	gu.mutation.ClearTasks()
	return gu
}

// RemoveTaskIDs removes the "tasks" edge to Task entities by IDs.
func (gu *GroupUpdate) RemoveTaskIDs(ids ...int) *GroupUpdate {
	gu.mutation.RemoveTaskIDs(ids...)
	return gu
}

// RemoveTasks removes "tasks" edges to Task entities.
func (gu *GroupUpdate) RemoveTasks(t ...*Task) *GroupUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return gu.RemoveTaskIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (gu *GroupUpdate) Save(ctx context.Context) (int, error) {
	if err := gu.defaults(); err != nil {
		return 0, err
	}
	return withHooks(ctx, gu.sqlSave, gu.mutation, gu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (gu *GroupUpdate) SaveX(ctx context.Context) int {
	affected, err := gu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (gu *GroupUpdate) Exec(ctx context.Context) error {
	_, err := gu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gu *GroupUpdate) ExecX(ctx context.Context) {
	if err := gu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (gu *GroupUpdate) defaults() error {
	if _, ok := gu.mutation.UpdatedAt(); !ok {
		if group.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized group.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := group.UpdateDefaultUpdatedAt()
		gu.mutation.SetUpdatedAt(v)
	}
	return nil
}

func (gu *GroupUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(group.Table, group.Columns, sqlgraph.NewFieldSpec(group.FieldID, field.TypeInt))
	if ps := gu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := gu.mutation.UpdatedAt(); ok {
		_spec.SetField(group.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := gu.mutation.DeleteAt(); ok {
		_spec.SetField(group.FieldDeleteAt, field.TypeTime, value)
	}
	if gu.mutation.DeleteAtCleared() {
		_spec.ClearField(group.FieldDeleteAt, field.TypeTime)
	}
	if value, ok := gu.mutation.BookTitle(); ok {
		_spec.SetField(group.FieldBookTitle, field.TypeString, value)
	}
	if value, ok := gu.mutation.Author(); ok {
		_spec.SetField(group.FieldAuthor, field.TypeString, value)
	}
	if value, ok := gu.mutation.MaxPage(); ok {
		_spec.SetField(group.FieldMaxPage, field.TypeInt, value)
	}
	if value, ok := gu.mutation.AddedMaxPage(); ok {
		_spec.AddField(group.FieldMaxPage, field.TypeInt, value)
	}
	if value, ok := gu.mutation.Publisher(); ok {
		_spec.SetField(group.FieldPublisher, field.TypeString, value)
	}
	if value, ok := gu.mutation.Description(); ok {
		_spec.SetField(group.FieldDescription, field.TypeString, value)
	}
	if value, ok := gu.mutation.BookMark(); ok {
		_spec.SetField(group.FieldBookMark, field.TypeInt, value)
	}
	if value, ok := gu.mutation.AddedBookMark(); ok {
		_spec.AddField(group.FieldBookMark, field.TypeInt, value)
	}
	if gu.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   group.UsersTable,
			Columns: group.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gu.mutation.RemovedUsersIDs(); len(nodes) > 0 && !gu.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   group.UsersTable,
			Columns: group.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gu.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   group.UsersTable,
			Columns: group.UsersPrimaryKey,
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
	if gu.mutation.PostsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   group.PostsTable,
			Columns: []string{group.PostsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(post.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gu.mutation.RemovedPostsIDs(); len(nodes) > 0 && !gu.mutation.PostsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   group.PostsTable,
			Columns: []string{group.PostsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(post.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gu.mutation.PostsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   group.PostsTable,
			Columns: []string{group.PostsColumn},
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
	if gu.mutation.TasksCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   group.TasksTable,
			Columns: []string{group.TasksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(task.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gu.mutation.RemovedTasksIDs(); len(nodes) > 0 && !gu.mutation.TasksCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   group.TasksTable,
			Columns: []string{group.TasksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(task.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gu.mutation.TasksIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   group.TasksTable,
			Columns: []string{group.TasksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(task.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, gu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{group.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	gu.mutation.done = true
	return n, nil
}

// GroupUpdateOne is the builder for updating a single Group entity.
type GroupUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *GroupMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (guo *GroupUpdateOne) SetUpdatedAt(t time.Time) *GroupUpdateOne {
	guo.mutation.SetUpdatedAt(t)
	return guo
}

// SetDeleteAt sets the "delete_at" field.
func (guo *GroupUpdateOne) SetDeleteAt(t time.Time) *GroupUpdateOne {
	guo.mutation.SetDeleteAt(t)
	return guo
}

// SetNillableDeleteAt sets the "delete_at" field if the given value is not nil.
func (guo *GroupUpdateOne) SetNillableDeleteAt(t *time.Time) *GroupUpdateOne {
	if t != nil {
		guo.SetDeleteAt(*t)
	}
	return guo
}

// ClearDeleteAt clears the value of the "delete_at" field.
func (guo *GroupUpdateOne) ClearDeleteAt() *GroupUpdateOne {
	guo.mutation.ClearDeleteAt()
	return guo
}

// SetBookTitle sets the "book_title" field.
func (guo *GroupUpdateOne) SetBookTitle(s string) *GroupUpdateOne {
	guo.mutation.SetBookTitle(s)
	return guo
}

// SetNillableBookTitle sets the "book_title" field if the given value is not nil.
func (guo *GroupUpdateOne) SetNillableBookTitle(s *string) *GroupUpdateOne {
	if s != nil {
		guo.SetBookTitle(*s)
	}
	return guo
}

// SetAuthor sets the "author" field.
func (guo *GroupUpdateOne) SetAuthor(s string) *GroupUpdateOne {
	guo.mutation.SetAuthor(s)
	return guo
}

// SetNillableAuthor sets the "author" field if the given value is not nil.
func (guo *GroupUpdateOne) SetNillableAuthor(s *string) *GroupUpdateOne {
	if s != nil {
		guo.SetAuthor(*s)
	}
	return guo
}

// SetMaxPage sets the "max_page" field.
func (guo *GroupUpdateOne) SetMaxPage(i int) *GroupUpdateOne {
	guo.mutation.ResetMaxPage()
	guo.mutation.SetMaxPage(i)
	return guo
}

// SetNillableMaxPage sets the "max_page" field if the given value is not nil.
func (guo *GroupUpdateOne) SetNillableMaxPage(i *int) *GroupUpdateOne {
	if i != nil {
		guo.SetMaxPage(*i)
	}
	return guo
}

// AddMaxPage adds i to the "max_page" field.
func (guo *GroupUpdateOne) AddMaxPage(i int) *GroupUpdateOne {
	guo.mutation.AddMaxPage(i)
	return guo
}

// SetPublisher sets the "publisher" field.
func (guo *GroupUpdateOne) SetPublisher(s string) *GroupUpdateOne {
	guo.mutation.SetPublisher(s)
	return guo
}

// SetNillablePublisher sets the "publisher" field if the given value is not nil.
func (guo *GroupUpdateOne) SetNillablePublisher(s *string) *GroupUpdateOne {
	if s != nil {
		guo.SetPublisher(*s)
	}
	return guo
}

// SetDescription sets the "description" field.
func (guo *GroupUpdateOne) SetDescription(s string) *GroupUpdateOne {
	guo.mutation.SetDescription(s)
	return guo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (guo *GroupUpdateOne) SetNillableDescription(s *string) *GroupUpdateOne {
	if s != nil {
		guo.SetDescription(*s)
	}
	return guo
}

// SetBookMark sets the "book_mark" field.
func (guo *GroupUpdateOne) SetBookMark(i int) *GroupUpdateOne {
	guo.mutation.ResetBookMark()
	guo.mutation.SetBookMark(i)
	return guo
}

// SetNillableBookMark sets the "book_mark" field if the given value is not nil.
func (guo *GroupUpdateOne) SetNillableBookMark(i *int) *GroupUpdateOne {
	if i != nil {
		guo.SetBookMark(*i)
	}
	return guo
}

// AddBookMark adds i to the "book_mark" field.
func (guo *GroupUpdateOne) AddBookMark(i int) *GroupUpdateOne {
	guo.mutation.AddBookMark(i)
	return guo
}

// AddUserIDs adds the "users" edge to the User entity by IDs.
func (guo *GroupUpdateOne) AddUserIDs(ids ...int) *GroupUpdateOne {
	guo.mutation.AddUserIDs(ids...)
	return guo
}

// AddUsers adds the "users" edges to the User entity.
func (guo *GroupUpdateOne) AddUsers(u ...*User) *GroupUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return guo.AddUserIDs(ids...)
}

// AddPostIDs adds the "posts" edge to the Post entity by IDs.
func (guo *GroupUpdateOne) AddPostIDs(ids ...int) *GroupUpdateOne {
	guo.mutation.AddPostIDs(ids...)
	return guo
}

// AddPosts adds the "posts" edges to the Post entity.
func (guo *GroupUpdateOne) AddPosts(p ...*Post) *GroupUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return guo.AddPostIDs(ids...)
}

// AddTaskIDs adds the "tasks" edge to the Task entity by IDs.
func (guo *GroupUpdateOne) AddTaskIDs(ids ...int) *GroupUpdateOne {
	guo.mutation.AddTaskIDs(ids...)
	return guo
}

// AddTasks adds the "tasks" edges to the Task entity.
func (guo *GroupUpdateOne) AddTasks(t ...*Task) *GroupUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return guo.AddTaskIDs(ids...)
}

// Mutation returns the GroupMutation object of the builder.
func (guo *GroupUpdateOne) Mutation() *GroupMutation {
	return guo.mutation
}

// ClearUsers clears all "users" edges to the User entity.
func (guo *GroupUpdateOne) ClearUsers() *GroupUpdateOne {
	guo.mutation.ClearUsers()
	return guo
}

// RemoveUserIDs removes the "users" edge to User entities by IDs.
func (guo *GroupUpdateOne) RemoveUserIDs(ids ...int) *GroupUpdateOne {
	guo.mutation.RemoveUserIDs(ids...)
	return guo
}

// RemoveUsers removes "users" edges to User entities.
func (guo *GroupUpdateOne) RemoveUsers(u ...*User) *GroupUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return guo.RemoveUserIDs(ids...)
}

// ClearPosts clears all "posts" edges to the Post entity.
func (guo *GroupUpdateOne) ClearPosts() *GroupUpdateOne {
	guo.mutation.ClearPosts()
	return guo
}

// RemovePostIDs removes the "posts" edge to Post entities by IDs.
func (guo *GroupUpdateOne) RemovePostIDs(ids ...int) *GroupUpdateOne {
	guo.mutation.RemovePostIDs(ids...)
	return guo
}

// RemovePosts removes "posts" edges to Post entities.
func (guo *GroupUpdateOne) RemovePosts(p ...*Post) *GroupUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return guo.RemovePostIDs(ids...)
}

// ClearTasks clears all "tasks" edges to the Task entity.
func (guo *GroupUpdateOne) ClearTasks() *GroupUpdateOne {
	guo.mutation.ClearTasks()
	return guo
}

// RemoveTaskIDs removes the "tasks" edge to Task entities by IDs.
func (guo *GroupUpdateOne) RemoveTaskIDs(ids ...int) *GroupUpdateOne {
	guo.mutation.RemoveTaskIDs(ids...)
	return guo
}

// RemoveTasks removes "tasks" edges to Task entities.
func (guo *GroupUpdateOne) RemoveTasks(t ...*Task) *GroupUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return guo.RemoveTaskIDs(ids...)
}

// Where appends a list predicates to the GroupUpdate builder.
func (guo *GroupUpdateOne) Where(ps ...predicate.Group) *GroupUpdateOne {
	guo.mutation.Where(ps...)
	return guo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (guo *GroupUpdateOne) Select(field string, fields ...string) *GroupUpdateOne {
	guo.fields = append([]string{field}, fields...)
	return guo
}

// Save executes the query and returns the updated Group entity.
func (guo *GroupUpdateOne) Save(ctx context.Context) (*Group, error) {
	if err := guo.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, guo.sqlSave, guo.mutation, guo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (guo *GroupUpdateOne) SaveX(ctx context.Context) *Group {
	node, err := guo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (guo *GroupUpdateOne) Exec(ctx context.Context) error {
	_, err := guo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (guo *GroupUpdateOne) ExecX(ctx context.Context) {
	if err := guo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (guo *GroupUpdateOne) defaults() error {
	if _, ok := guo.mutation.UpdatedAt(); !ok {
		if group.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized group.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := group.UpdateDefaultUpdatedAt()
		guo.mutation.SetUpdatedAt(v)
	}
	return nil
}

func (guo *GroupUpdateOne) sqlSave(ctx context.Context) (_node *Group, err error) {
	_spec := sqlgraph.NewUpdateSpec(group.Table, group.Columns, sqlgraph.NewFieldSpec(group.FieldID, field.TypeInt))
	id, ok := guo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Group.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := guo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, group.FieldID)
		for _, f := range fields {
			if !group.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != group.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := guo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := guo.mutation.UpdatedAt(); ok {
		_spec.SetField(group.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := guo.mutation.DeleteAt(); ok {
		_spec.SetField(group.FieldDeleteAt, field.TypeTime, value)
	}
	if guo.mutation.DeleteAtCleared() {
		_spec.ClearField(group.FieldDeleteAt, field.TypeTime)
	}
	if value, ok := guo.mutation.BookTitle(); ok {
		_spec.SetField(group.FieldBookTitle, field.TypeString, value)
	}
	if value, ok := guo.mutation.Author(); ok {
		_spec.SetField(group.FieldAuthor, field.TypeString, value)
	}
	if value, ok := guo.mutation.MaxPage(); ok {
		_spec.SetField(group.FieldMaxPage, field.TypeInt, value)
	}
	if value, ok := guo.mutation.AddedMaxPage(); ok {
		_spec.AddField(group.FieldMaxPage, field.TypeInt, value)
	}
	if value, ok := guo.mutation.Publisher(); ok {
		_spec.SetField(group.FieldPublisher, field.TypeString, value)
	}
	if value, ok := guo.mutation.Description(); ok {
		_spec.SetField(group.FieldDescription, field.TypeString, value)
	}
	if value, ok := guo.mutation.BookMark(); ok {
		_spec.SetField(group.FieldBookMark, field.TypeInt, value)
	}
	if value, ok := guo.mutation.AddedBookMark(); ok {
		_spec.AddField(group.FieldBookMark, field.TypeInt, value)
	}
	if guo.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   group.UsersTable,
			Columns: group.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := guo.mutation.RemovedUsersIDs(); len(nodes) > 0 && !guo.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   group.UsersTable,
			Columns: group.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := guo.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   group.UsersTable,
			Columns: group.UsersPrimaryKey,
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
	if guo.mutation.PostsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   group.PostsTable,
			Columns: []string{group.PostsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(post.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := guo.mutation.RemovedPostsIDs(); len(nodes) > 0 && !guo.mutation.PostsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   group.PostsTable,
			Columns: []string{group.PostsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(post.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := guo.mutation.PostsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   group.PostsTable,
			Columns: []string{group.PostsColumn},
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
	if guo.mutation.TasksCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   group.TasksTable,
			Columns: []string{group.TasksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(task.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := guo.mutation.RemovedTasksIDs(); len(nodes) > 0 && !guo.mutation.TasksCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   group.TasksTable,
			Columns: []string{group.TasksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(task.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := guo.mutation.TasksIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   group.TasksTable,
			Columns: []string{group.TasksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(task.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Group{config: guo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, guo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{group.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	guo.mutation.done = true
	return _node, nil
}