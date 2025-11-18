#all: sqlc templ build imagen levantar test fin
all: sqlc templ build imagen levantar fin

sqlc:
	@printf "Aplicando sqlc...\n"
	@cd server && sqlc generate
	@printf "Sqlc generado.\n\n\n"
	@sleep 2

templ:
	@printf "Compilando archivos templ...\n"
	@cd server/views && templ generate
	@printf "Archivos templ compilados.\n\n\n"
	@sleep 2

build:
	@printf "Construyendo app...\n"
	@cd server && go build -o exe .
	@printf "App construida.\n\n\n"
	@sleep 2

imagen:
	@printf "Compilando imagen Docker...\n"
	@docker compose build --no-cache
	@printf "Imagen compilada.\n\n\n"
	@sleep 2

levantar:
	@printf "Levantando Docker...\n"
	@docker compose up -d
	@printf "Docker levantado.\n"
	@printf "Espere 20 segundos a que la base de datos termine de iniciar...\n\n\n"
	@sleep 20

#test:
#	@echo "Realizando pruebas...\n"
#	@hurl --test requests.hurl
#	@echo "Pruebas realizadas exitosamente.\n\n\n"

fin:
	@printf "Recursos cargados exitosamente, puede seguir usando la aplicacion.\n"
	@printf "Aplicacion corriendo en: http://localhost:8080.\n"

off:
	@docker compose down