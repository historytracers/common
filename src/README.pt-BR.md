# src

Este diretório contém o código-fonte do pacote `common`.

## Conteúdo

| Diretório    | Linguagem         | Descrição |
|--------------|-------------------|-----------|
| `android/`   | Java (Android)    | Biblioteca Android com modelos de dados, utilitários de carimbo de data/hora e auxiliares de criação para aplicativos History Tracers |
| `go/`        | Go                | Tipos, utilitários e configuração compartilhados usados em projetos do History Tracers |
| `smartphone/` | JSON              | Conteúdo das lições para smartphones, um diretório por idioma |
| `publisher/`  | Go                | Publicador local que reescreve e minifica os arquivos JSON dos smartphones |

### android/

Módulo de biblioteca Android que expõe POJOs Java (com serialização Gson) espelhando as estruturas de dados Go. Inclui `HTConfigBase`, todas as classes de modelo de conteúdo (`Family`, `ClassIdx`, `ClassTemplateFile`, `HTSourceFile`, etc.), `TimestampUtils` e `CreatorUtils`.

Compilar com Gradle:

```bash
./gradlew :common-android:assembleRelease
```

### go/

Módulo Go (`github.com/historytracers/common`) que expõe o pacote `common` com estruturas de dados, tipos de configuração, carimbos de data/hora e auxiliares de indexação para registros de classes e famílias.

Compilar com autotools:

```bash
./bootstrap
./configure
make
```

### smartphone/

Conteúdo das lições para smartphones, um diretório por idioma (`en-US/`,
`es-ES/`, `pt-BR/`) com um arquivo JSON por lição.

### publisher/

Comando em Go que reescreve (normaliza) e minifica os arquivos JSON dos
smartphones.

Compile-o com o script auxiliar a partir da raiz do repositório:

```bash
./build-publisher.sh
```

Reescreva (normalize e minifique) todos os arquivos JSON:

```bash
./build/historytracers-publisher -minify
```

Os arquivos gerados são gravados em `build/www/lang/<lang>/smartphone/`; os
arquivos de origem não são modificados. Verifique os arquivos JSON com:

```bash
./build/historytracers-publisher -validate
```

Consulte [publisher/README.pt-BR.md](publisher/README.pt-BR.md) para mais
detalhes.
