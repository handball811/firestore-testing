package format

import "cloud.google.com/go/firestore"

var _ Formatter = new(ColIter)

type ColIter struct {
	Prefix  string
	Target  *firestore.CollectionIterator
	current *ColRef
}

func (c *ColIter) Next() (string, error) {
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

func (c *ColIter) next() error {
	cur, err := c.Target.Next()
	if err != nil {
		return ErrFin
	}
	c.current = &ColRef{
		Prefix: c.Prefix + AppendPrefix,
		Target: cur,
	}
	return nil
}
