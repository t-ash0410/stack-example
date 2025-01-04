package firestorex

import (
	"fmt"
	"iter"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

func ReadEach[T any](iter *firestore.DocumentIterator) iter.Seq2[*T, error] {
	return func(yield func(*T, error) bool) {
		for {
			doc, err := iter.Next()
			if err == iterator.Done {
				return
			}
			if err != nil {
				yield(nil, fmt.Errorf("failed to read: %w", err))
				return
			}

			var d T
			if err := doc.DataTo(&d); err != nil {
				yield(nil, fmt.Errorf("failed to unmarshal: %w", err))
				return
			}
			if !yield(&d, nil) {
				return
			}
		}
	}
}
