# publisher

Publicador local para el contenido de los teléfonos almacenado en
`src/smartphone/??-??/`.

Aplica las mismas transformaciones que el publicador principal de History
Tracers (emoticonos predeterminados, URL de las páginas de `source_menu` y, cuando
un archivo se modifica en Git, su marca de tiempo `last_update`) y escribe el
JSON minificado en el directorio de contenido.

## Compilación

El script auxiliar en la raíz del repositorio genera el sistema de compilación
(cuando es necesario), configura, compila el publicador y luego reescribe y
verifica el contenido de los teléfonos:

```bash
./build-publisher.sh
```

Equivale a, desde la raíz del repositorio:

```bash
./bootstrap      # solo la primera vez, genera el sistema de compilación
./configure
make publisher   # produce build/historytracers-publisher(.exe)
./build/historytracers-publisher -minify
./build/historytracers-publisher -validate
```

## Uso

Ejecute el binario desde la raíz del repositorio:

```bash
# Reescribe y minifica cada archivo de teléfono en build/www/lang/<lang>/smartphone/
./build/historytracers-publisher -minify

# Verifica los archivos JSON de los teléfonos sin producir salida
./build/historytracers-publisher -validate

# Muestra los directorios en uso
./build/historytracers-publisher -compilation
```

### Opciones

| Opción | Descripción |
| --- | --- |
| `-minify` | Reescribe y minifica todos los archivos JSON de los teléfonos en el directorio de contenido. |
| `-validate` | Verifica los archivos JSON de los teléfonos sin producir salida; termina con un estado distinto de cero cuando un archivo no es válido. |
| `-verbose` | Imprime mensajes de información durante el procesamiento. |
| `-compilation` | Muestra los directorios usados por esta compilación. |
| `-src` | Directorio que contiene los archivos fuente (predeterminado: `.`). |
| `-www` | Directorio para el contenido generado (predeterminado: `build/www/`). |
| `-logfile` | Redirige toda la salida al archivo indicado. |
| `-db` | Ruta a la base de datos SQLite de fuentes académicas (opcional). |

Cuando la base de datos de fuentes opcional (`lang/sources/history_tracers.db`
de forma predeterminada) está presente, los datos de citas referenciados por
los archivos de los teléfonos se cargan durante la publicación. El publicador
funciona sin ella.

## Reutilizar el paquete

La lógica de reescritura y minificación está en el paquete importable
`historytracers-publisher/smartphone`, de modo que otros proyectos de History
Tracers pueden llamarla directamente:

```go
import "historytracers-publisher/smartphone"

cfg := smartphone.Config{
    SrcPath:     "/ruta/a/common",
    ContentPath: "/ruta/de/salida",
}

if err := smartphone.HTMinifyAllFiles(cfg); err != nil {
    // manejar el error
}

if invalid := smartphone.HTValidateSMGameFormats(cfg); invalid > 0 {
    // manejar los archivos inválidos
}
```

Los consumidores referencian este módulo con una directiva `replace` (el mismo
patrón usado para `github.com/historytracers/common`):

```
require historytracers-publisher v0.0.0

replace historytracers-publisher => ../common/src/publisher
```

El comando de este directorio es solo un envoltorio delgado sobre el paquete.
