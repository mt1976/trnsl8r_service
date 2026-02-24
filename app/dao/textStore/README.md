# textStore

`textStore` is a DAO package for the TextStore table in the frantic-amphora framework.

## Overview

This package provides:

- **Type-safe database operations** using strongly-typed field queries
- **Audit trail integration** for all CRUD operations
- **Cache management** with automatic hydration and synchronization
- **Background worker** support for async operations
- **Import/Export** capabilities (JSON and CSV formats)
- **Validation** using struct tags

## Entity Definition

The `TextStore` struct represents records in the TextStore table.

## Field Definitions

The `TextStore` struct contains the following fields:

| Field Name | Field Reference | Type | Tags | Purpose |
|------------|----------------|------|------|---------|
| **ID** (required) | `Fields.ID` | `int` | `storm:"id,increment=100"` | Primary key with auto-increment |
| **Key** (required) | `Fields.Key` | `string` | `storm:"index,unique"` | Encoded unique identifier |
| **Raw** (required) | `Fields.Raw` | `string` | `storm:"index,unique"` | Raw unique identifier |
| **Audit** (required) | `Fields.Audit` | `audit.Audit` | `csv:"-"` | Audit trail information |
| **Lock** (required) | `Fields.Lock` | `sync.Mutex` | `csv:"-"` | Record locking for concurrent updates |
| Signature | `Fields.Signature` | `string` | `csv:"-"` |  |
| Domain | `Fields.Domain` | `string` | `csv:"-"` |  |
| Type | `Fields.Type` | `string` | `csv:"-"` |  |
| Original | `Fields.Original` | `string` | `csv:"original"` |  |
| Message | `Fields.Message` | `string` | `csv:"message"` |  |
| SourceApplication | `Fields.SourceApplication` | `string` | `csv:"-"` |  |
| SourceLocale | `Fields.SourceLocale` | `string` | `csv:"-"` |  |
| ConsumedBy | `Fields.ConsumedBy` | `[]string` | `csv:"-"` |  |
| Localised | `Fields.Localised` | `map[string]string` | `csv:"-"` |  |
| Notes | `Fields.Notes` | `string` | `csv:"-"` |  |

| Signature | `Fields.Signature` | `string` | `csv:"-"` |  |
| Domain | `Fields.Domain` | `string` | `csv:"-"` |  |
| Type | `Fields.Type` | `string` | `csv:"-"` |  |
| Original | `Fields.Original` | `string` | `csv:"original"` |  |
| Message | `Fields.Message` | `string` | `csv:"message"` |  |
| SourceApplication | `Fields.SourceApplication` | `string` | `csv:"-"` |  |
| SourceLocale | `Fields.SourceLocale` | `string` | `csv:"-"` |  |
| ConsumedBy | `Fields.ConsumedBy` | `[]string` | `csv:"-"` |  |
| Localised | `Fields.Localised` | `map[string]string` | `csv:"-"` |  |
| Notes | `Fields.Notes` | `string` | `csv:"-"` |  |


**Note:** Fields marked as **(required)** are mandatory framework fields and must not be modified or removed.

### Using Field References

Field references enable type-safe queries throughout the DAO:

```go
// Get a record by a specific field
record, err := textStore.GetBy(textStore.Fields.Key, "abc123")

// Query with WHERE conditions
records, err := textStore.GetAllWhere(textStore.Fields.SomeField, value)

// Count records matching criteria
count, err := textStore.CountWhere(textStore.Fields.Active, true)
```

## Public API

### Exported types/vars

- `type TextStore struct { ... }`
- `var TableName entities.Table`
- `var Fields fieldNames`

### Database lifecycle

- `func Initialise(ctx context.Context, cached bool)`
- `func IsInitialised() bool`
- `func Close()`
- `func GetDatabaseConnections() func() ([]*database.DB, error)`

### Queries

- `func Count() (int, error)`
- `func CountWhere(field entities.Field, value any) (int, error)`
- `func GetBy(field entities.Field, value any) (*TextStore, error)`
- `func GetAll() ([]TextStore, error)`
- `func GetAllWhere(field entities.Field, value any) ([]TextStore, error)`

### Mutations

- `func Delete(ctx context.Context, id int, note string) error`
- `func DeleteBy(ctx context.Context, field entities.Field, value any, note string) error`
- `func Drop() error`
- `func ClearDown(ctx context.Context) error`

### Record methods

- `func (record *TextStore) Validate() error`
- `func (record *TextStore) Update(ctx context.Context, note string) (*TextStore, error)`
- `func (record *TextStore) UpdateWithAction(ctx context.Context, auditAction audit.Action, note string) (*TextStore, error)`
- `func (record *TextStore) Clone(ctx context.Context) (*TextStore, error)`

### Lookups

- `func GetDefaultLookup() (lookup.Lookup, error)`
- `func GetLookup(field, value entities.Field) (lookup.Lookup, error)`

### Cache integration

- `func CacheHydrator(ctx context.Context) func() ([]any, error)`
- `func CacheSynchroniser(ctx context.Context) func(any) error`

### Construction

- `func New() *TextStore`
- `func Create(ctx context.Context, basis *TextStore) (*TextStore, error)`

### Import / Export

- `func (record *TextStore) ExportRecordToJSON(name string)`
- `func ExportAllToJSON(message string)`
- `func (record *TextStore) ExportRecordToCSV(name string) error`
- `func ExportAllToCSV(msg string) error`
- `func ImportAllFromCSV() error`

### Worker

- `func Worker(j jobs.Job, db *database.DB)`

### Debug

- `func (record *TextStore) Spew()`

## Regenerate

- From this package directory, run: `go generate ./...`

## Next edits

- Adjust the domain fields in the model file.
- Update validation/defaulting hooks.
- Replace any placeholder logic (e.g. clone, import processor) with real implementations.

---

## Generation Information

**Generated Date:** 24/02/2026 & 10:04  
**Generated By:** matttownsend (orion)  
**Generated From Template Version:** 0.5.24 - 2026-01-31
