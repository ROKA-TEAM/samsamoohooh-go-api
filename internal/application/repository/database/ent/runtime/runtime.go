// Code generated by ent, DO NOT EDIT.

package runtime

import (
	"samsamoohooh-go-api/internal/application/repository/database/ent/comment"
	"samsamoohooh-go-api/internal/application/repository/database/ent/group"
	"samsamoohooh-go-api/internal/application/repository/database/ent/post"
	"samsamoohooh-go-api/internal/application/repository/database/ent/schema"
	"samsamoohooh-go-api/internal/application/repository/database/ent/task"
	"samsamoohooh-go-api/internal/application/repository/database/ent/topic"
	"samsamoohooh-go-api/internal/application/repository/database/ent/user"
	"time"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	commentMixin := schema.Comment{}.Mixin()
	commentMixinHooks1 := commentMixin[1].Hooks()
	comment.Hooks[0] = commentMixinHooks1[0]
	commentMixinInters1 := commentMixin[1].Interceptors()
	comment.Interceptors[0] = commentMixinInters1[0]
	commentMixinFields0 := commentMixin[0].Fields()
	_ = commentMixinFields0
	commentFields := schema.Comment{}.Fields()
	_ = commentFields
	// commentDescCreatedAt is the schema descriptor for created_at field.
	commentDescCreatedAt := commentMixinFields0[0].Descriptor()
	// comment.DefaultCreatedAt holds the default value on creation for the created_at field.
	comment.DefaultCreatedAt = commentDescCreatedAt.Default.(func() time.Time)
	// commentDescUpdatedAt is the schema descriptor for updated_at field.
	commentDescUpdatedAt := commentMixinFields0[1].Descriptor()
	// comment.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	comment.DefaultUpdatedAt = commentDescUpdatedAt.Default.(func() time.Time)
	// comment.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	comment.UpdateDefaultUpdatedAt = commentDescUpdatedAt.UpdateDefault.(func() time.Time)
	groupMixin := schema.Group{}.Mixin()
	groupMixinHooks1 := groupMixin[1].Hooks()
	group.Hooks[0] = groupMixinHooks1[0]
	groupMixinInters1 := groupMixin[1].Interceptors()
	group.Interceptors[0] = groupMixinInters1[0]
	groupMixinFields0 := groupMixin[0].Fields()
	_ = groupMixinFields0
	groupFields := schema.Group{}.Fields()
	_ = groupFields
	// groupDescCreatedAt is the schema descriptor for created_at field.
	groupDescCreatedAt := groupMixinFields0[0].Descriptor()
	// group.DefaultCreatedAt holds the default value on creation for the created_at field.
	group.DefaultCreatedAt = groupDescCreatedAt.Default.(func() time.Time)
	// groupDescUpdatedAt is the schema descriptor for updated_at field.
	groupDescUpdatedAt := groupMixinFields0[1].Descriptor()
	// group.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	group.DefaultUpdatedAt = groupDescUpdatedAt.Default.(func() time.Time)
	// group.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	group.UpdateDefaultUpdatedAt = groupDescUpdatedAt.UpdateDefault.(func() time.Time)
	postMixin := schema.Post{}.Mixin()
	postMixinHooks1 := postMixin[1].Hooks()
	post.Hooks[0] = postMixinHooks1[0]
	postMixinInters1 := postMixin[1].Interceptors()
	post.Interceptors[0] = postMixinInters1[0]
	postMixinFields0 := postMixin[0].Fields()
	_ = postMixinFields0
	postFields := schema.Post{}.Fields()
	_ = postFields
	// postDescCreatedAt is the schema descriptor for created_at field.
	postDescCreatedAt := postMixinFields0[0].Descriptor()
	// post.DefaultCreatedAt holds the default value on creation for the created_at field.
	post.DefaultCreatedAt = postDescCreatedAt.Default.(func() time.Time)
	// postDescUpdatedAt is the schema descriptor for updated_at field.
	postDescUpdatedAt := postMixinFields0[1].Descriptor()
	// post.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	post.DefaultUpdatedAt = postDescUpdatedAt.Default.(func() time.Time)
	// post.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	post.UpdateDefaultUpdatedAt = postDescUpdatedAt.UpdateDefault.(func() time.Time)
	taskMixin := schema.Task{}.Mixin()
	taskMixinHooks1 := taskMixin[1].Hooks()
	task.Hooks[0] = taskMixinHooks1[0]
	taskMixinInters1 := taskMixin[1].Interceptors()
	task.Interceptors[0] = taskMixinInters1[0]
	taskMixinFields0 := taskMixin[0].Fields()
	_ = taskMixinFields0
	taskFields := schema.Task{}.Fields()
	_ = taskFields
	// taskDescCreatedAt is the schema descriptor for created_at field.
	taskDescCreatedAt := taskMixinFields0[0].Descriptor()
	// task.DefaultCreatedAt holds the default value on creation for the created_at field.
	task.DefaultCreatedAt = taskDescCreatedAt.Default.(func() time.Time)
	// taskDescUpdatedAt is the schema descriptor for updated_at field.
	taskDescUpdatedAt := taskMixinFields0[1].Descriptor()
	// task.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	task.DefaultUpdatedAt = taskDescUpdatedAt.Default.(func() time.Time)
	// task.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	task.UpdateDefaultUpdatedAt = taskDescUpdatedAt.UpdateDefault.(func() time.Time)
	topicMixin := schema.Topic{}.Mixin()
	topicMixinHooks1 := topicMixin[1].Hooks()
	topic.Hooks[0] = topicMixinHooks1[0]
	topicMixinInters1 := topicMixin[1].Interceptors()
	topic.Interceptors[0] = topicMixinInters1[0]
	topicMixinFields0 := topicMixin[0].Fields()
	_ = topicMixinFields0
	topicFields := schema.Topic{}.Fields()
	_ = topicFields
	// topicDescCreatedAt is the schema descriptor for created_at field.
	topicDescCreatedAt := topicMixinFields0[0].Descriptor()
	// topic.DefaultCreatedAt holds the default value on creation for the created_at field.
	topic.DefaultCreatedAt = topicDescCreatedAt.Default.(func() time.Time)
	// topicDescUpdatedAt is the schema descriptor for updated_at field.
	topicDescUpdatedAt := topicMixinFields0[1].Descriptor()
	// topic.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	topic.DefaultUpdatedAt = topicDescUpdatedAt.Default.(func() time.Time)
	// topic.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	topic.UpdateDefaultUpdatedAt = topicDescUpdatedAt.UpdateDefault.(func() time.Time)
	userMixin := schema.User{}.Mixin()
	userMixinHooks1 := userMixin[1].Hooks()
	user.Hooks[0] = userMixinHooks1[0]
	userMixinInters1 := userMixin[1].Interceptors()
	user.Interceptors[0] = userMixinInters1[0]
	userMixinFields0 := userMixin[0].Fields()
	_ = userMixinFields0
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescCreatedAt is the schema descriptor for created_at field.
	userDescCreatedAt := userMixinFields0[0].Descriptor()
	// user.DefaultCreatedAt holds the default value on creation for the created_at field.
	user.DefaultCreatedAt = userDescCreatedAt.Default.(func() time.Time)
	// userDescUpdatedAt is the schema descriptor for updated_at field.
	userDescUpdatedAt := userMixinFields0[1].Descriptor()
	// user.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	user.DefaultUpdatedAt = userDescUpdatedAt.Default.(func() time.Time)
	// user.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	user.UpdateDefaultUpdatedAt = userDescUpdatedAt.UpdateDefault.(func() time.Time)
}

const (
	Version = "v0.14.1"                                         // Version of ent codegen.
	Sum     = "h1:fUERL506Pqr92EPHJqr8EYxbPioflJo6PudkrEA8a/s=" // Sum of ent codegen.
)