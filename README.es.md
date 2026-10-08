[![Contributor Covenant](https://img.shields.io/badge/Contributor%20Covenant-2.1-4baaaa.svg)](CODE_OF_CONDUCT.md)
![CodeRabbit Pull Request Reviews](https://img.shields.io/coderabbit/prs/github/historytracers/common?utm_source=oss&utm_medium=github&utm_campaign=historytracers%2Fcommon&labelColor=171717&color=FF570A&link=https%3A%2F%2Fcoderabbit.ai&label=CodeRabbit+Reviews)

# Common

Código común compartido para la organización [History Tracers](https://github.com/historytracers).

Este repositorio contiene librerías, utilidades, tipos y configuración compartidos entre varios proyectos de History Tracers, incluyendo:

- [historytracers/historytracers](https://github.com/historytracers/historytracers) — la aplicación principal de History Tracers
- Otros proyectos bajo la organización [historytracers](https://github.com/historytracers)

## Propósito

Mantener el código común en un solo lugar garantiza consistencia, reduce duplicación y facilita el mantenimiento de la lógica compartida en todo el ecosistema.

## Uso

Cada proyecto consumidor referencia este repositorio como una dependencia. Consulte la documentación del proyecto respectivo para instrucciones de instalación y uso.

## Desarrollo

### Compilación

```bash
./bootstrap          # solo la primera vez, genera el sistema de compilación
./configure
make                 # compila el paquete Go compartido y el publicador local
```

El publicador local se compila como `build/historytracers-publisher(.exe)`. Para
bootstrap, configurar, compilar, reescribir y verificar en un solo paso:

```bash
./build-publisher.sh
```

### Reescribir los archivos JSON

El contenido de los teléfonos en `src/smartphone/??-??/` se reescribe
(normaliza) y minifica con `-minify`:

```bash
./build/historytracers-publisher -minify
```

Los archivos minificados se escriben en `build/www/lang/<lang>/smartphone/`;
los archivos fuente nunca se modifican. Para verificar todos los archivos JSON
de los teléfonos (termina con un estado distinto de cero si hay errores):

```bash
./build/historytracers-publisher -validate
```

### Generar audio

Para generar los archivos de entrada para la síntesis de voz de cada pantalla
de los teléfonos:

```bash
./build/historytracers-publisher -audio
```

Los archivos de texto plano se escriben en `audio/`.

Consulte [src/publisher/README.es.md](src/publisher/README.es.md) para la
documentación completa.

## Contribuir

Vea [CONTRIBUTING.md](./CONTRIBUTING.md) si está presente; de lo contrario, consulte la guía de contribución en el repositorio principal.

Todas las contribuciones deben seguir el [Código de Conducta](./CODE_OF_CONDUCT.md).

## Licencia

Vea [LICENSE](./LICENSE) si está presente; de lo contrario, consulte la licencia en el repositorio principal.
