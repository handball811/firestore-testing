package format

import "cloud.google.com/go/firestore"

var _ Formatter = new(DocIter)

type DocIter struct {
	Prefix  string
	Target  *firestore.DocumentRefIterator
	current *DocRef
}

func (c *DocIter) Next() (string, error) {
	if c.current == nil {
		if c.next() != nil {
			return "", ErrFin
		}
	}
	for {
		resp, err := c.current.Next()
		if err != nil {
			if c.next() != nil {
				return "", ErrFin
			}
		} else {
			return resp, nil
		}
	}
}

func (c *DocIter) next() error {
	cur, err := c.Target.Next()
	if err != nil {
		return ErrFin
	}
	c.current = &DocRef{
		Prefix: c.Prefix + AppendPrefix,
		Target: cur,
	}
	return nil
}
