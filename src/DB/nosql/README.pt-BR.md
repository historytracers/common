# nosql

Este diretório contém scripts NoSQL (MongoDB) para criar e manter o banco de dados do pacote `common`, análogos à estrutura SQL ANSI em [`src/DB/sql/`](../sql/).

## Coleções

### `sources`

Armazena entradas de fontes sem duplicatas. Cada campo mapeia para `HTSourceElement` definido em [`src/go/data-type.go`](../../go/data-type.go#L32-L38):

| Campo           | Tipo BSON   | Descrição                          | Campo Go    |
|-----------------|-------------|------------------------------------|-------------|
| `src_id`        | `binData`   | Identificador binário único        | `ID`        |
| `sfo_id`         | `binData`   | UUID do formato de citação         | `SfoID`     |
| `src_citation`   | `string`    | Texto da citação da fonte          | `Citation`  |
| `src_date`      | `string`    | Data associada à fonte             | `Date`      |
| `src_publish_date`| `string`  | Data de publicação da fonte        | `PublishDate`|
| `src_url`       | `string`    | URL da fonte                       | `URL`       |

Indexado em `src_citation` e `src_id` (único) para buscas mais rápidas.

### `source_format`

Define os formatos de citação referenciados por `sources`:

| Campo           | Tipo BSON | Descrição                       |
|-----------------|-----------|---------------------------------|
| `sfo_id`        | `binData` | Identificador binário único     |
| `sfo_name`      | `string`  | Nome do formato                 |
| `sfo_description`| `string` | Descrição do formato            |

### `files`

Armazena entradas de arquivos:

| Campo     | Tipo BSON | Descrição                       |
|-----------|-----------|---------------------------------|
| `fil_id`  | `binData` | Identificador binário único     |
| `fil_desc`| `string`  | Descrição do arquivo            |

### `citation`

Define um relacionamento N para M entre `files` e `source` com um qualificador de tipo:

| Campo     | Tipo BSON | Descrição                                                   |
|-----------|-----------|-------------------------------------------------------------|
| `fil_id`  | `binData` | Referência a uma entrada `files`                             |
| `src_id`  | `binData` | Referência a uma entrada `source`                           |
| `cit_type`| `int`     | Tipo de citação (0=Primary, 1=References, 2=Holy, 3=Social Media) |

Validado com `$jsonSchema` garantindo que `cit_type` seja um de `[0, 1, 2, 3]`.

## Arquivos

| Arquivo | Descrição |
|---------|-----------|
| `00-create-collection.js`      | Cria a coleção `sources` com validação de esquema JSON |
| `01-create-source-format.js`  | Cria a coleção `source_format` com validação de esquema JSON |
| `02-create-indexes.js`        | Cria índices em `src_citation` e `src_id` |
| `03-create-files.js`           | Cria a coleção `files` com validação de esquema JSON |
| `04-create-citation.js`        | Cria a coleção `citation` com validação de esquema JSON |
