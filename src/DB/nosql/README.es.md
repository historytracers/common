# nosql

Este directorio contiene scripts NoSQL (MongoDB) para crear y mantener la base de datos del paquete `common`, análogos a la estructura SQL ANSI en [`src/DB/sql/`](../sql/).

## Colección

### `sources`

Almacena entradas de fuentes sin duplicados. Cada campo se corresponde con `HTSourceElement` definido en [`src/go/data-type.go`](../../go/data-type.go#L32-L38):

| Campo            | Tipo BSON   | Descripción                         | Campo Go    |
|------------------|-------------|-------------------------------------|-------------|
| `src_id`         | `binData`   | Identificador binario único         | `ID`        |
| `src_citation`   | `string`    | Texto de la cita de la fuente       | `Citation`  |
| `src_date`       | `string`    | Fecha asociada a la fuente          | `Date`      |
| `src_publish_date`| `string`   | Fecha de publicación de la fuente   | `PublishDate`|
| `src_url`        | `string`    | URL de la fuente                    | `URL`       |

Indexado en `src_citation` y `src_id` (único) para búsquedas más rápidas.

## Archivos

| Archivo | Descripción |
|---------|-------------|
| `00-create-collection.js` | Crea la colección `sources` con validación de esquema JSON |
| `01-create-indexes.js`    | Crea índices en `src_citation` y `src_id` |
