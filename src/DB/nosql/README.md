# nosql

This directory contains NoSQL scripts (MongoDB) to create and maintain the database for the `common` package, analogous to the ANSI SQL structure in [`src/DB/sql/`](../sql/).

## Collection

### `sources`

Stores source entries without duplicates. Each field maps to `HTSourceElement` defined in [`src/go/data-type.go`](../../go/data-type.go#L32-L38):

| Field             | BSON Type   | Description                        | Go field      |
|-------------------|-------------|------------------------------------|---------------|
| `src_id`          | `binData`   | Unique binary identifier           | `ID`          |
| `src_citation`    | `string`    | Source citation text               | `Citation`    |
| `src_date`        | `string`    | Date associated with the source    | `Date`        |
| `src_publish_date`| `string`    | Publication date of the source     | `PublishDate` |
| `src_url`         | `string`    | Source URL                         | `URL`         |

Indexed on `src_citation` and `src_id` (unique) for faster lookups.

## Files

| File | Description |
|------|-------------|
| `00-create-collection.js` | Creates the `sources` collection with JSON Schema validation |
| `01-create-indexes.js`    | Creates indexes on `src_citation` and `src_id` |
