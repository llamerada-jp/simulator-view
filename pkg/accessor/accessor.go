/**
 * Copyright 2020-2020 Yuji Ito <llamerada.jp@gmail.com>
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package accessor

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	timeout = 10 * time.Second
)

// Accessor contain mongodb client and collections
type Accessor struct {
	client     *mongo.Client
	collection *mongo.Collection
}

// NewAccessor makes new connection to mongoDB using target URI and etc
func NewAccessor(uri, database, collection string) (*Accessor, error) {
	// make context
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	// connect mongodb
	c, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}
	defer c.Disconnect(ctx)

	coll := c.Database(database).Collection(collection)

	return &Accessor{
		client:     c,
		collection: coll,
	}, nil
}

// Disconnect close the connection form mongDB
func (acc *Accessor) Disconnect() {
	// make context
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	acc.client.Disconnect(ctx)
}
