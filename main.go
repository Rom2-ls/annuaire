package main

import (
	"annuaire/pkg/storage"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	// Définir les flags
	var action = flag.String("action", "", "Action à effectuer: ajouter, rechercher, lister, supprimer, modifier, serveur")
	var nom = flag.String("nom", "", "Nom du contact")
	var prenom = flag.String("prenom", "", "Prénom du contact")
	var telephone = flag.String("tel", "", "Numéro de téléphone du contact")
	var ancienNom = flag.String("ancien-nom", "", "Ancien nom du contact (pour modifier)")
	var ancienPrenom = flag.String("ancien-prenom", "", "Ancien prénom du contact (pour modifier)")
	var fichier = flag.String("fichier", "annuaire.json", "Fichier de sauvegarde de l'annuaire")
	var export = flag.String("export", "", "Exporter l'annuaire vers un fichier JSON")
	var importFile = flag.String("import", "", "Importer l'annuaire depuis un fichier JSON")

	flag.Parse()

	if *action == "" {
		fmt.Println("Usage: go run main.go --action <action> [options]")
		fmt.Println("Actions disponibles:")
		fmt.Println("  ajouter    : Ajouter un contact (--nom, --prenom, --tel requis)")
		fmt.Println("  rechercher : Rechercher un contact (--nom requis)")
		fmt.Println("  lister     : Lister tous les contacts")
		fmt.Println("  supprimer  : Supprimer un contact (--nom, --prenom requis)")
		fmt.Println("  modifier   : Modifier un contact (--ancien-nom, --ancien-prenom, --nom, --prenom, --tel requis)")
		fmt.Println("  serveur    : Démarrer le serveur web (--port optionnel, défaut: 8080)")
		fmt.Println("\nOptions:")
		fmt.Println("  --fichier  : Fichier de sauvegarde (défaut: annuaire.json)")
		fmt.Println("  --export   : Exporter vers un fichier JSON")
		fmt.Println("  --import   : Importer depuis un fichier JSON")
		fmt.Println("  --port     : Port pour le serveur web (défaut: 8080)")
		fmt.Println("\nExemples:")
		fmt.Println("  go run main.go --action ajouter --nom \"Dupont\" --prenom \"Jean\" --tel \"0123456789\"")
		fmt.Println("  go run main.go --action rechercher --nom \"Dupont\"")
		fmt.Println("  go run main.go --action lister")
		os.Exit(1)
	}

	// Créer l'annuaire
	annuaire := storage.NewAnnuaire(*fichier)

	// Charger l'annuaire existant
	if err := annuaire.ChargerJSON(); err != nil {
		log.Printf("Avertissement: impossible de charger l'annuaire existant: %v", err)
	}

	// Gérer l'import
	if *importFile != "" {
		annuaireImport := storage.NewAnnuaire(*importFile)
		if err := annuaireImport.ChargerJSON(); err != nil {
			log.Fatalf("Erreur lors de l'import: %v", err)
		}
		// Copier tous les contacts
		for _, contact := range annuaireImport.Lister() {
			annuaire.Ajouter(contact.Nom, contact.Prenom, contact.Telephone)
		}
		fmt.Printf("Import réussi depuis %s\n", *importFile)
	}

	// Exécuter l'action demandée
	switch *action {
	case "ajouter":
		if *nom == "" || *prenom == "" || *telephone == "" {
			log.Fatal("Pour ajouter un contact, --nom, --prenom et --tel sont requis")
		}
		
		err := annuaire.Ajouter(*nom, *prenom, *telephone)
		if err != nil {
			log.Fatalf("Erreur lors de l'ajout: %v", err)
		}
		
		fmt.Printf("Contact ajouté: %s %s (%s)\n", *prenom, *nom, *telephone)

	case "rechercher":
		if *nom == "" {
			log.Fatal("Pour rechercher un contact, --nom est requis")
		}
		
		resultats, err := annuaire.Rechercher(*nom)
		if err != nil {
			log.Fatalf("Erreur lors de la recherche: %v", err)
		}
		
		fmt.Printf("Contacts trouvés (%d):\n", len(resultats))
		for _, contact := range resultats {
			fmt.Printf("- %s\n", contact.String())
		}

	case "lister":
		contacts := annuaire.Lister()
		if len(contacts) == 0 {
			fmt.Println("Aucun contact dans l'annuaire")
		} else {
			fmt.Printf("Contacts dans l'annuaire (%d):\n", len(contacts))
			for _, contact := range contacts {
				fmt.Printf("- %s\n", contact.String())
			}
		}

	case "supprimer":
		if *nom == "" || *prenom == "" {
			log.Fatal("Pour supprimer un contact, --nom et --prenom sont requis")
		}
		
		err := annuaire.Supprimer(*nom, *prenom)
		if err != nil {
			log.Fatalf("Erreur lors de la suppression: %v", err)
		}
		
		fmt.Printf("Contact supprimé: %s %s\n", *prenom, *nom)

	case "modifier":
		if *ancienNom == "" || *ancienPrenom == "" || *nom == "" || *prenom == "" || *telephone == "" {
			log.Fatal("Pour modifier un contact, --ancien-nom, --ancien-prenom, --nom, --prenom et --tel sont requis")
		}
		
		err := annuaire.Modifier(*ancienNom, *ancienPrenom, *nom, *prenom, *telephone)
		if err != nil {
			log.Fatalf("Erreur lors de la modification: %v", err)
		}
		
		fmt.Printf("Contact modifié: %s %s -> %s %s (%s)\n", 
			*ancienPrenom, *ancienNom, *prenom, *nom, *telephone)

	default:
		log.Fatalf("Action inconnue: %s", *action)
	}

	// Sauvegarder l'annuaire
	if err := annuaire.SauvegarderJSON(); err != nil {
		log.Fatalf("Erreur lors de la sauvegarde: %v", err)
	}

	// Gérer l'export
	if *export != "" {
		annuaireExport := storage.NewAnnuaire(*export)
		// Copier tous les contacts
		for _, contact := range annuaire.Lister() {
			annuaireExport.Ajouter(contact.Nom, contact.Prenom, contact.Telephone)
		}
		if err := annuaireExport.SauvegarderJSON(); err != nil {
			log.Fatalf("Erreur lors de l'export: %v", err)
		}
		fmt.Printf("Export réussi vers %s\n", *export)
	}
}
