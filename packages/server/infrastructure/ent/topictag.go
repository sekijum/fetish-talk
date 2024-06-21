// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"server/infrastructure/ent/topictag"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// TopicTag is the model entity for the TopicTag schema.
type TopicTag struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// CreatedAt holds the value of the "createdAt" field.
	CreatedAt time.Time `json:"createdAt,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the TopicTagQuery when eager-loading is set.
	Edges        TopicTagEdges `json:"edges"`
	selectValues sql.SelectValues
}

// TopicTagEdges holds the relations/edges for other nodes in the graph.
type TopicTagEdges struct {
	// Topics holds the value of the topics edge.
	Topics []*Topic `json:"topics,omitempty"`
	// TopicTaggings holds the value of the topic_taggings edge.
	TopicTaggings []*TopicTagging `json:"topic_taggings,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// TopicsOrErr returns the Topics value or an error if the edge
// was not loaded in eager-loading.
func (e TopicTagEdges) TopicsOrErr() ([]*Topic, error) {
	if e.loadedTypes[0] {
		return e.Topics, nil
	}
	return nil, &NotLoadedError{edge: "topics"}
}

// TopicTaggingsOrErr returns the TopicTaggings value or an error if the edge
// was not loaded in eager-loading.
func (e TopicTagEdges) TopicTaggingsOrErr() ([]*TopicTagging, error) {
	if e.loadedTypes[1] {
		return e.TopicTaggings, nil
	}
	return nil, &NotLoadedError{edge: "topic_taggings"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*TopicTag) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case topictag.FieldID:
			values[i] = new(sql.NullInt64)
		case topictag.FieldName:
			values[i] = new(sql.NullString)
		case topictag.FieldCreatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the TopicTag fields.
func (tt *TopicTag) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case topictag.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			tt.ID = int(value.Int64)
		case topictag.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				tt.Name = value.String
			}
		case topictag.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field createdAt", values[i])
			} else if value.Valid {
				tt.CreatedAt = value.Time
			}
		default:
			tt.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the TopicTag.
// This includes values selected through modifiers, order, etc.
func (tt *TopicTag) Value(name string) (ent.Value, error) {
	return tt.selectValues.Get(name)
}

// QueryTopics queries the "topics" edge of the TopicTag entity.
func (tt *TopicTag) QueryTopics() *TopicQuery {
	return NewTopicTagClient(tt.config).QueryTopics(tt)
}

// QueryTopicTaggings queries the "topic_taggings" edge of the TopicTag entity.
func (tt *TopicTag) QueryTopicTaggings() *TopicTaggingQuery {
	return NewTopicTagClient(tt.config).QueryTopicTaggings(tt)
}

// Update returns a builder for updating this TopicTag.
// Note that you need to call TopicTag.Unwrap() before calling this method if this TopicTag
// was returned from a transaction, and the transaction was committed or rolled back.
func (tt *TopicTag) Update() *TopicTagUpdateOne {
	return NewTopicTagClient(tt.config).UpdateOne(tt)
}

// Unwrap unwraps the TopicTag entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (tt *TopicTag) Unwrap() *TopicTag {
	_tx, ok := tt.config.driver.(*txDriver)
	if !ok {
		panic("ent: TopicTag is not a transactional entity")
	}
	tt.config.driver = _tx.drv
	return tt
}

// String implements the fmt.Stringer.
func (tt *TopicTag) String() string {
	var builder strings.Builder
	builder.WriteString("TopicTag(")
	builder.WriteString(fmt.Sprintf("id=%v, ", tt.ID))
	builder.WriteString("name=")
	builder.WriteString(tt.Name)
	builder.WriteString(", ")
	builder.WriteString("createdAt=")
	builder.WriteString(tt.CreatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// TopicTags is a parsable slice of TopicTag.
type TopicTags []*TopicTag
