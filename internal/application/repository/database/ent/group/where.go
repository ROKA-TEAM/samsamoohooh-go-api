// Code generated by ent, DO NOT EDIT.

package group

import (
	"samsamoohooh-go-api/internal/application/repository/database/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Group {
	return predicate.Group(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Group {
	return predicate.Group(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Group {
	return predicate.Group(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Group {
	return predicate.Group(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Group {
	return predicate.Group(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Group {
	return predicate.Group(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Group {
	return predicate.Group(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Group {
	return predicate.Group(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Group {
	return predicate.Group(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Group {
	return predicate.Group(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Group {
	return predicate.Group(sql.FieldEQ(FieldUpdatedAt, v))
}

// BookTitle applies equality check predicate on the "book_title" field. It's identical to BookTitleEQ.
func BookTitle(v string) predicate.Group {
	return predicate.Group(sql.FieldEQ(FieldBookTitle, v))
}

// Author applies equality check predicate on the "author" field. It's identical to AuthorEQ.
func Author(v string) predicate.Group {
	return predicate.Group(sql.FieldEQ(FieldAuthor, v))
}

// MaxPage applies equality check predicate on the "max_page" field. It's identical to MaxPageEQ.
func MaxPage(v int) predicate.Group {
	return predicate.Group(sql.FieldEQ(FieldMaxPage, v))
}

// Publisher applies equality check predicate on the "publisher" field. It's identical to PublisherEQ.
func Publisher(v string) predicate.Group {
	return predicate.Group(sql.FieldEQ(FieldPublisher, v))
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.Group {
	return predicate.Group(sql.FieldEQ(FieldDescription, v))
}

// BookMark applies equality check predicate on the "book_mark" field. It's identical to BookMarkEQ.
func BookMark(v int) predicate.Group {
	return predicate.Group(sql.FieldEQ(FieldBookMark, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Group {
	return predicate.Group(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Group {
	return predicate.Group(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Group {
	return predicate.Group(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Group {
	return predicate.Group(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Group {
	return predicate.Group(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Group {
	return predicate.Group(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Group {
	return predicate.Group(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Group {
	return predicate.Group(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Group {
	return predicate.Group(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Group {
	return predicate.Group(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Group {
	return predicate.Group(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Group {
	return predicate.Group(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Group {
	return predicate.Group(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Group {
	return predicate.Group(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Group {
	return predicate.Group(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Group {
	return predicate.Group(sql.FieldLTE(FieldUpdatedAt, v))
}

// BookTitleEQ applies the EQ predicate on the "book_title" field.
func BookTitleEQ(v string) predicate.Group {
	return predicate.Group(sql.FieldEQ(FieldBookTitle, v))
}

// BookTitleNEQ applies the NEQ predicate on the "book_title" field.
func BookTitleNEQ(v string) predicate.Group {
	return predicate.Group(sql.FieldNEQ(FieldBookTitle, v))
}

// BookTitleIn applies the In predicate on the "book_title" field.
func BookTitleIn(vs ...string) predicate.Group {
	return predicate.Group(sql.FieldIn(FieldBookTitle, vs...))
}

// BookTitleNotIn applies the NotIn predicate on the "book_title" field.
func BookTitleNotIn(vs ...string) predicate.Group {
	return predicate.Group(sql.FieldNotIn(FieldBookTitle, vs...))
}

// BookTitleGT applies the GT predicate on the "book_title" field.
func BookTitleGT(v string) predicate.Group {
	return predicate.Group(sql.FieldGT(FieldBookTitle, v))
}

// BookTitleGTE applies the GTE predicate on the "book_title" field.
func BookTitleGTE(v string) predicate.Group {
	return predicate.Group(sql.FieldGTE(FieldBookTitle, v))
}

// BookTitleLT applies the LT predicate on the "book_title" field.
func BookTitleLT(v string) predicate.Group {
	return predicate.Group(sql.FieldLT(FieldBookTitle, v))
}

// BookTitleLTE applies the LTE predicate on the "book_title" field.
func BookTitleLTE(v string) predicate.Group {
	return predicate.Group(sql.FieldLTE(FieldBookTitle, v))
}

// BookTitleContains applies the Contains predicate on the "book_title" field.
func BookTitleContains(v string) predicate.Group {
	return predicate.Group(sql.FieldContains(FieldBookTitle, v))
}

// BookTitleHasPrefix applies the HasPrefix predicate on the "book_title" field.
func BookTitleHasPrefix(v string) predicate.Group {
	return predicate.Group(sql.FieldHasPrefix(FieldBookTitle, v))
}

// BookTitleHasSuffix applies the HasSuffix predicate on the "book_title" field.
func BookTitleHasSuffix(v string) predicate.Group {
	return predicate.Group(sql.FieldHasSuffix(FieldBookTitle, v))
}

// BookTitleEqualFold applies the EqualFold predicate on the "book_title" field.
func BookTitleEqualFold(v string) predicate.Group {
	return predicate.Group(sql.FieldEqualFold(FieldBookTitle, v))
}

// BookTitleContainsFold applies the ContainsFold predicate on the "book_title" field.
func BookTitleContainsFold(v string) predicate.Group {
	return predicate.Group(sql.FieldContainsFold(FieldBookTitle, v))
}

// AuthorEQ applies the EQ predicate on the "author" field.
func AuthorEQ(v string) predicate.Group {
	return predicate.Group(sql.FieldEQ(FieldAuthor, v))
}

// AuthorNEQ applies the NEQ predicate on the "author" field.
func AuthorNEQ(v string) predicate.Group {
	return predicate.Group(sql.FieldNEQ(FieldAuthor, v))
}

// AuthorIn applies the In predicate on the "author" field.
func AuthorIn(vs ...string) predicate.Group {
	return predicate.Group(sql.FieldIn(FieldAuthor, vs...))
}

// AuthorNotIn applies the NotIn predicate on the "author" field.
func AuthorNotIn(vs ...string) predicate.Group {
	return predicate.Group(sql.FieldNotIn(FieldAuthor, vs...))
}

// AuthorGT applies the GT predicate on the "author" field.
func AuthorGT(v string) predicate.Group {
	return predicate.Group(sql.FieldGT(FieldAuthor, v))
}

// AuthorGTE applies the GTE predicate on the "author" field.
func AuthorGTE(v string) predicate.Group {
	return predicate.Group(sql.FieldGTE(FieldAuthor, v))
}

// AuthorLT applies the LT predicate on the "author" field.
func AuthorLT(v string) predicate.Group {
	return predicate.Group(sql.FieldLT(FieldAuthor, v))
}

// AuthorLTE applies the LTE predicate on the "author" field.
func AuthorLTE(v string) predicate.Group {
	return predicate.Group(sql.FieldLTE(FieldAuthor, v))
}

// AuthorContains applies the Contains predicate on the "author" field.
func AuthorContains(v string) predicate.Group {
	return predicate.Group(sql.FieldContains(FieldAuthor, v))
}

// AuthorHasPrefix applies the HasPrefix predicate on the "author" field.
func AuthorHasPrefix(v string) predicate.Group {
	return predicate.Group(sql.FieldHasPrefix(FieldAuthor, v))
}

// AuthorHasSuffix applies the HasSuffix predicate on the "author" field.
func AuthorHasSuffix(v string) predicate.Group {
	return predicate.Group(sql.FieldHasSuffix(FieldAuthor, v))
}

// AuthorEqualFold applies the EqualFold predicate on the "author" field.
func AuthorEqualFold(v string) predicate.Group {
	return predicate.Group(sql.FieldEqualFold(FieldAuthor, v))
}

// AuthorContainsFold applies the ContainsFold predicate on the "author" field.
func AuthorContainsFold(v string) predicate.Group {
	return predicate.Group(sql.FieldContainsFold(FieldAuthor, v))
}

// MaxPageEQ applies the EQ predicate on the "max_page" field.
func MaxPageEQ(v int) predicate.Group {
	return predicate.Group(sql.FieldEQ(FieldMaxPage, v))
}

// MaxPageNEQ applies the NEQ predicate on the "max_page" field.
func MaxPageNEQ(v int) predicate.Group {
	return predicate.Group(sql.FieldNEQ(FieldMaxPage, v))
}

// MaxPageIn applies the In predicate on the "max_page" field.
func MaxPageIn(vs ...int) predicate.Group {
	return predicate.Group(sql.FieldIn(FieldMaxPage, vs...))
}

// MaxPageNotIn applies the NotIn predicate on the "max_page" field.
func MaxPageNotIn(vs ...int) predicate.Group {
	return predicate.Group(sql.FieldNotIn(FieldMaxPage, vs...))
}

// MaxPageGT applies the GT predicate on the "max_page" field.
func MaxPageGT(v int) predicate.Group {
	return predicate.Group(sql.FieldGT(FieldMaxPage, v))
}

// MaxPageGTE applies the GTE predicate on the "max_page" field.
func MaxPageGTE(v int) predicate.Group {
	return predicate.Group(sql.FieldGTE(FieldMaxPage, v))
}

// MaxPageLT applies the LT predicate on the "max_page" field.
func MaxPageLT(v int) predicate.Group {
	return predicate.Group(sql.FieldLT(FieldMaxPage, v))
}

// MaxPageLTE applies the LTE predicate on the "max_page" field.
func MaxPageLTE(v int) predicate.Group {
	return predicate.Group(sql.FieldLTE(FieldMaxPage, v))
}

// PublisherEQ applies the EQ predicate on the "publisher" field.
func PublisherEQ(v string) predicate.Group {
	return predicate.Group(sql.FieldEQ(FieldPublisher, v))
}

// PublisherNEQ applies the NEQ predicate on the "publisher" field.
func PublisherNEQ(v string) predicate.Group {
	return predicate.Group(sql.FieldNEQ(FieldPublisher, v))
}

// PublisherIn applies the In predicate on the "publisher" field.
func PublisherIn(vs ...string) predicate.Group {
	return predicate.Group(sql.FieldIn(FieldPublisher, vs...))
}

// PublisherNotIn applies the NotIn predicate on the "publisher" field.
func PublisherNotIn(vs ...string) predicate.Group {
	return predicate.Group(sql.FieldNotIn(FieldPublisher, vs...))
}

// PublisherGT applies the GT predicate on the "publisher" field.
func PublisherGT(v string) predicate.Group {
	return predicate.Group(sql.FieldGT(FieldPublisher, v))
}

// PublisherGTE applies the GTE predicate on the "publisher" field.
func PublisherGTE(v string) predicate.Group {
	return predicate.Group(sql.FieldGTE(FieldPublisher, v))
}

// PublisherLT applies the LT predicate on the "publisher" field.
func PublisherLT(v string) predicate.Group {
	return predicate.Group(sql.FieldLT(FieldPublisher, v))
}

// PublisherLTE applies the LTE predicate on the "publisher" field.
func PublisherLTE(v string) predicate.Group {
	return predicate.Group(sql.FieldLTE(FieldPublisher, v))
}

// PublisherContains applies the Contains predicate on the "publisher" field.
func PublisherContains(v string) predicate.Group {
	return predicate.Group(sql.FieldContains(FieldPublisher, v))
}

// PublisherHasPrefix applies the HasPrefix predicate on the "publisher" field.
func PublisherHasPrefix(v string) predicate.Group {
	return predicate.Group(sql.FieldHasPrefix(FieldPublisher, v))
}

// PublisherHasSuffix applies the HasSuffix predicate on the "publisher" field.
func PublisherHasSuffix(v string) predicate.Group {
	return predicate.Group(sql.FieldHasSuffix(FieldPublisher, v))
}

// PublisherEqualFold applies the EqualFold predicate on the "publisher" field.
func PublisherEqualFold(v string) predicate.Group {
	return predicate.Group(sql.FieldEqualFold(FieldPublisher, v))
}

// PublisherContainsFold applies the ContainsFold predicate on the "publisher" field.
func PublisherContainsFold(v string) predicate.Group {
	return predicate.Group(sql.FieldContainsFold(FieldPublisher, v))
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.Group {
	return predicate.Group(sql.FieldEQ(FieldDescription, v))
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.Group {
	return predicate.Group(sql.FieldNEQ(FieldDescription, v))
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.Group {
	return predicate.Group(sql.FieldIn(FieldDescription, vs...))
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.Group {
	return predicate.Group(sql.FieldNotIn(FieldDescription, vs...))
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.Group {
	return predicate.Group(sql.FieldGT(FieldDescription, v))
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.Group {
	return predicate.Group(sql.FieldGTE(FieldDescription, v))
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.Group {
	return predicate.Group(sql.FieldLT(FieldDescription, v))
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.Group {
	return predicate.Group(sql.FieldLTE(FieldDescription, v))
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.Group {
	return predicate.Group(sql.FieldContains(FieldDescription, v))
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.Group {
	return predicate.Group(sql.FieldHasPrefix(FieldDescription, v))
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.Group {
	return predicate.Group(sql.FieldHasSuffix(FieldDescription, v))
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.Group {
	return predicate.Group(sql.FieldEqualFold(FieldDescription, v))
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.Group {
	return predicate.Group(sql.FieldContainsFold(FieldDescription, v))
}

// BookMarkEQ applies the EQ predicate on the "book_mark" field.
func BookMarkEQ(v int) predicate.Group {
	return predicate.Group(sql.FieldEQ(FieldBookMark, v))
}

// BookMarkNEQ applies the NEQ predicate on the "book_mark" field.
func BookMarkNEQ(v int) predicate.Group {
	return predicate.Group(sql.FieldNEQ(FieldBookMark, v))
}

// BookMarkIn applies the In predicate on the "book_mark" field.
func BookMarkIn(vs ...int) predicate.Group {
	return predicate.Group(sql.FieldIn(FieldBookMark, vs...))
}

// BookMarkNotIn applies the NotIn predicate on the "book_mark" field.
func BookMarkNotIn(vs ...int) predicate.Group {
	return predicate.Group(sql.FieldNotIn(FieldBookMark, vs...))
}

// BookMarkGT applies the GT predicate on the "book_mark" field.
func BookMarkGT(v int) predicate.Group {
	return predicate.Group(sql.FieldGT(FieldBookMark, v))
}

// BookMarkGTE applies the GTE predicate on the "book_mark" field.
func BookMarkGTE(v int) predicate.Group {
	return predicate.Group(sql.FieldGTE(FieldBookMark, v))
}

// BookMarkLT applies the LT predicate on the "book_mark" field.
func BookMarkLT(v int) predicate.Group {
	return predicate.Group(sql.FieldLT(FieldBookMark, v))
}

// BookMarkLTE applies the LTE predicate on the "book_mark" field.
func BookMarkLTE(v int) predicate.Group {
	return predicate.Group(sql.FieldLTE(FieldBookMark, v))
}

// HasUsers applies the HasEdge predicate on the "users" edge.
func HasUsers() predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, UsersTable, UsersPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUsersWith applies the HasEdge predicate on the "users" edge with a given conditions (other predicates).
func HasUsersWith(preds ...predicate.User) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		step := newUsersStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasPosts applies the HasEdge predicate on the "posts" edge.
func HasPosts() predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, PostsTable, PostsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPostsWith applies the HasEdge predicate on the "posts" edge with a given conditions (other predicates).
func HasPostsWith(preds ...predicate.Post) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		step := newPostsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasTasks applies the HasEdge predicate on the "tasks" edge.
func HasTasks() predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, TasksTable, TasksColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTasksWith applies the HasEdge predicate on the "tasks" edge with a given conditions (other predicates).
func HasTasksWith(preds ...predicate.Task) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		step := newTasksStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Group) predicate.Group {
	return predicate.Group(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Group) predicate.Group {
	return predicate.Group(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Group) predicate.Group {
	return predicate.Group(sql.NotPredicates(p))
}
