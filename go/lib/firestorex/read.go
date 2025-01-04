package firestorex

import (
	"fmt"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

func ReadAll[T any](iter *firestore.DocumentIterator) ([]*T, error) {
	var ret []*T
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("failed to read: %w", err)
		}

		var d T
		if err := doc.DataTo(&d); err != nil {
			return nil, fmt.Errorf("failed to unmarshal: %w", err)
		}
		ret = append(ret, &d)
	}
	return ret, nil
}
