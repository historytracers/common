# DB

Este directorio contiene scripts SQL ANSI para crear y mantener la base de datos del paquete `common`.

## Base de datos

La base de datos `history_tracers` almacena registros de fuentes sin duplicados utilizados en proyectos de History Tracers.

## Tablas

### `sources`

Almacena entradas de fuentes sin duplicados. Cada campo se corresponde con `HTSourceElement` definido en [`src/go/data-type.go`](../go/data-type.go#L32-L38):

| Columna           | Tipo        | Descripción                         | Campo Go    |
|-------------------|-------------|-------------------------------------|-------------|
| `src_id`          | `BINARY(16)`| Identificador único (clave primaria) | `ID`        |
| `src_citation`    | `TEXT`      | Texto de la cita de la fuente       | `Citation`  |
| `src_date`        | `TEXT`      | Fecha asociada a la fuente          | `Date`      |
| `src_publish_date`| `TEXT`      | Fecha de publicación de la fuente   | `PublishDate`|
| `src_url`         | `TEXT`      | URL de la fuente                    | `URL`       |

Indexado en `src_citation` para búsquedas más rápidas.

## Archivos

| Archivo | Descripción |
|---------|-------------|
| `00-create-database.sql`   | Crea la base de datos `history_tracers` |
| `01-create-sources.sql`    | Crea la tabla `sources` |
| `02-index-sources-citation.sql` | Añade un índice en `src_citation` |
