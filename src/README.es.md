# src

Este directorio contiene el código fuente del paquete `common`.

## Contenido

| Directorio    | Lenguaje         | Descripción |
|---------------|------------------|-------------|
| `android/`    | Java (Android)   | Biblioteca Android con modelos de datos, utilidades de marca de tiempo y ayudantes de creación para aplicaciones de History Tracers |
| `go/`         | Go               | Tipos, utilidades y configuración compartidos usados en proyectos de History Tracers |

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
