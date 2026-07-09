# sql

Este diretório contém scripts SQL ANSI para criar e manter o banco de dados do pacote `common`. É análogo à estrutura NoSQL em [`src/DB/nosql/`](../nosql/).

## Banco de dados

O banco de dados `history_tracers` armazena registros de fontes sem duplicatas usados em projetos do History Tracers.

## Tabelas

### `sources`

Armazena entradas de fontes sem duplicatas. Cada campo mapeia para `HTSourceElement` definido em [`src/go/data-type.go`](../../go/data-type.go#L32-L38):

| Coluna          | Tipo        | Descrição                          | Campo Go    |
|-----------------|-------------|------------------------------------|-------------|
| `src_id`        | `BINARY(16)`| Identificador único (chave primária)| `ID`        |
| `src_citation`  | `TEXT`      | Texto da citação da fonte          | `Citation`  |
| `src_date`      | `TEXT`      | Data associada à fonte             | `Date`      |
| `src_publish_date`| `TEXT`    | Data de publicação da fonte        | `PublishDate`|
| `src_url`       | `TEXT`      | URL da fonte                       | `URL`       |

Indexado em `src_citation` para buscas mais rápidas.

### `file`

Armazena entradas de arquivos. Cada arquivo tem um identificador binário único e uma descrição:

| Coluna   | Tipo        | Descrição                            |
|----------|-------------|--------------------------------------|
| `fil_id` | `BINARY(16)`| Identificador único (chave primária) |
| `fil_desc`| `TEXT`     | Descrição do arquivo                 |

### `citation`

Define um relacionamento N para M entre `file` e `source` com um qualificador de tipo:

| Coluna   | Tipo        | Descrição                                                  |
|----------|-------------|------------------------------------------------------------|
| `fil_id` | `BINARY(16)`| Referência a uma entrada `file`                            |
| `src_id` | `BINARY(16)`| Referência a uma entrada `source`                          |
| `cit_type`| `TINYINT`  | Tipo de citação (0=Primary, 1=References, 2=Holy, 3=Social Media) |

As três colunas juntas formam a chave primária composta. Indexado em `fil_id` para buscas mais rápidas. Uma restrição `CHECK` limita `cit_type` aos valores 0–3.

## Arquivos

| Arquivo | Descrição |
|---------|-----------|
| `00-create-database.sql`   | Cria o banco de dados `history_tracers` |
| `01-create-sources.sql`    | Cria a tabela `sources` |
| `02-index-sources-citation.sql` | Adiciona um índice em `src_citation` |
| `03-create-file.sql`       | Cria a tabela `file` |
| `04-create-citation.sql`   | Cria a tabela `citation` |
