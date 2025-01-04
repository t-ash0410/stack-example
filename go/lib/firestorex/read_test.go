package firestorex_test

import (
	"context"
	"fmt"
	"math"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/t-ash0410/stack-example/go/lib/firestoretest"
	"github.com/t-ash0410/stack-example/go/lib/firestorex"
)

const (
	collectionNameDummy = "dummy-collection"
)

type DummyData struct {
	ID   string    `firestore:"ID"`
	Tag  string    `firestore:"Tag"`
	Date time.Time `firestore:"Date"`
}

func TestReadAll(t *testing.T) {
	date := time.Date(2025, 1, 25, 0, 0, 0, 0, time.UTC)

	t.Run("Success: Read all", func(t *testing.T) {
		ctx := context.Background()

		// Setup firestore client
		fsc, err := firestoretest.InitFirestoreClient(ctx, collectionNameDummy)
		if err != nil {
			t.Fatalf("Failed to create firestore client: %v", err)
		}

		// Prepare dummy data
		bw := fsc.BulkWriter(ctx)
		for i := 0; i < 5; i++ {
			var (
				id = fmt.Sprintf("dummy-%d", i)
				dd = &DummyData{
					ID:   id,
					Tag:  fmt.Sprintf("tag-%v", math.Mod(float64(i), 2)),
					Date: date,
				}
			)
			if _, err := bw.Create(fsc.Doc(fmt.Sprintf("%s/%s", collectionNameDummy, id)), dd); err != nil {
				t.Fatalf("Failed to create dummy data: %v", err)
			}
		}
		bw.End()

		// Run
		var (
			iter = fsc.Collection(collectionNameDummy).Where("Tag", "==", "tag-0").Documents(ctx)
			want = []*DummyData{
				{
					ID:   "dummy-0",
					Tag:  "tag-0",
					Date: date,
				},
				{
					ID:   "dummy-2",
					Tag:  "tag-0",
					Date: date,
				},
				{
					ID:   "dummy-4",
					Tag:  "tag-0",
					Date: date,
				},
			}
		)
		got := []*DummyData{}
		for d, err := range firestorex.ReadEach[DummyData](iter) {
			if !assert.Nil(t, err) {
				return
			}
			got = append(got, d.Data)
		}
		assert.EqualValues(t, want, got)
	})

	t.Run("Success: Read once", func(t *testing.T) {
		ctx := context.Background()

		// Setup firestore client
		fsc, err := firestoretest.InitFirestoreClient(ctx, collectionNameDummy)
		if err != nil {
			t.Fatalf("Failed to create firestore client: %v", err)
		}

		// Prepare dummy data
		bw := fsc.BulkWriter(ctx)
		for i := 0; i < 5; i++ {
			var (
				id = fmt.Sprintf("dummy-%d", i)
				dd = &DummyData{
					ID:   id,
					Tag:  "tag", // All data have the same tag
					Date: date,
				}
			)
			if _, err := bw.Create(fsc.Doc(fmt.Sprintf("%s/%s", collectionNameDummy, id)), dd); err != nil {
				t.Fatalf("Failed to create dummy data: %v", err)
			}
		}
		bw.End()

		// Run
		var (
			iter = fsc.Collection(collectionNameDummy).Where("Tag", "==", "tag").Documents(ctx)
			want = []*DummyData{
				{
					ID:   "dummy-0",
					Tag:  "tag",
					Date: date,
				},
			}
		)
		got := []*DummyData{}
		for d, err := range firestorex.ReadEach[DummyData](iter) {
			if !assert.Nil(t, err) {
				return
			}
			got = append(got, d.Data)
			break // important
		}
		assert.EqualValues(t, want, got)
	})

	t.Run("Fail: Cancelled context", func(t *testing.T) {
		ctx := context.Background()

		// Setup firestore client
		fsc, err := firestoretest.InitFirestoreClient(ctx, collectionNameDummy)
		if err != nil {
			t.Fatalf("Failed to create firestore client: %v", err)
		}

		cctx, cancel := context.WithCancel(ctx)
		cancel() // important

		// Run
		iter := fsc.Collection(collectionNameDummy).Where("Tag", "==", "tag-0").Documents(cctx)
		for _, err := range firestorex.ReadEach[DummyData](iter) {
			assert.ErrorContains(t, err, "failed to read")
		}
	})

	t.Run("Fail: Unmarshal fail", func(t *testing.T) {
		ctx := context.Background()

		// Setup firestore client
		fsc, err := firestoretest.InitFirestoreClient(ctx, collectionNameDummy)
		if err != nil {
			t.Fatalf("Failed to create firestore client: %v", err)
		}

		// Prepare dummy data
		var (
			invalid = map[string]any{
				"Date": "invalid date", // important
			}
			bw = fsc.BulkWriter(ctx)
		)
		if _, err := bw.Create(fsc.Doc(fmt.Sprintf("%s/%s", collectionNameDummy, "dummy")), invalid); err != nil {
			t.Fatalf("Failed to create dummy data: %v", err)
		}
		bw.End()

		// Run
		iter := fsc.Collection(collectionNameDummy).Where("Date", "==", "invalid date").Documents(ctx)
		for _, err := range firestorex.ReadEach[DummyData](iter) {
			assert.ErrorContains(t, err, "failed to unmarshal")
		}
	})
}