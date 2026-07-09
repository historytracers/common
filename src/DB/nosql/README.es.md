# nosql

Este directorio contiene scripts NoSQL (MongoDB) para crear y mantener la base de datos del paquete `common`, análogos a la estructura SQL ANSI en [`src/DB/sql/`](../sql/).

## Colecciones

### `sources`

Almacena entradas de fuentes sin duplicados. Cada campo se corresponde con `HTSourceElement` definido en [`src/go/data-type.go`](../../go/data-type.go#L32-L38):

| Campo            | Tipo BSON   | Descripción                         | Campo Go    |
|------------------|-------------|-------------------------------------|-------------|
| `src_id`         | `binData`   | Identificador binario único         | `ID`        |
| `sfo_id`         | `binData`   | UUID del formato de cita            | `SfoID`     |
| `src_citation`   | `string`    | Texto de la cita de la fuente       | `Citation`  |
| `src_date`       | `string`    | Fecha asociada a la fuente          | `Date`      |
| `src_publish_date`| `string`   | Fecha de publicación de la fuente   | `PublishDate`|
| `src_url`        | `string`    | URL de la fuente                    | `URL`       |

Indexado en `src_citation` y `src_id` (único) para búsquedas más rápidas.

### `source_format`

Define los formatos de cita referenciados por `sources`:

| Campo           | Tipo BSON | Descripción                     |
|-----------------|-----------|---------------------------------|
| `sfo_id`        | `binData` | Identificador binario único     |
| `sfo_name`      | `string`  | Nombre del formato              |
| `sfo_description`| `string` | Descripción del formato         |

### `files`

Almacena entradas de archivos:

| Campo     | Tipo BSON | Descripción                      |
|-----------|-----------|----------------------------------|
| `fil_id`  | `binData` | Identificador binario único      |
| `fil_desc`| `string`  | Descripción del archivo          |

### `citation`

Define una relación N a M entre `files` y `source` con un calificador de tipo:

| Campo     | Tipo BSON | Descripción                                                |
|-----------|-----------|------------------------------------------------------------|
| `fil_id`  | `binData` | Referencia a una entrada `files`                            |
| `src_id`  | `binData` | Referencia a una entrada `source`                          |
| `cit_type`| `int`     | Tipo de cita (0=Primary, 1=References, 2=Holy, 3=Social Media) |

Validado con `$jsonSchema` asegurando que `cit_type` sea uno de `[0, 1, 2, 3]`.

## Archivos

| Archivo | Descripción |
|---------|-------------|
| `00-create-collection.js`      | Crea la colección `sources` con validación de esquema JSON |
| `01-create-source-format.js`  | Crea la colección `source_format` con validación de esquema JSON |
| `02-create-indexes.js`        | Crea índices en `src_citation` y `src_id` |
| `03-create-files.js`           | Crea la colección `files` con validación de esquema JSON |
| `04-create-citation.js`        | Crea la colección `citation` con validación de esquema JSON |
