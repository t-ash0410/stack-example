package firestorex

import (
	"fmt"
	"iter"
	"time"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

type ResultWithMeta[T any] struct {
	Data       *T
	CreateTime time.Time
	UpdateTime time.Time
}

func ReadEach[T any](iter *firestore.DocumentIterator) iter.Seq2[*ResultWithMeta[T], error] {
	return func(yield func(*ResultWithMeta[T], error) bool) {
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
			ret := ResultWithMeta[T]{
				Data:       &d,
				CreateTime: doc.CreateTime,
				UpdateTime: doc.UpdateTime,
			}
			if !yield(&ret, nil) {
				return
			}
		}
	}
}
