# requests.hurl

# 1. Crear un task
POST http://localhost:8080/tasks
Content-Type: application/json
{
    "title": "Ejemplo",
    "description": "Ejemeplo description",
    "completed": false
}
HTTP 201

# 2. Obtener todos los tasks
GET http://localhost:8080/tasks
HTTP 200

# 3. Obtener el task con id 1
GET http://localhost:8080/tasks/1
HTTP 200

# 4. Modificar el task con id 1
PUT http://localhost:8080/tasks/1
Content-Type: application/json
{
    "title": "Ejemplo",
    "description": "Ejemeplo description",
    "completed": false
}
HTTP 200


# 5. Eliminar el task con id 1
DELETE http://localhost:8080/tasks/1
HTTP 204