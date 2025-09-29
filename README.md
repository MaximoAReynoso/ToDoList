# ToDoList

Esta aplicación, como se puede ver, es una lista de tareas.

## Interfaz

La interfaz de la aplicación es la siguiente: Está su título y el espacio donde se van a listar los elementos de una tabla con tareas, marcadas como pendientes o 
completadas. <br/>
    Cada fila tiene sus elementos, estos son el ID para identificar el elemento, su título, descripción, y si fue completada o no, siendo por defecto la tarea marcada como no completada. <br/>
    Se agregaron al diseño un icono acorde a la aplicacion (visible al levantar el server) y algunos estilos.

## Estructura de la pagina

Toda la aplicación se encuentra dentro de la carpeta server, dentro de ella está la carpeta "pages" que contiene la pagina principal (index.html) y el icono de la aplicación, por otro lado, en el mismo nivel de la carpeta antes mencionada hay dos archivos, "main.go" y "go.mod". <br/> El archivo "go.mod" muestra la raíz del modulo y la versión de go en la que se desarrolló el proyecto, pero de entre los dos el más importante es "main.go", ya que ejecutandolo se inicia el servidor de todo el proyecto. 

## Como ejecutar el servidor

### Fijarse estado de Golang
Lo primero que se debe hacer es fijarse si se tiene el lenguaje de programacion que corre el servidor instalado. El que ejecuta nuestro proyecto es [Golang](https://go.dev/). Para corrobarar la instalación o chequear si se tiene instalada una versión se debe ejecutar en la terminal (del sistema operativo que se tenga) el comando siguiente:
```
go version
```
Si esto resulto en una respuesta del tipo `go version go1.24.7 linux/amd64` o similar, se tiene instalado el lenguaje y ya puede correr el servidor.

### Ejecutar el servidor
Ahora con el lenguaje [Golang](https://go.dev/) instalado, debemos abrir la carpeta donde se encuentra este mismo archivo en la terminal de nuestro sistema operativo y ejecutar:
```
cd server
go run .
```
Luego de unos momentos, en el terminal aparecerá el mensaje `Servidor escuchando en http://localhost:8080`, para ver el proyecto solo se debe hacer click en el link generado.

### Detener el servidor
Para apagar el servidor, solo se debe clickear en la terminal nuevamente y usar la convinacion de teclas Ctrl+Z, que matará el programa ejecutando.

