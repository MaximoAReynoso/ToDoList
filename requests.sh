#!/usr/bin/env bash
set -e

API_URL="http://localhost:8080/tasks"

echo "-----------------------------"
echo "Creating a task..."
response=$(curl -s -o /dev/null -w "%{http_code}" -X POST "$API_URL" \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Ejemplo",
    "description": "Ejemplo description",
    "completed": false
  }')
echo "Status: $response (Expected: 201)"
echo "-----------------------------"

echo "Getting all tasks..."
curl -s -X GET "$API_URL" -i
echo "-----------------------------"

echo "Getting task with id 1..."
curl -s -X GET "$API_URL/1" -i
echo "-----------------------------"

echo "Updating task with id 1..."
response=$(curl -s -o /dev/null -w "%{http_code}" -X PUT "$API_URL/1" \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Ejemplo (modificado)",
    "description": "Ejemplo description actualizada",
    "completed": true
  }')
echo "Status: $response (Expected: 200)"
echo "-----------------------------"

echo "Deleting task with id 1..."
response=$(curl -s -o /dev/null -w "%{http_code}" -X DELETE "$API_URL/1")
echo "Status: $response (Expected: 204)"
echo "-----------------------------"

echo "Done!"