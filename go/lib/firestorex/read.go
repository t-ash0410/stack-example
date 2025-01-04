package firestorex

import (
	"context"
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

			ret, err := marshalResultWithMeta[T](doc)
			if err != nil {
				yield(nil, fmt.Errorf("failed to unmarshal: %w", err))
				return
			}
			if !yield(ret, nil) {
				return
			}
		}
	}
}

func ReadOne[T any](ctx context.Context, ref *firestore.DocumentRef) (*ResultWithMeta[T], error) {
	doc, err := ref.Get(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to read: %w", err)
	}
	return marshalResultWithMeta[T](doc)
}

func ReadOneWithTxn[T any](txn *firestore.Transaction, ref *firestore.DocumentRef) (*ResultWithMeta[T], error) {
	doc, err := txn.Get(ref)
	if err != nil {
		return nil, fmt.Errorf("failed to read: %w", err)
	}
	return marshalResultWithMeta[T](doc)
}

func marshalResultWithMeta[T any](doc *firestore.DocumentSnapshot) (*ResultWithMeta[T], error) {
	var d T
	if err := doc.DataTo(&d); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return &ResultWithMeta[T]{
		Data:       &d,
		CreateTime: doc.CreateTime,
		UpdateTime: doc.UpdateTime,
	}, nil
}
