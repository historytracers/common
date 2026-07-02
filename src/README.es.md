# src

Este directorio contiene el código fuente del paquete `common`.

## Contenido

| Directorio | Lenguaje | Descripción |
|------------|----------|-------------|
| `go/`      | Go       | Tipos, utilidades y configuración compartidos usados en proyectos de History Tracers |

El directorio `go/` es un módulo de Go (`github.com/historytracers/common`) que expone el paquete `common` con estructuras de datos, tipos de configuración, marcas de tiempo y ayudantes de indexación para registros de clases y familias.

## Compilación

```bash
# Regenerar archivos de autotools (si configure.ac o Makefile.am cambiaron)
./bootstrap

# Configurar y compilar
./configure
make
```
