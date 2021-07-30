package format

import (
	"context"

	"cloud.google.com/go/firestore"
)

var _ Formatter = new(ColRef)

type ColRef struct {
	Prefix  string
	Target  *firestore.CollectionRef
	current *DocIter
	init    bool
}

func (c *ColRef) Next() (string, error) {
	if !c.init {
		c.init = true
		c.current = &DocIter{
			Prefix: c.Prefix + AppendPrefix,
			Target: c.Target.DocumentRefs(context.Background()),
		}
		return c.Prefix + c.Target.ID + "/", nil
	}
	return c.current.Next()
}
