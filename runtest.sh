#!/bin/bash

set -e

echo -e "Aplicando sqlc..."
cd server
sqlc generate
cd ..
echo -e "Sqlc generado.\n\n\n"

sleep 5

echo -e "Construyendo app..."
cd server
go build -o exe .
cd ..
echo -e "App construida.\n\n\n"

sleep 5

echo -e "Compilando imagen Docker..."
docker compose build --no-cache
echo -e "Imagen compilada.\n\n\n"

sleep 5

echo -e "Levantando Docker..."
docker compose up -d
echo -e "Docker levantado."
echo -e "Espere 20 segundos a que la base de datos termine de iniciar...\n\n\n"

sleep 20

echo -e "Ejecutando tests..."
./requests.sh
echo -e "Test realizados.\n\n\n"

sleep 5

echo "Recursos cargados exitosamente, puede seguir usando la aplicacion."
echo -e "Aplicacion corriendo en: http://localhost:8080. \n"