package schema

import "entgo.io/ent"

// Subscriber holds the schema definition for the Subscriber entity.
type Subscriber struct {
	ent.Schema
}

// Fields of the Subscriber.
func (Subscriber) Fields() []ent.Field {
	return nil
}

// Edges of the Subscriber.
func (Subscriber) Edges() []ent.Edge {
	return nil
}
