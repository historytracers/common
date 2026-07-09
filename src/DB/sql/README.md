# sql

This directory contains ANSI SQL scripts to create and maintain the database for the `common` package. It mirrors the NoSQL structure in [`src/DB/nosql/`](../nosql/).

## Database

The database `history_tracers` stores deduplicated source records used across History Tracers projects.

## Tables

### `sources`

Stores source entries without duplicates. Each field maps to `HTSourceElement` defined in [`src/go/data-type.go`](../../go/data-type.go#L32-L38):

| Column            | Type        | Description                        | Go field      |
|-------------------|-------------|------------------------------------|---------------|
| `src_id`          | `BINARY(16)`| Unique identifier (primary key)    | `ID`          |
| `src_citation`    | `TEXT`      | Source citation text               | `Citation`    |
| `src_date`        | `TEXT`      | Date associated with the source    | `Date`        |
| `src_publish_date`| `TEXT`      | Publication date of the source     | `PublishDate` |
| `src_url`         | `TEXT`      | Source URL                         | `URL`         |

Indexed on `src_citation` for faster search lookups.

### `file`

Stores file entries. Each file has a unique binary identifier and a description:

| Column   | Type        | Description                     |
|----------|-------------|---------------------------------|
| `fil_id` | `BINARY(16)`| Unique identifier (primary key) |
| `fil_desc`| `TEXT`     | File description                |

### `citation`

Defines an N-to-M relationship between `file` and `source` with a type qualifier:

| Column    | Type        | Description                                      |
|-----------|-------------|--------------------------------------------------|
| `fil_id`  | `BINARY(16)`| Reference to `file` entry                        |
| `src_id`  | `BINARY(16)`| Reference to `source` entry                      |
| `cit_type`| `TINYINT`   | Type of citation (0=Primary, 1=References, 2=Holy, 3=Social Media) |

The three columns together form the composite primary key. Indexed on `fil_id` for faster lookups. A `CHECK` constraint limits `cit_type` to values 0–3.

## Files

| File | Description |
|------|-------------|
| `00-create-database.sql`   | Creates the `history_tracers` database |
| `01-create-sources.sql`    | Creates the `sources` table |
| `02-index-sources-citation.sql` | Adds an index on `src_citation` |
| `03-create-file.sql`       | Creates the `file` table |
| `04-create-citation.sql`   | Creates the `citation` table |
