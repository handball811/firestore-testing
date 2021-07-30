package format

import (
	"context"

	"cloud.google.com/go/firestore"
)

var _ Formatter = new(DocRef)

/*

show fields

show sub collections

*/
type DocRef struct {
	Prefix string
	Target *firestore.DocumentRef

	current Formatter
	state   int
}

func (c *DocRef) Next() (string, error) {
	if c.state == 0 {
		c.state = 1
		snap, err := c.Target.Get(context.Background())
		if err == nil {
			c.current = &Map{
				Prefix: c.Prefix + AppendPrefix,
				Target: snap.Data(),
			}

		} else {
			c.state = 2
		}
		return c.Prefix + c.Target.ID + "/", nil
	}
	if c.state == 1 {
		resp, err := c.current.Next()
		if err == nil {
			return resp, nil
		}
		c.state = 2
		c.current = &ColIter{
			Prefix: c.Prefix + AppendPrefix,
			Target: c.Target.Collections(context.Background()),
		}
	}
	return c.current.Next()
}
