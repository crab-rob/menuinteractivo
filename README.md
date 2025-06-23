# MenuInteractivo

MenuInteractivo es una aplicación de terminal escrita en Go que permite ejecutar comandos de GNU/Linux y ver su explicación detallada en tiempo real. Es ideal como herramienta educativa y de consulta rápida para usuarios que desean aprender o recordar el funcionamiento de comandos comunes del sistema.

## Características

- Ejecuta comandos de shell directamente desde la interfaz.
- Muestra la salida del comando en un panel izquierdo.
- Muestra la descripción, uso y detalles del comando en un panel derecho, extraídos de una base de datos JSON.
- Soporte para decenas de comandos populares de Linux.
- Interfaz de doble panel ajustable al tamaño de la terminal.
- Fácil de extender agregando nuevos comandos a `db.json`.

## Instalación

Clona el repositorio y compila el proyecto:

```sh
git clone https://github.com/jclumbiarres/menuinteractivo.git
cd menuinteractivo
go build
```

## Uso

1. Ejecuta el binario:

   ```sh
   ./menuinteractivo
   ```

2. Escribe cualquier comando de Linux (por ejemplo, `ls`, `cat`, `grep`, etc.) y presiona Enter.
3. Verás la salida del comando a la izquierda y su explicación a la derecha.
4. Escribe `exit` para salir.

## Base de datos de comandos

La información de los comandos se encuentra en el archivo [`db.json`](db.json). Puedes agregar o modificar comandos editando este archivo.

## Ejemplo de uso

```
SHELL> ls
$ ls
archivo1.txt  archivo2.txt

Comando: ls

Lista archivos y directorios del directorio actual.

Muestra el contenido de un directorio. Ejemplo: ls -lah. Flags: -l (listado largo), -a (ocultos), -h (legible).
```

## Licencia

MIT

---

Desarrollado por [@jclumbiarres](https://github.com/jclumbiarres)
