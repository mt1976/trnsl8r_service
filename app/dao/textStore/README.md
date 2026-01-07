# Text Store DAO

Package `TextStore` provides Data Access Object (DAO) functionality for managing `TextStore` entities.

## Domain Information

- **Domain**: `Text` (from `Domain`)
- **Table Name**: `TextStore` (from `TableName = Domain + "Store"`)

## Package Variables

- `Domain = "Text"`
- `TableName = Domain + "Store"`
- `Fields` provides a structured way to reference model field names.

## Struct Definition

### TextStore

`TextStore` represents a User entity.

| Field           | Type            | Tags                                      | Description                     |
| --------------- | --------------- | ----------------------------------------- | ------------------------------- |
| `ID`            | `int`           | `storm:"id,increment=100"`                | Primary key with auto increment |
| `Key`           | `string`        | `storm:"unique"`                          | Key                             |
| `Raw`           | `string`        | `storm:"unique"`                          | Raw ID before encoding          |
| `UID`           | `string`        | `validate:"required"`                     |                                 |
| `GID`           | `string`        | `storm:"index" validate:"required"`       |                                 |
| `RealName`      | `string`        | `validate:"required,min=5"`               | This field will not be indexed  |
| `UserName`      | `string`        | `validate:"required,min=5"`               |                                 |
| `UserCode`      | `string`        | `storm:"index" validate:"required,min=5"` |                                 |
| `Email`         | `string`        |                                           |                                 |
| `Notes`         | `string`        | `validate:"max=75"`                       |                                 |
| `Active`        | `dao.StormBool` |                                           |                                 |
| `ExampleInt`    | `dao.Int`       |                                           |                                 |
| `ExampleFloat`  | `dao.Float`     |                                           |                                 |
| `ExampleBool`   | `dao.Bool`      |                                           |                                 |
| `ExampleDate`   | `time.Time`     |                                           |                                 |
| `ExampleString` | `string`        |                                           |                                 |
| `LastLogin`     | `time.Time`     |                                           |                                 |
| `LastHost`      | `string`        |                                           |                                 |
| `Audit`         | `audit.Audit`   | `csv:"-"`                                 | Audit data                      |

### Field Constants

```go
Fields.ID            = "ID"
Fields.Key           = "Key"
Fields.Raw           = "Raw"
Fields.Audit         = "Audit"
Fields.UID           = "UID"
Fields.GID           = "GID"
Fields.RealName      = "RealName"
Fields.UserName      = "UserName"
Fields.UserCode      = "UserCode"
Fields.Email         = "Email"
Fields.Notes         = "Notes"
Fields.Active        = "Active"
Fields.ExampleInt    = "ExampleInt"
Fields.ExampleFloat  = "ExampleFloat"
Fields.ExampleBool   = "ExampleBool"
Fields.ExampleDate   = "ExampleDate"
Fields.ExampleString = "ExampleString"
Fields.LastLogin     = "LastLogin"
Fields.LastHost      = "LastHost"
```

## Operations

### Initialisation / Lifecycle

| Function          | Signature                                           | Description |
| ---------------- | --------------------------------------------------- | ----------- |
| `Initialise`     | `Initialise(ctx context.Context)`                   | Sets up the database connection and prepares the DAO for operations. |
| `IsInitialised`  | `IsInitialised() bool`                              | Returns the initialisation status of the DAO. |
| `Close`          | `Close()`                                           | Terminates the connection to the database used by the DAO. |
| `Drop`           | `Drop() error`                                      | Removes the DAO's database entirely. |
| `ClearDown`      | `ClearDown(ctx context.Context) error`              | Removes all records from the DAO's database. |
| `GetDatabaseConnections` | `GetDatabaseConnections() func() ([]*database.DB, error)` | Returns a function that fetches the current database instances. |

### Create

| Function | Signature                                                                                        | Description |
| -------- | ------------------------------------------------------------------------------------------------ | ----------- |
| `New`    | `New() TextStore`                                                                            | Creates a new Text instance. |
| `Create` | `Create(ctx context.Context, userName, uid, realName, email, gid string) (TextStore, error)` | Creates a new Text instance in the database. |
| `Add`    | `Add(ctx context.Context) (TextStore, error)`                                                | (No doc comment) |
| `(*TextStore) Create` | `(record *TextStore) Create(ctx context.Context, note string) error`            | Inserts a new TextStore record into the database. |

### Read

| Function      | Signature                                                          | Description |
| ------------- | ------------------------------------------------------------------ | ----------- |
| `GetBy`       | `GetBy(field dao.Field, value any) (TextStore, error)`         | Retrieves a record by specified field and value. |
| `GetById`     | `GetById(id any) (TextStore, error)`                           | Retrieves a record by ID. |
| `GetByKey`    | `GetByKey(key any) (TextStore, error)`                         | Retrieves a record by key. |
| `GetAll`      | `GetAll() ([]TextStore, error)`                                | Retrieves all records. |
| `GetAllWhere` | `GetAllWhere(field dao.Field, value any) ([]TextStore, error)` | Retrieves all records matching criteria. |
| `Count`       | `Count() (int, error)`                                             | Returns the total number of records. |
| `CountWhere`  | `CountWhere(field dao.Field, value any) (int, error)`              | Counts records matching criteria. |

### Update

| Function           | Signature                                                                                                      | Description |
| ------------------ | -------------------------------------------------------------------------------------------------------------- | ----------- |
| `(*TextStore) Update`           | `(record *TextStore) Update(ctx context.Context, note string) error`                         | Updates the record in the database. |
| `(*TextStore) UpdateWithAction` | `(record *TextStore) UpdateWithAction(ctx context.Context, auditAction audit.Action, note string) error` | Updates with a specified audit action. |
| `(*TextStore) SetName`          | `(u *TextStore) SetName(name string) error`                                                  | (No doc comment) |

### Delete

| Function      | Signature                                                                      | Description |
| ------------- | ------------------------------------------------------------------------------ | ----------- |
| `Delete`      | `Delete(ctx context.Context, id int, note string) error`                       | Deletes a record by ID. |
| `DeleteBy`    | `DeleteBy(ctx context.Context, field dao.Field, value any, note string) error` | Deletes a record by specified field and value. |
| `DeleteByKey` | `DeleteByKey(ctx context.Context, key string, note string) error`              | Deletes a record by key. |

### Validation / Utility

| Function | Signature                                           | Description |
| -------- | --------------------------------------------------- | ----------- |
| `(*TextStore) Validate` | `(record *TextStore) Validate() error`                    | Checks if the record is valid. |
| `(*TextStore) Clone`    | `(record *TextStore) Clone(ctx context.Context) (TextStore, error)` | Clones the current record in the database. |
| `(*TextStore) Spew`     | `(record *TextStore) Spew()`                               | Outputs the record contents to the Info log. |
| `Login`                     | `Login(ctx context.Context)`                                   | (No doc comment) |

### Lookup

| Function           | Signature                                          | Description |
| ------------------ | -------------------------------------------------- | ----------- |
| `GetDefaultLookup` | `GetDefaultLookup() (lookup.Lookup, error)`        | Builds a default lookup using `Key` and `Raw`. |
| `GetLookup`        | `GetLookup(field, value dao.Field) (lookup.Lookup, error)` | Builds a lookup for specified fields. |

### Import / Export

| Function | Signature | Description |
| -------- | --------- | ----------- |
| `ExportAllAsCSV`  | `ExportAllAsCSV() error`              | Exports all records as a CSV file. |
| `ExportAllAsJSON` | `ExportAllAsJSON(message string)`     | Exports all records as JSON files. |
| `ImportAllFromCSV` | `ImportAllFromCSV() error`           | (No doc comment) |
| `(*TextStore) ExportRecordAsCSV`  | `(record *TextStore) ExportRecordAsCSV(name string) error` | Exports a single record as CSV. |
| `(*TextStore) ExportRecordAsJSON` | `(record *TextStore) ExportRecordAsJSON(name string)`      | Exports a single record as JSON. |

### Worker / Cache

| Function  | Signature                                   | Description |
| --------- | ------------------------------------------- | ----------- |
| `PreLoad` | `PreLoad(ctx context.Context) error`        | Preloads the cache from the database. |
| `Worker`  | `Worker(j jobs.Job, db *database.DB)`       | Job scheduled to run at a predefined interval. |

## Usage Examples

> Note: call `Initialise(ctx)` once during application startup (before CRUD operations), and `Close()` during shutdown.

### Initialise and create a record

```go
package main

import (
    "context"
    "log"

    "github.com/mt1976/frantic-core/dao/test/TextStore"
)

func main() {
    ctx := context.Background()

    TextStore.Initialise(ctx)
    defer TextStore.Close()

    rec, err := TextStore.Create(ctx, "jdoe", "1001", "John Doe", "jdoe@example.com", "group-1")
    if err != nil {
        log.Fatal(err)
    }

    log.Printf("created ID=%d Key=%s", rec.ID, rec.Key)
}
```

### Read, update, validate, and delete

```go
package main

import (
    "context"
    "log"

    "github.com/mt1976/frantic-core/dao/test/TextStore"
)

func main() {
    ctx := context.Background()

    TextStore.Initialise(ctx)
    defer TextStore.Close()

    // Fetch by a field (recommended)
    rec, err := TextStore.GetBy(TextStore.Fields.UserCode, "1001:jdoe")
    if err != nil {
        log.Fatal(err)
    }

    // Update
    rec.Email = "john.doe@newdomain.test"
    if err := rec.Validate(); err != nil {
        log.Fatal(err)
    }
    if err := rec.Update(ctx, "updated email"); err != nil {
        log.Fatal(err)
    }

    // Delete by ID
    if err := TextStore.Delete(ctx, rec.ID, "cleanup"); err != nil {
        log.Fatal(err)
    }
}
```

### Lookup and export

```go
package main

import (
    "context"
    "log"

    "github.com/mt1976/frantic-core/dao/test/TextStore"
)

func main() {
    ctx := context.Background()

    TextStore.Initialise(ctx)
    defer TextStore.Close()

    // Build a default lookup (Key -> Raw)
    lu, err := TextStore.GetDefaultLookup()
    if err != nil {
        log.Fatal(err)
    }
    _ = lu // use lookup

    // Export all records
    if err := TextStore.ExportAllAsCSV(); err != nil {
        log.Fatal(err)
    }
    TextStore.ExportAllAsJSON("nightly backup")
}
```

### Cache preload

```go
package main

import (
    "context"
    "log"

    "github.com/mt1976/frantic-core/dao/test/TextStore"
)

func main() {
    ctx := context.Background()

    TextStore.Initialise(ctx)
    defer TextStore.Close()

    if err := TextStore.PreLoad(ctx); err != nil {
        log.Fatal(err)
    }
}
```

## Package Files

- `TextStore.go`
- `TextStoreCache.go`
- `TextStoreDB.go`
- `TextStoreHelpers.go`
- `TextStoreImpex.go`
- `TextStoreInternals.go`
- `TextStoreModel.go`
- `TextStoreNew.go`
- `TextStoreWorker.go`
