# DB

This directory contains ANSI SQL scripts to create and maintain the database for the `common` package.

## Database

The database `history_tracers` stores deduplicated source records used across History Tracers projects.

## Tables

### `sources`

Stores source entries without duplicates. Each field maps to `HTSourceElement` defined in [`src/go/data-type.go`](../go/data-type.go#L32-L38):

| Column            | Type        | Description                        | Go field      |
|-------------------|-------------|------------------------------------|---------------|
| `src_id`          | `BINARY(16)`| Unique identifier (primary key)    | `ID`          |
| `src_citation`    | `TEXT`      | Source citation text               | `Citation`    |
| `src_date`        | `TEXT`      | Date associated with the source    | `Date`        |
| `src_publish_date`| `TEXT`      | Publication date of the source     | `PublishDate` |
| `src_url`         | `TEXT`      | Source URL                         | `URL`         |

Indexed on `src_citation` for faster search lookups.

## Files

| File | Description |
|------|-------------|
| `00-create-database.sql`   | Creates the `history_tracers` database |
| `01-create-sources.sql`    | Creates the `sources` table |
| `02-index-sources-citation.sql` | Adds an index on `src_citation` |
