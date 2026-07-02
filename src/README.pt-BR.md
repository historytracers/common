# src

Este diretório contém o código-fonte do pacote `common`.

## Conteúdo

| Diretório | Linguagem | Descrição |
|-----------|-----------|-----------|
| `go/`     | Go        | Tipos, utilitários e configuração compartilhados usados em projetos do History Tracers |

O diretório `go/` é um módulo Go (`github.com/historytracers/common`) que expõe o pacote `common` com estruturas de dados, tipos de configuração, carimbos de data/hora e auxiliares de indexação para registros de classes e famílias.

## Compilação

```bash
# Regenerar arquivos do autotools (se configure.ac ou Makefile.am mudaram)
./bootstrap

# Configurar e compilar
./configure
make
```
