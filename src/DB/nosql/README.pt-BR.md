# nosql

Este diretório contém scripts NoSQL (MongoDB) para criar e manter o banco de dados do pacote `common`, análogos à estrutura SQL ANSI em [`src/DB/sql/`](../sql/).

## Coleções

### `sources`

Armazena entradas de fontes sem duplicatas. Cada campo mapeia para `HTSourceElement` definido em [`src/go/data-type.go`](../../go/data-type.go#L32-L38):

| Campo           | Tipo BSON   | Descrição                          | Campo Go    |
|-----------------|-------------|------------------------------------|-------------|
| `src_id`        | `binData`   | Identificador binário único        | `ID`        |
| `src_citation`  | `string`    | Texto da citação da fonte          | `Citation`  |
| `src_date`      | `string`    | Data associada à fonte             | `Date`      |
| `src_publish_date`| `string`  | Data de publicação da fonte        | `PublishDate`|
| `src_url`       | `string`    | URL da fonte                       | `URL`       |

Indexado em `src_citation` e `src_id` (único) para buscas mais rápidas.

### `file`

Armazena entradas de arquivos:

| Campo     | Tipo BSON | Descrição                       |
|-----------|-----------|---------------------------------|
| `fil_id`  | `binData` | Identificador binário único     |
| `fil_desc`| `string`  | Descrição do arquivo            |

### `citation`

Define um relacionamento N para M entre `file` e `source` com um qualificador de tipo:

| Campo     | Tipo BSON | Descrição                                                   |
|-----------|-----------|-------------------------------------------------------------|
| `fil_id`  | `binData` | Referência a uma entrada `file`                             |
| `src_id`  | `binData` | Referência a uma entrada `source`                           |
| `cit_type`| `int`     | Tipo de citação (0=Primary, 1=References, 2=Holy, 3=Social Media) |

Validado com `$jsonSchema` garantindo que `cit_type` seja um de `[0, 1, 2, 3]`.

## Arquivos

| Arquivo | Descrição |
|---------|-----------|
| `00-create-collection.js` | Cria a coleção `sources` com validação de esquema JSON |
| `01-create-indexes.js`    | Cria índices em `src_citation` e `src_id` |
| `02-create-file.js`       | Cria a coleção `file` com validação de esquema JSON |
| `03-create-citation.js`   | Cria a coleção `citation` com validação de esquema JSON |
