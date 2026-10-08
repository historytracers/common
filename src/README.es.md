# src

Este directorio contiene el código fuente del paquete `common`.

## Contenido

| Directorio    | Lenguaje         | Descripción |
|---------------|------------------|-------------|
| `android/`    | Java (Android)   | Biblioteca Android con modelos de datos, utilidades de marca de tiempo y ayudantes de creación para aplicaciones de History Tracers |
| `go/`         | Go               | Tipos, utilidades y configuración compartidos usados en proyectos de History Tracers |
| `smartphone/` | JSON             | Contenido de las lecciones para teléfonos, un directorio por idioma |
| `publisher/`  | Go               | Publicador local que reescribe y minifica los archivos JSON de los teléfonos |

### android/

Módulo de biblioteca Android que expone POJOs Java (con serialización Gson) que reflejan las estructuras de datos de Go. Incluye `HTConfigBase`, todas las clases de modelo de contenido (`Family`, `ClassIdx`, `ClassTemplateFile`, `HTSourceFile`, etc.), `TimestampUtils` y `CreatorUtils`.

Compilar con Gradle:

```bash
./gradlew :common-android:assembleRelease
```

### go/

Módulo de Go (`github.com/historytracers/common`) que expone el paquete `common` con estructuras de datos, tipos de configuración, marcas de tiempo y ayudantes de indexación para registros de clases y familias.

Compilar con autotools:

```bash
./bootstrap
./configure
make
```

### smartphone/

Contenido de las lecciones para teléfonos, un directorio por idioma (`en-US/`,
`es-ES/`, `pt-BR/`) con un archivo JSON por lección.

### publisher/

Comando en Go que reescribe (normaliza) y minifica los archivos JSON de los
teléfonos.

Compílelo con el script auxiliar desde la raíz del repositorio:

```bash
./build-publisher.sh
```

Reescriba (normalice y minifique) todos los archivos JSON:

```bash
./build/historytracers-publisher -minify
```

Los archivos generados se escriben en `build/www/lang/<lang>/smartphone/`; los
archivos fuente no se modifican. Verifique los archivos JSON con:

```bash
./build/historytracers-publisher -validate
```

Genere los archivos de entrada para la síntesis de voz con:

```bash
./build/historytracers-publisher -audio
```

Los archivos de texto de audio se escriben en `audio/`.

Consulte [publisher/README.es.md](publisher/README.es.md) para más detalles.
