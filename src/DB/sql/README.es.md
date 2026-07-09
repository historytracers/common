# sql

Este directorio contiene scripts SQL ANSI para crear y mantener la base de datos del paquete `common`. Es análogo a la estructura NoSQL en [`src/DB/nosql/`](../nosql/).

## Base de datos

La base de datos `history_tracers` almacena registros de fuentes sin duplicados utilizados en proyectos de History Tracers.

## Tablas

### `sources`

Almacena entradas de fuentes sin duplicados. Cada campo se corresponde con `HTSourceElement` definido en [`src/go/data-type.go`](../../go/data-type.go#L32-L38):

| Columna           | Tipo        | Descripción                         | Campo Go    |
|-------------------|-------------|-------------------------------------|-------------|
| `src_id`          | `BINARY(16)`| Identificador único (clave primaria) | `ID`        |
| `src_citation`    | `TEXT`      | Texto de la cita de la fuente       | `Citation`  |
| `src_date`        | `TEXT`      | Fecha asociada a la fuente          | `Date`      |
| `src_publish_date`| `TEXT`      | Fecha de publicación de la fuente   | `PublishDate`|
| `src_url`         | `TEXT`      | URL de la fuente                    | `URL`       |

Indexado en `src_citation` para búsquedas más rápidas.

### `file`

Almacena entradas de archivos. Cada archivo tiene un identificador binario único y una descripción:

| Columna   | Tipo        | Descripción                          |
|-----------|-------------|--------------------------------------|
| `fil_id`  | `BINARY(16)`| Identificador único (clave primaria) |
| `fil_desc`| `TEXT`      | Descripción del archivo              |

### `citation`

Define una relación N a M entre `file` y `source` con un calificador de tipo:

| Columna   | Tipo        | Descripción                                               |
|-----------|-------------|-----------------------------------------------------------|
| `fil_id`  | `BINARY(16)`| Referencia a la entrada `file`                            |
| `src_id`  | `BINARY(16)`| Referencia a la entrada `source`                          |
| `cit_type`| `TINYINT`   | Tipo de cita (0=Primary, 1=References, 2=Holy, 3=Social Media) |

Las tres columnas juntas forman la clave primaria compuesta. Indexado en `fil_id` para búsquedas más rápidas. Una restricción `CHECK` limita `cit_type` a valores 0–3.

## Archivos

| Archivo | Descripción |
|---------|-------------|
| `00-create-database.sql`   | Crea la base de datos `history_tracers` |
| `01-create-sources.sql`    | Crea la tabla `sources` |
| `02-index-sources-citation.sql` | Añade un índice en `src_citation` |
| `03-create-file.sql`       | Crea la tabla `file` |
| `04-create-citation.sql`   | Crea la tabla `citation` |
