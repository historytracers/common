[![Contributor Covenant](https://img.shields.io/badge/Contributor%20Covenant-2.1-4baaaa.svg)](CODE_OF_CONDUCT.md)
![CodeRabbit Pull Request Reviews](https://img.shields.io/coderabbit/prs/github/historytracers/common?utm_source=oss&utm_medium=github&utm_campaign=historytracers%2Fcommon&labelColor=171717&color=FF570A&link=https%3A%2F%2Fcoderabbit.ai&label=CodeRabbit+Reviews)

# Common

Código comum compartilhado para a organização [History Tracers](https://github.com/historytracers).

Este repositório contém bibliotecas, utilitários, tipos e configurações compartilhados entre vários projetos do History Tracers, incluindo:

- [historytracers/historytracers](https://github.com/historytracers/historytracers) — o aplicativo principal do History Tracers
- Outros projetos sob a organização [historytracers](https://github.com/historytracers)

## Propósito

Manter o código comum em um único lugar garante consistência, reduz duplicação e facilita a manutenção da lógica compartilhada em todo o ecossistema.

## Uso

Cada projeto consumidor referencia este repositório como uma dependência. Consulte a documentação do respectivo projeto para instruções de instalação e uso.

## Desenvolvimento

### Compilação

```bash
./bootstrap          # somente na primeira vez, gera o sistema de compilação
./configure
make                 # compila o pacote Go compartilhado e o publicador local
```

O publicador local é compilado como `build/historytracers-publisher(.exe)`. Para
bootstrap, configurar, compilar, reescrever e verificar em uma única etapa:

```bash
./build-publisher.sh
```

### Reescrever os arquivos JSON

O conteúdo dos smartphones em `src/smartphone/??-??/` é reescrito
(normalizado) e minificado com `-minify`:

```bash
./build/historytracers-publisher -minify
```

Os arquivos minificados são gravados em `build/www/lang/<lang>/smartphone/`;
os arquivos de origem nunca são modificados. Para verificar todos os arquivos
JSON dos smartphones (termina com status diferente de zero em caso de erro):

```bash
./build/historytracers-publisher -validate
```

### Gerar áudio

Para gerar os arquivos de entrada para a síntese de voz de cada tela dos
smartphones:

```bash
./build/historytracers-publisher -audio
```

Os arquivos de texto simples são gravados em `audio/`.

Consulte [src/publisher/README.pt-BR.md](src/publisher/README.pt-BR.md) para a
documentação completa.

## Contribuindo

Veja [CONTRIBUTING.md](./CONTRIBUTING.md) se presente; caso contrário, consulte o guia de contribuição no repositório principal.

Todas as contribuições devem seguir o [Código de Conduta](./CODE_OF_CONDUCT.md).

## Licença

Veja [LICENSE](./LICENSE) se presente; caso contrário, consulte a licença no repositório principal.
