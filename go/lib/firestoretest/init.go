package firestoretest

import (
	"context"
	"fmt"

	"cloud.google.com/go/firestore"
	"github.com/google/uuid"
	"google.golang.org/api/iterator"
)

var newFirestoreClient = firestore.NewClient

func InitFirestoreClient(ctx context.Context, collections ...string) (*firestore.Client, error) {
	fsc, err := newFirestoreClient(ctx, uuid.NewString())
	if err != nil {
		return nil, fmt.Errorf("failed to create firestore client: %w", err)
	}

	bw := fsc.BulkWriter(ctx)
	for _, c := range collections {
		if err := enqRemoveAllDocs(ctx, bw, fsc, c); err != nil {
			return nil, fmt.Errorf("failed to remove all documents, collection = %s: %w", c, err)
		}
	}
	bw.End()

	return fsc, nil
}

func enqRemoveAllDocs(ctx context.Context, bw *firestore.BulkWriter, fsc *firestore.Client, col string) error {
	iter := fsc.Collection(col).Documents(ctx)
	defer iter.Stop()
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return err
		}
		if _, err := bw.Delete(doc.Ref); err != nil {
			return err
		}
	}
	return nil
}
