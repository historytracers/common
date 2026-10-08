# publisher

Publicador local para o conteúdo dos smartphones armazenado em
`src/smartphone/??-??/`.

Ele aplica as mesmas transformações usadas pelo publicador principal do
History Tracers (emoticons padrão, URL das páginas de `source_menu` e, quando um
arquivo é modificado no Git, seu carimbo de data/hora `last_update`), grava o
JSON minificado no diretório de conteúdo e gera os arquivos de entrada para a
síntese de voz.

## Compilação

O script auxiliar na raiz do repositório gera o sistema de compilação (quando
necessário), configura, compila o publicador e, em seguida, reescreve, gera o
áudio e verifica o conteúdo dos smartphones:

```bash
./build-publisher.sh
```

Equivale a, a partir da raiz do repositório:

```bash
./bootstrap      # somente na primeira vez, gera o sistema de compilação
./configure
make publisher   # produz build/historytracers-publisher(.exe)
./build/historytracers-publisher -minify
./build/historytracers-publisher -audio
./build/historytracers-publisher -validate
```

## Uso

Execute o binário a partir da raiz do repositório:

```bash
# Reescreve e minifica cada arquivo de smartphone em build/www/lang/<lang>/smartphone/
./build/historytracers-publisher -minify

# Gera os arquivos de entrada para a síntese de voz em audio/
./build/historytracers-publisher -audio

# Verifica os arquivos JSON dos smartphones sem produzir saída
./build/historytracers-publisher -validate

# Mostra os diretórios em uso
./build/historytracers-publisher -compilation
```

### Opções

| Opção | Descrição |
| --- | --- |
| `-minify` | Reescreve e minifica todos os arquivos JSON dos smartphones no diretório de conteúdo. |
| `-audio` | Gera os arquivos de entrada para a síntese de voz de cada tela no diretório de áudio. |
| `-validate` | Verifica os arquivos JSON dos smartphones sem produzir saída; termina com status diferente de zero quando um arquivo é inválido. |
| `-verbose` | Imprime mensagens de informação durante o processamento. |
| `-compilation` | Mostra os diretórios usados por esta compilação. |
| `-src` | Diretório que contém os arquivos de origem (padrão: `.`). |
| `-www` | Diretório para o conteúdo gerado (padrão: `build/www/`). |
| `-audiodir` | Diretório para os arquivos de texto de áudio gerados (padrão: `audio/`). |
| `-logfile` | Redireciona toda a saída para o arquivo informado. |
| `-db` | Caminho para o banco de dados SQLite de fontes acadêmicas (opcional). |

Quando o banco de dados de fontes opcional (`lang/sources/history_tracers.db`
por padrão) está presente, os dados de citação referenciados pelos arquivos dos
smartphones são carregados durante a publicação. O publicador funciona sem ele.

## Arquivos de áudio

`-audio` grava um arquivo de texto simples por tela, chamado
`audio/<uuid-lição>_<uuid-tela>_<lang>.txt`. Os arquivos destinam-se a um motor
de síntese de voz como o [Piper](https://github.com/rhasspy/piper). Marcação
HTML e Markdown, emoticons e URLs de redes sociais são simplificados para que o
texto seja lido naturalmente. O diretório de saída pode ser alterado com
`-audiodir`.

## Reutilizar o pacote

A lógica de reescrita, minificação e áudio está no pacote importável
`historytracers-publisher/smartphone`, de modo que outros projetos do History
Tracers podem chamá-la diretamente:

```go
import "historytracers-publisher/smartphone"

cfg := smartphone.Config{
    SrcPath:     "/caminho/para/common",
    ContentPath: "/caminho/de/saida",
    AudioPath:   "/caminho/de/audio",
}

if err := smartphone.HTMinifyAllFiles(cfg); err != nil {
    // tratar o erro
}

if err := smartphone.HTGenerateAudio(cfg); err != nil {
    // tratar o erro
}

if invalid := smartphone.HTValidateSMGameFormats(cfg); invalid > 0 {
    // tratar os arquivos inválidos
}
```

Os consumidores referenciam este módulo com uma diretiva `replace` (o mesmo
padrão usado para `github.com/historytracers/common`):

```
require historytracers-publisher v0.0.0

replace historytracers-publisher => ../common/src/publisher
```

O comando deste diretório é apenas um invólucro fino sobre o pacote.
