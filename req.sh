#!/usr/bin/bash
set -e

URL="http://localhost:8080/tasks"

echo "Creando task..."
res=$(curl -i -X POST http://localhost:8080/tasks \
  -H "Content-Type: application/x-www-form-urlencoded" \
  -d "title=Comprar pan&description=Ir a la tienda&completed=on")
echo -e "Resultado: $res (Caso exitoso: 200)\n\n"

echo "Obteniendo todos los tasks..."
res=$(curl -i -X GET http://localhost:8080/tasks \
  -H "Accept: text/html")
echo -e "Resultado: $res (Caso exitoso: 200)\n\n"

echo -e "El verbo de HTTP DELETE puede ser probado en la aplicacion.\n\n\n"