package firestoretest_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"cloud.google.com/go/firestore"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/stretchr/testify/assert"

	"github.com/t-ash0410/stack-example/go/lib/firestoretest"
)

type DummyData struct {
	ID string `firestore:"ID"`
}

func TestInitFirestoreClient(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		ctx := context.Background()

		fsc, err := firestore.NewClient(ctx, os.Getenv("FIRESTORE_PROJECT_ID"))
		if err != nil {
			t.Fatalf("Failed to create firestore client: %v", err)
		}

		var (
			c  = "dummy-collection"
			id = "dummy"
			dd = &DummyData{
				ID: id,
			}
			bw = fsc.BulkWriter(ctx)
		)
		if _, err := bw.Create(fsc.Doc(fmt.Sprintf("%s/%s", c, id)), dd); err != nil {
			t.Fatalf("Failed to create dummy data: %v", err)
		}
		bw.End()

		fsc, err = firestoretest.InitFirestoreClient(c)
		if err != nil {
			t.Errorf("Unexpected error occurred: %v", err)
			return
		}

		_, err = fsc.Collection(c).Doc(id).Get(ctx)
		if s, ok := status.FromError(err); !ok || s.Code() != codes.NotFound {
			t.Errorf("Unexpected error occurred: %v", err)
			return
		}
	})

	t.Run("Fail: Empty project id", func(t *testing.T) {
		var (
			envKey = "FIRESTORE_PROJECT_ID"
			pid    = os.Getenv(envKey)
		)
		os.Setenv(envKey, "")
		t.Cleanup(func() {
			os.Setenv(envKey, pid)
		})

		_, err := firestoretest.InitFirestoreClient("dummy")
		assert.ErrorContains(t, err, "failed to create firestore client")
	})
}
