# ToDoList

Esta aplicación, como se puede ver, es una lista de tareas.

## Interfaz

La interfaz de la aplicación es la siguiente: Está su título como se va arriba, y el espacio donde se van a listar los elementos de una tabla con tareas pendientes o 
completadas. Con un recuadro fijo en la esquina de la pantalla, siendo la plantilla a completar para generar la nueva tarea.<br/>
Cada tarea generada tiene sus elementos visibles, estos son su título, descripción, y si fue completada o no, siendo por defecto la tarea marcada como no completada. Tienen un elemento no visible pero presente en la base de datos que es su identificador o ID, generado automaticamente de manera incremental.<br/>
Se agrego al diseño un icono acorde a la aplicación.

## Estructura de la pagina

Toda la aplicación se encuentra dentro de la carpeta server, dentro de ella está la carpeta "static" que contiene los elementos estaticos, por otro lado, en el mismo nivel de la carpeta antes mencionada hay dos archivos, "main.go" y "go.mod". <br/> El archivo "go.mod" muestra la raíz del modulo y la versión de go en la que se desarrolló el proyecto, pero de entre los dos el más importante es "main.go", ya que ejecutándolo se inicia el servidor de todo el proyecto. <br/><br/>

Junto a "static" hay otra carpeta llamada "db" con sus subcarpetas "queries" (con un archivo con las operaciones permitidas sobre la base de datos), "schema" (con su archivo para la construcción de la tabla) y "sqlc" con el código Go seguro y tipado generado por el comando homónimo al iniciar el programa.<br/>
También, junto a la carpeta "db" se agregó un archivo "sqlc.yaml" para especificar las variables con las que se va a ejecutar dicho comando.<br/><br/>

La carpeta logic adyacente a las anteriores contiene los conectores y operaciones necesarias para la interacción entre la base de datos y el servidor. Todo lo anterior engloba el contenido de la carpeta server, junto con el `Dockerfile` utilizado en el montado de la aplicación.<br/><br/>

Por ultimo, views, con archivos templ que al compilarse generan ahi mismo los archivos golang. Dichos archivos proveen la forma de renderizar la aplicacion del lado del servidor.<br/><br/>

Por fuera de /server, esta el archivo de testeo `requests.hurl`, `docker-compose.yml` encargado de maquetar como se debe conteinearizar el programa entero y el archivo con los comandos que levantan la aplicación `makefile`.

## Como ejecutar la aplicación

### Fijarse estado del lenguaje de programación del servidor
Lo primero que se debe hacer es fijarse si se tiene el lenguaje de programación que corre el servidor instalado. El que ejecuta nuestro proyecto es [Golang](https://go.dev/). Para corrobarar la instalación o chequear si se tiene instalada una versión se debe ejecutar en la terminal (del sistema operativo que se tenga) el comando siguiente:
```
go version
```
Si esto resulto en una respuesta del tipo `go version go1.24.7 linux/amd64` o similar, se tiene instalado el lenguaje y ya puede correr el servidor.

### Corroborar tener Docker instalado 
Se debe tener instalado Docker para ejecutar en sus respectivos contenedores de manera aislada el servidor y la base de datos. Puede fijarse si tiene instalado el programa utilizando el comando:
```
docker version
```
En caso de no tener una respecta satisfactoria, se puede instalar la herramienta desde su pagina [oficial](https://docs.docker.com/engine/install/) siguiendo las instrucciones correspondientes a su sistema operativo.

### Ejecutar la aplicación
Ahora con el lenguaje y herramienta instalados ya podemos ejecutar la aplicación. Se provee un archivo llamado makefile, todo lo que se debe hacer para montar y empezar a usar es usar el siguiente comando:
```
make
```
Con la iniciación correcta, se muestra todo el proceso de montado y testeo. Para ver el proyecto en el navegador solo se debe clickear en el link generado al final del proyecto.

### Detener la aplicación
Para apagar el servidor, solo se debe clickear en la terminal nuevamente y usar siguiente comando, que dentendrá y eliminará lo generado en docker:
```
make off
```