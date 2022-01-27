# ExTypes

Extra data types useful for database

## JSON Object

JSON Object is useful for accepting any type for postgres JSONB, Array, and more. You can't put or retrieve them
directly.

## Example

```go
package sample

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/nasermirzaei89/extypes"
)

type Foo struct {
	ID       string
	Tags     []string
	Metadata map[string]interface{}
}

func InsertFoo(ctx context.Context, db *sql.DB, foo Foo) error {
	tags := extypes.NewJSONObject(foo.Tags)
	metadata := extypes.NewJSONObject(foo.Metadata)

	_, err := db.ExecContext(ctx, "INSERT INTO foo (id, tags, metadata) VALUES ($1, $2, $3)", foo.ID, tags, metadata)
	if err != nil {
		return fmt.Errorf("error on exec insert query")
	}

	return nil
}

func FindFoo(ctx context.Context, db *sql.DB, fooID string) (*Foo, error) {
	var (
		res      Foo
		tags     extypes.JSONObject
		metadata extypes.JSONObject
	)

	err := db.QueryRowContext(ctx, "SELECT id, tags, metadata FROM foo WHERE id = $1", fooID).Scan(&res.ID, &tags, &metadata)
	if err != nil {
		return nil, fmt.Errorf("error on query row")
	}

	res.Tags = tags.GetStringSlice()
	res.Metadata = tags.GetStringInterfaceMap()

	return &res, nil
}
```
