package debug

import (
	"context"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

func DeleteAll(client *firestore.Client) error {
	ctx := context.Background()
	return DeleteCollections(ctx, client, client.Collections(ctx), 24)
}

func DeleteCollections(ctx context.Context, client *firestore.Client,
	ref *firestore.CollectionIterator, batchSize int) error {
	for {
		next, err := ref.Next()
		if err != nil {
			return nil
		}
		if err := DeleteCollection(ctx, client, next, batchSize); err != nil {
			return err
		}
	}
}

/*
Retrieved from
https://github.com/GoogleCloudPlatform/golang-samples/blob/master/firestore/firestore_snippets/save.go
*/
func DeleteCollection(ctx context.Context, client *firestore.Client,
	ref *firestore.CollectionRef, batchSize int) error {

	for {
		// Get a batch of documents
		iter := ref.Limit(batchSize).Documents(ctx)
		numDeleted := 0

		// Iterate through the documents, adding
		// a delete operation for each one to a
		// WriteBatch.
		batch := client.Batch()
		for {
			doc, err := iter.Next()
			if err == iterator.Done {
				break
			}
			if err != nil {
				return err
			}

			batch.Delete(doc.Ref)
			numDeleted++
		}

		// If there are no documents to delete,
		// the process is over.
		if numDeleted == 0 {
			return nil
		}

		_, err := batch.Commit(ctx)
		if err != nil {
			return err
		}
	}
}
