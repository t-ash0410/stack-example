package firestoretest

import (
	"context"
	"fmt"
	"os"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

func InitFirestoreClient(collections ...string) (*firestore.Client, error) {
	ctx := context.Background()

	fsc, err := firestore.NewClient(ctx, os.Getenv("FIRESTORE_PROJECT_ID"))
	if err != nil {
		return nil, fmt.Errorf("failed to create firestore client: %w", err)
	}

	for _, c := range collections {
		if err := removeAllDocs(ctx, fsc, c); err != nil {
			return nil, fmt.Errorf("failed to remove all documents, collection = %s: %w", c, err)
		}
	}

	return fsc, nil
}

func removeAllDocs(ctx context.Context, fsc *firestore.Client, c string) error {
	bw := fsc.BulkWriter(ctx)
	iter := fsc.Collection(c).Documents(ctx)
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
	bw.End()
	return nil
}
