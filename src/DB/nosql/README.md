# nosql

This directory contains NoSQL scripts (MongoDB) to create and maintain the database for the `common` package, analogous to the ANSI SQL structure in [`src/DB/sql/`](../sql/).

## Collections

### `sources`

Stores source entries without duplicates. Each field maps to `HTSourceElement` defined in [`src/go/data-type.go`](../../go/data-type.go#L32-L38):

| Field             | BSON Type   | Description                        | Go field      |
|-------------------|-------------|------------------------------------|---------------|
| `src_id`          | `binData`   | Unique binary identifier           | `ID`          |
| `sfo_id`          | `binData`   | UUID of citation format            | `SfoID`       |
| `src_citation`    | `string`    | Source citation text               | `Citation`    |
| `src_date`        | `string`    | Date associated with the source    | `Date`        |
| `src_publish_date`| `string`    | Publication date of the source     | `PublishDate` |
| `src_url`         | `string`    | Source URL                         | `URL`         |

Indexed on `src_citation` and `src_id` (unique) for faster lookups.

### `source_format`

Defines citation formats referenced by `sources`:

| Field           | BSON Type | Description                     |
|-----------------|-----------|---------------------------------|
| `sfo_id`        | `binData` | Unique binary identifier        |
| `sfo_name`      | `string`  | Format name                     |
| `sfo_description`| `string` | Format description              |

### `files`

Stores file entries:

| Field     | BSON Type | Description                     |
|-----------|-----------|---------------------------------|
| `fil_id`  | `binData` | Unique binary identifier        |
| `fil_desc`| `string`  | File description                |

### `citation`

Defines an N-to-M relationship between `files` and `source` with a type qualifier:

| Field     | BSON Type | Description                                               |
|-----------|-----------|-----------------------------------------------------------|
| `fil_id`  | `binData` | Reference to a `files` entry                               |
| `src_id`  | `binData` | Reference to a `source` entry                             |
| `cit_type`| `int`     | Type of citation (0=Primary, 1=References, 2=Holy, 3=Social Media) |

Validated with `$jsonSchema` ensuring `cit_type` is one of `[0, 1, 2, 3]`.

## Files

| File | Description |
|------|-------------|
| `00-create-collection.js`      | Creates the `sources` collection with JSON Schema validation |
| `01-create-source-format.js`  | Creates the `source_format` collection with JSON Schema validation |
| `02-create-indexes.js`        | Creates indexes on `src_citation` and `src_id` |
| `03-create-files.js`           | Creates the `files` collection with JSON Schema validation |
| `04-create-citation.js`        | Creates the `citation` collection with JSON Schema validation |
