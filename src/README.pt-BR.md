# src

Este diretório contém o código-fonte do pacote `common`.

## Conteúdo

| Diretório    | Linguagem         | Descrição |
|--------------|-------------------|-----------|
| `android/`   | Java (Android)    | Biblioteca Android com modelos de dados, utilitários de carimbo de data/hora e auxiliares de criação para aplicativos History Tracers |
| `go/`        | Go                | Tipos, utilitários e configuração compartilhados usados em projetos do History Tracers |

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
