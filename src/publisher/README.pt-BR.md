# publisher

Publicador local para o conteúdo dos smartphones armazenado em
`src/smartphone/??-??/`.

Ele aplica as mesmas transformações usadas pelo publicador principal do
History Tracers (emoticons padrão, URL das páginas de `source_menu` e, quando um
arquivo é modificado no Git, seu carimbo de data/hora `last_update`) e grava o
JSON minificado no diretório de conteúdo.

## Compilação

O script auxiliar na raiz do repositório gera o sistema de compilação (quando
necessário), configura, compila o publicador e, em seguida, reescreve e
verifica o conteúdo dos smartphones:

```bash
./build-publisher.sh
```

Equivale a, a partir da raiz do repositório:

```bash
./bootstrap      # somente na primeira vez, gera o sistema de compilação
./configure
make publisher   # produz build/historytracers-publisher(.exe)
./build/historytracers-publisher -minify
./build/historytracers-publisher -validate
```

## Uso

Execute o binário a partir da raiz do repositório:

```bash
# Reescreve e minifica cada arquivo de smartphone em build/www/lang/<lang>/smartphone/
./build/historytracers-publisher -minify

# Verifica os arquivos JSON dos smartphones sem produzir saída
./build/historytracers-publisher -validate

# Mostra os diretórios em uso
./build/historytracers-publisher -compilation
```

### Opções

| Opção | Descrição |
| --- | --- |
| `-minify` | Reescreve e minifica todos os arquivos JSON dos smartphones no diretório de conteúdo. |
| `-validate` | Verifica os arquivos JSON dos smartphones sem produzir saída; termina com status diferente de zero quando um arquivo é inválido. |
| `-verbose` | Imprime mensagens de informação durante o processamento. |
| `-compilation` | Mostra os diretórios usados por esta compilação. |
| `-src` | Diretório que contém os arquivos de origem (padrão: `.`). |
| `-www` | Diretório para o conteúdo gerado (padrão: `build/www/`). |
| `-logfile` | Redireciona toda a saída para o arquivo informado. |
| `-db` | Caminho para o banco de dados SQLite de fontes acadêmicas (opcional). |

Quando o banco de dados de fontes opcional (`lang/sources/history_tracers.db`
por padrão) está presente, os dados de citação referenciados pelos arquivos dos
smartphones são carregados durante a publicação. O publicador funciona sem ele.
