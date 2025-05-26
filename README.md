# Annuaire en Go

## 📋 Fonctionnalités

- ✅ **Ajouter** un contact (nom, prénom, téléphone)
- ✅ **Rechercher** des contacts par nom/prénom
- ✅ **Lister** tous les contacts
- ✅ **Supprimer** un contact
- ✅ **Modifier** un contact existant
- ✅ **Import/Export** JSON
- ✅ **Vérification** des doublons
- ✅ **Tests unitaires** complets
- ✅ **Sauvegarde automatique** en JSON
- 🎉 **Interface web** moderne avec Go
- 🎉 **Auto-refresh** en temps réel
- 🎉 **Serveur Node.js** avec WebSocket
- 🎉 **Mode développement** avec compilation automatique

## 🚀 Installation et Configuration

### Prérequis

- Go 1.19 ou plus récent
- Git (optionnel)
- Node.js 16+ (pour l'interface web bonus)
- npm (pour les dépendances Node.js)

### Installation

```bash
# Cloner le projet
git clone <votre-repo-url>
cd annuaire

# Installer les dépendances
make install

# Compiler le projet
make build
```

## 📖 Utilisation

### Commandes de base

#### Ajouter un contact

```bash
go run main.go --action ajouter --nom "Dupont" --prenom "Jean" --tel "0123456789"
```

#### Rechercher un contact

```bash
go run main.go --action rechercher --nom "Dupont"
```

#### Lister tous les contacts

```bash
go run main.go --action lister
```

#### Supprimer un contact

```bash
go run main.go --action supprimer --nom "Dupont" --prenom "Jean"
```

#### Modifier un contact

```bash
go run main.go --action modifier --ancien-nom "Dupont" --ancien-prenom "Jean" --nom "Martin" --prenom "Pierre" --tel "0987654321"
```

### Fonctionnalités avancées

#### Export vers JSON

```bash
go run main.go --action lister --export "backup.json"
```

#### Import depuis JSON

```bash
go run main.go --action lister --import "backup.json"
```

#### Utiliser un fichier de sauvegarde personnalisé

```bash
go run main.go --action lister --fichier "mon_annuaire.json"
```

### Exemples d'utilisation

```bash
go run main.go --action ajouter --nom "Charlie" --prenom "Alice" --tel "0811223344"
go run main.go --action rechercher --nom "Alice"
go run main.go --action lister
go run main.go --action rechercher --nom "Charlie"
```

## 🌐 Interface Web

```bash
make node-web
```

Interface disponible sur : http://localhost:3000

## 🧪 Tests

### Exécuter tous les tests

```bash
make test
```

### Tests unitaires détaillés

```bash
go test -v ./pkg/contact/
go test -v ./pkg/storage/
```

### Démonstration rapide

```bash
make demo
```

## 🏗️ Structure du projet

```
annuaire/
├── main.go                      # Point d'entrée principal
├── pkg/
│   ├── contact/
│   │   ├── contact.go           # Structure Contact
│   │   └── contact_test.go      # Tests du package contact
│   └── storage/
│       ├── annuaire.go          # Logique de l'annuaire
│       └── annuaire_test.go     # Tests du package storage
├── Makefile                     # Automatisation des tâches
├── README.md                    # Documentation
├── go.mod                       # Module Go
├── package.json                 # Dépendances Node.js pour l'interface web
├── web-server.js                # Serveur Node.js avec auto-refresh
└── annuaire.json               # Fichier de sauvegarde (généré)
```

## 🐛 Résolution des problèmes

### Le fichier JSON est corrompu

```bash
# Supprimer le fichier et redémarrer
rm annuaire.json
go run main.go --action lister
```

### Erreur de compilation

```bash
# Nettoyer et recompiler
make clean
make install
make build
```
