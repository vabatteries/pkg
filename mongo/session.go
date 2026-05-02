package mongo

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/mongo"
)


func (c *MongoClient) WithinSession(ctx context.Context, cb WithinSessionFunc) (any, error) {
	
	// 1. Start a new session
	sess, err := c.Client.StartSession()
	if err != nil {
		return nil, err
	}

	defer sess.EndSession(ctx)

	// 2. Create a new context
	ctxNew := mongo.NewSessionContext(ctx, sess)

	// 3. Start transaction
	if err = sess.StartTransaction(); err != nil {
		return nil, err
	}

	res, err := cb(ctxNew, sess)
	if err != nil {
		_ = sess.AbortTransaction(context.Background())

		return nil, err
	}

	if err := sess.CommitTransaction(context.Background()); err != nil {
		return nil, err
	}

	return res, nil
}