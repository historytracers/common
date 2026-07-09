# DB

Este diretório contém scripts SQL ANSI para criar e manter o banco de dados do pacote `common`.

## Banco de dados

O banco de dados `history_tracers` armazena registros de fontes sem duplicatas usados em projetos do History Tracers.

## Tabelas

### `sources`

Armazena entradas de fontes sem duplicatas. Cada campo mapeia para `HTSourceElement` definido em [`src/go/data-type.go`](../go/data-type.go#L32-L38):

| Coluna          | Tipo        | Descrição                          | Campo Go    |
|-----------------|-------------|------------------------------------|-------------|
| `src_id`        | `BINARY(16)`| Identificador único (chave primária)| `ID`        |
| `src_citation`  | `TEXT`      | Texto da citação da fonte          | `Citation`  |
| `src_date`      | `TEXT`      | Data associada à fonte             | `Date`      |
| `src_publish_date`| `TEXT`    | Data de publicação da fonte        | `PublishDate`|
| `src_url`       | `TEXT`      | URL da fonte                       | `URL`       |

Indexado em `src_citation` para buscas mais rápidas.

## Arquivos

| Arquivo | Descrição |
|---------|-----------|
| `00-create-database.sql`   | Cria o banco de dados `history_tracers` |
| `01-create-sources.sql`    | Cria a tabela `sources` |
| `02-index-sources-citation.sql` | Adiciona um índice em `src_citation` |
