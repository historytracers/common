# nosql

Este diretório contém scripts NoSQL (MongoDB) para criar e manter o banco de dados do pacote `common`, análogos à estrutura SQL ANSI em [`src/DB/sql/`](../sql/).

## Coleção

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

## Arquivos

| Arquivo | Descrição |
|---------|-----------|
| `00-create-collection.js` | Cria a coleção `sources` com validação de esquema JSON |
| `01-create-indexes.js`    | Cria índices em `src_citation` e `src_id` |
