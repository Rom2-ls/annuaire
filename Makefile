BINARY_NAME=annuaire
MAIN_PATH=main.go

.PHONY: all build test clean run install

all: test build

build:
	go build -o ${BINARY_NAME} ${MAIN_PATH}

test:
	go test -v ./...

clean:
	go clean
	rm -f ${BINARY_NAME}

run:
	go run ${MAIN_PATH}

install:
	go mod tidy

# Exemples d'utilisation
demo:
	@echo "Démonstration de l'annuaire:"
	@echo "1. Ajout de contacts..."
	go run ${MAIN_PATH} --action ajouter --nom "Dupont" --prenom "Jean" --tel "0123456789"
	go run ${MAIN_PATH} --action ajouter --nom "Martin" --prenom "Pierre" --tel "0987654321"
	go run ${MAIN_PATH} --action ajouter --nom "Charlie" --prenom "Alice" --tel "0811223344"
	@echo "2. Liste des contacts:"
	go run ${MAIN_PATH} --action lister
	@echo "3. Recherche d'un contact:"
	go run ${MAIN_PATH} --action rechercher --nom "Charlie"


# Serveur Node.js
node-web:
	npm install
	npm start

# Tests de performance
benchmark:
	go test -bench=. ./...

# Nettoyage complet
deep-clean: clean
	rm -rf vendor/
	rm -rf node_modules/
	rm -f annuaire
	go clean -modcache
