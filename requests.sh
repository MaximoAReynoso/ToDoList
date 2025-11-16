#!/usr/bin/bash
set -e

URL="http://localhost:8080/tasks"

echo "Creando task..."
res=$(curl -s -o /dev/null -w "%{http_code}" -X POST "$URL" \
  -d "title=Ejemplo" \
  -d "description=Ejemplo description" \
  -d "completed=false")
echo -e "Resultado: $res (Caso exitoso: 201)\n\n"

echo "Obteniendo todos los tasks..."
curl -s -X GET "$URL" -i
echo -e "\n\n"

echo "Obteniendo task con ID = 1..."
curl -s -X GET "$URL/1" -i
echo -e "\n\n"

echo "Actualizando task con ID = 1..."
res=$(curl -s -o /dev/null -w "%{http_code}" -X PUT "$URL/1" \
  -d "title=Ejemplo (modificado)" \
  -d "description=Ejemplo description actualizada" \
  -d "completed=true")
echo -e "Resultado: $res (Resultado exitoso: 200)\n\n"

echo "Eliminando task con ID = 1..."
res=$(curl -s -o /dev/null -w "%{http_code}" -X DELETE "$URL/1")
echo -e "Resultado: $res (Resultado exitoso: 204)\n\n"