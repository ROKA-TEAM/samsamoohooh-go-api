// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"samsamoohooh-go-api/internal/repository/database/ent/predicate"
	"samsamoohooh-go-api/internal/repository/database/ent/task"
	"samsamoohooh-go-api/internal/repository/database/ent/topic"
	"samsamoohooh-go-api/internal/repository/database/ent/user"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TopicUpdate is the builder for updating Topic entities.
type TopicUpdate struct {
	config
	hooks    []Hook
	mutation *TopicMutation
}

// Where appends a list predicates to the TopicUpdate builder.
func (tu *TopicUpdate) Where(ps ...predicate.Topic) *TopicUpdate {
	tu.mutation.Where(ps...)
	return tu
}

// SetUpdatedAt sets the "updated_at" field.
func (tu *TopicUpdate) SetUpdatedAt(t time.Time) *TopicUpdate {
	tu.mutation.SetUpdatedAt(t)
	return tu
}

// SetDeleteAt sets the "delete_at" field.
func (tu *TopicUpdate) SetDeleteAt(t time.Time) *TopicUpdate {
	tu.mutation.SetDeleteAt(t)
	return tu
}

// SetNillableDeleteAt sets the "delete_at" field if the given value is not nil.
func (tu *TopicUpdate) SetNillableDeleteAt(t *time.Time) *TopicUpdate {
	if t != nil {
		tu.SetDeleteAt(*t)
	}
	return tu
}

// ClearDeleteAt clears the value of the "delete_at" field.
func (tu *TopicUpdate) ClearDeleteAt() *TopicUpdate {
	tu.mutation.ClearDeleteAt()
	return tu
}

// SetField sets the "field" field.
func (tu *TopicUpdate) SetField(s string) *TopicUpdate {
	tu.mutation.SetFieldField(s)
	return tu
}

// SetNillableField sets the "field" field if the given value is not nil.
func (tu *TopicUpdate) SetNillableField(s *string) *TopicUpdate {
	if s != nil {
		tu.SetField(*s)
	}
	return tu
}

// SetFeeling sets the "feeling" field.
func (tu *TopicUpdate) SetFeeling(s string) *TopicUpdate {
	tu.mutation.SetFeeling(s)
	return tu
}

// SetNillableFeeling sets the "feeling" field if the given value is not nil.
func (tu *TopicUpdate) SetNillableFeeling(s *string) *TopicUpdate {
	if s != nil {
		tu.SetFeeling(*s)
	}
	return tu
}

// SetUserID sets the "user" edge to the User entity by ID.
func (tu *TopicUpdate) SetUserID(id int) *TopicUpdate {
	tu.mutation.SetUserID(id)
	return tu
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (tu *TopicUpdate) SetNillableUserID(id *int) *TopicUpdate {
	if id != nil {
		tu = tu.SetUserID(*id)
	}
	return tu
}

// SetUser sets the "user" edge to the User entity.
func (tu *TopicUpdate) SetUser(u *User) *TopicUpdate {
	return tu.SetUserID(u.ID)
}

// SetTaskID sets the "task" edge to the Task entity by ID.
func (tu *TopicUpdate) SetTaskID(id int) *TopicUpdate {
	tu.mutation.SetTaskID(id)
	return tu
}

// SetNillableTaskID sets the "task" edge to the Task entity by ID if the given value is not nil.
func (tu *TopicUpdate) SetNillableTaskID(id *int) *TopicUpdate {
	if id != nil {
		tu = tu.SetTaskID(*id)
	}
	return tu
}

// SetTask sets the "task" edge to the Task entity.
func (tu *TopicUpdate) SetTask(t *Task) *TopicUpdate {
	return tu.SetTaskID(t.ID)
}

// Mutation returns the TopicMutation object of the builder.
func (tu *TopicUpdate) Mutation() *TopicMutation {
	return tu.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (tu *TopicUpdate) ClearUser() *TopicUpdate {
	tu.mutation.ClearUser()
	return tu
}

// ClearTask clears the "task" edge to the Task entity.
func (tu *TopicUpdate) ClearTask() *TopicUpdate {
	tu.mutation.ClearTask()
	return tu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tu *TopicUpdate) Save(ctx context.Context) (int, error) {
	if err := tu.defaults(); err != nil {
		return 0, err
	}
	return withHooks(ctx, tu.sqlSave, tu.mutation, tu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tu *TopicUpdate) SaveX(ctx context.Context) int {
	affected, err := tu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tu *TopicUpdate) Exec(ctx context.Context) error {
	_, err := tu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tu *TopicUpdate) ExecX(ctx context.Context) {
	if err := tu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tu *TopicUpdate) defaults() error {
	if _, ok := tu.mutation.UpdatedAt(); !ok {
		if topic.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized topic.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := topic.UpdateDefaultUpdatedAt()
		tu.mutation.SetUpdatedAt(v)
	}
	return nil
}

func (tu *TopicUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(topic.Table, topic.Columns, sqlgraph.NewFieldSpec(topic.FieldID, field.TypeInt))
	if ps := tu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tu.mutation.UpdatedAt(); ok {
		_spec.SetField(topic.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := tu.mutation.DeleteAt(); ok {
		_spec.SetField(topic.FieldDeleteAt, field.TypeTime, value)
	}
	if tu.mutation.DeleteAtCleared() {
		_spec.ClearField(topic.FieldDeleteAt, field.TypeTime)
	}
	if value, ok := tu.mutation.GetField(); ok {
		_spec.SetField(topic.FieldField, field.TypeString, value)
	}
	if value, ok := tu.mutation.Feeling(); ok {
		_spec.SetField(topic.FieldFeeling, field.TypeString, value)
	}
	if tu.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   topic.UserTable,
			Columns: []string{topic.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   topic.UserTable,
			Columns: []string{topic.UserColumn},
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
	if tu.mutation.TaskCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   topic.TaskTable,
			Columns: []string{topic.TaskColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(task.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.TaskIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   topic.TaskTable,
			Columns: []string{topic.TaskColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, tu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{topic.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	tu.mutation.done = true
	return n, nil
}

// TopicUpdateOne is the builder for updating a single Topic entity.
type TopicUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TopicMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (tuo *TopicUpdateOne) SetUpdatedAt(t time.Time) *TopicUpdateOne {
	tuo.mutation.SetUpdatedAt(t)
	return tuo
}

// SetDeleteAt sets the "delete_at" field.
func (tuo *TopicUpdateOne) SetDeleteAt(t time.Time) *TopicUpdateOne {
	tuo.mutation.SetDeleteAt(t)
	return tuo
}

// SetNillableDeleteAt sets the "delete_at" field if the given value is not nil.
func (tuo *TopicUpdateOne) SetNillableDeleteAt(t *time.Time) *TopicUpdateOne {
	if t != nil {
		tuo.SetDeleteAt(*t)
	}
	return tuo
}

// ClearDeleteAt clears the value of the "delete_at" field.
func (tuo *TopicUpdateOne) ClearDeleteAt() *TopicUpdateOne {
	tuo.mutation.ClearDeleteAt()
	return tuo
}

// SetField sets the "field" field.
func (tuo *TopicUpdateOne) SetField(s string) *TopicUpdateOne {
	tuo.mutation.SetFieldField(s)
	return tuo
}

// SetNillableField sets the "field" field if the given value is not nil.
func (tuo *TopicUpdateOne) SetNillableField(s *string) *TopicUpdateOne {
	if s != nil {
		tuo.SetField(*s)
	}
	return tuo
}

// SetFeeling sets the "feeling" field.
func (tuo *TopicUpdateOne) SetFeeling(s string) *TopicUpdateOne {
	tuo.mutation.SetFeeling(s)
	return tuo
}

// SetNillableFeeling sets the "feeling" field if the given value is not nil.
func (tuo *TopicUpdateOne) SetNillableFeeling(s *string) *TopicUpdateOne {
	if s != nil {
		tuo.SetFeeling(*s)
	}
	return tuo
}

// SetUserID sets the "user" edge to the User entity by ID.
func (tuo *TopicUpdateOne) SetUserID(id int) *TopicUpdateOne {
	tuo.mutation.SetUserID(id)
	return tuo
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (tuo *TopicUpdateOne) SetNillableUserID(id *int) *TopicUpdateOne {
	if id != nil {
		tuo = tuo.SetUserID(*id)
	}
	return tuo
}

// SetUser sets the "user" edge to the User entity.
func (tuo *TopicUpdateOne) SetUser(u *User) *TopicUpdateOne {
	return tuo.SetUserID(u.ID)
}

// SetTaskID sets the "task" edge to the Task entity by ID.
func (tuo *TopicUpdateOne) SetTaskID(id int) *TopicUpdateOne {
	tuo.mutation.SetTaskID(id)
	return tuo
}

// SetNillableTaskID sets the "task" edge to the Task entity by ID if the given value is not nil.
func (tuo *TopicUpdateOne) SetNillableTaskID(id *int) *TopicUpdateOne {
	if id != nil {
		tuo = tuo.SetTaskID(*id)
	}
	return tuo
}

// SetTask sets the "task" edge to the Task entity.
func (tuo *TopicUpdateOne) SetTask(t *Task) *TopicUpdateOne {
	return tuo.SetTaskID(t.ID)
}

// Mutation returns the TopicMutation object of the builder.
func (tuo *TopicUpdateOne) Mutation() *TopicMutation {
	return tuo.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (tuo *TopicUpdateOne) ClearUser() *TopicUpdateOne {
	tuo.mutation.ClearUser()
	return tuo
}

// ClearTask clears the "task" edge to the Task entity.
func (tuo *TopicUpdateOne) ClearTask() *TopicUpdateOne {
	tuo.mutation.ClearTask()
	return tuo
}

// Where appends a list predicates to the TopicUpdate builder.
func (tuo *TopicUpdateOne) Where(ps ...predicate.Topic) *TopicUpdateOne {
	tuo.mutation.Where(ps...)
	return tuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tuo *TopicUpdateOne) Select(field string, fields ...string) *TopicUpdateOne {
	tuo.fields = append([]string{field}, fields...)
	return tuo
}

// Save executes the query and returns the updated Topic entity.
func (tuo *TopicUpdateOne) Save(ctx context.Context) (*Topic, error) {
	if err := tuo.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, tuo.sqlSave, tuo.mutation, tuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tuo *TopicUpdateOne) SaveX(ctx context.Context) *Topic {
	node, err := tuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tuo *TopicUpdateOne) Exec(ctx context.Context) error {
	_, err := tuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tuo *TopicUpdateOne) ExecX(ctx context.Context) {
	if err := tuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tuo *TopicUpdateOne) defaults() error {
	if _, ok := tuo.mutation.UpdatedAt(); !ok {
		if topic.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized topic.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := topic.UpdateDefaultUpdatedAt()
		tuo.mutation.SetUpdatedAt(v)
	}
	return nil
}

func (tuo *TopicUpdateOne) sqlSave(ctx context.Context) (_node *Topic, err error) {
	_spec := sqlgraph.NewUpdateSpec(topic.Table, topic.Columns, sqlgraph.NewFieldSpec(topic.FieldID, field.TypeInt))
	id, ok := tuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Topic.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := tuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, topic.FieldID)
		for _, f := range fields {
			if !topic.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != topic.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tuo.mutation.UpdatedAt(); ok {
		_spec.SetField(topic.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := tuo.mutation.DeleteAt(); ok {
		_spec.SetField(topic.FieldDeleteAt, field.TypeTime, value)
	}
	if tuo.mutation.DeleteAtCleared() {
		_spec.ClearField(topic.FieldDeleteAt, field.TypeTime)
	}
	if value, ok := tuo.mutation.GetField(); ok {
		_spec.SetField(topic.FieldField, field.TypeString, value)
	}
	if value, ok := tuo.mutation.Feeling(); ok {
		_spec.SetField(topic.FieldFeeling, field.TypeString, value)
	}
	if tuo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   topic.UserTable,
			Columns: []string{topic.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   topic.UserTable,
			Columns: []string{topic.UserColumn},
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
	if tuo.mutation.TaskCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   topic.TaskTable,
			Columns: []string{topic.TaskColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(task.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.TaskIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   topic.TaskTable,
			Columns: []string{topic.TaskColumn},
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
	_node = &Topic{config: tuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{topic.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	tuo.mutation.done = true
	return _node, nil
}
