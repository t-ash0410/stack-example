package firestoretest

import (
	"context"
	"fmt"
	"os"
	"testing"

	"cloud.google.com/go/firestore"
	"github.com/stretchr/testify/assert"
	"google.golang.org/api/option"
)

func TestInitFirestoreClient_fail(t *testing.T) {
	t.Run("Failed to create firestore client", func(t *testing.T) {
		newFirestoreClient = func(ctx context.Context, projectID string, opts ...option.ClientOption) (*firestore.Client, error) {
			return nil, fmt.Errorf("some error")
		}
		t.Cleanup(func() {
			newFirestoreClient = firestore.NewClient
		})
		_, err := InitFirestoreClient(context.Background())
		assert.ErrorContains(t, err, "failed to create firestore client")
	})
}

func Test_enqRemoveAllDocs_fail(t *testing.T) {
	type DummyData struct {
		ID string `firestore:"ID"`
	}

	t.Run("Use closed bulk writer", func(t *testing.T) {
		ctx := context.Background()

		fsc, err := firestore.NewClient(ctx, os.Getenv("FIRESTORE_PROJECT_ID"), option.WithScopes())
		if err != nil {
			t.Fatalf("Failed to create firestore client: %v", err)
		}

		var (
			bw = fsc.BulkWriter(ctx)
			c  = "dummy-collection"
			id = "dummy"
			dd = &DummyData{
				ID: id,
			}
		)
		if _, err := bw.Create(fsc.Doc(fmt.Sprintf("%s/%s", c, id)), dd); err != nil {
			t.Fatalf("Failed to create dummy data: %v", err)
		}
		bw.End()

		err = enqRemoveAllDocs(ctx, bw, fsc, c)
		assert.ErrorContains(t, err, "BulkWriter has been closed")
	})
}
