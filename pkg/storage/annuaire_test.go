package storage

import (
	"os"
	"testing"
)

func TestAjouterContact(t *testing.T) {
	annuaire := NewAnnuaire("test_annuaire.json")
	
	// Test d'ajout normal
	err := annuaire.Ajouter("Dupont", "Jean", "0123456789")
	if err != nil {
		t.Errorf("Erreur lors de l'ajout du contact: %v", err)
	}
	
	if annuaire.NombreContacts() != 1 {
		t.Errorf("Nombre de contacts = %d, expected 1", annuaire.NombreContacts())
	}
	
	// Test d'ajout d'un contact en double
	err = annuaire.Ajouter("Dupont", "Jean", "0987654321")
	if err == nil {
		t.Error("L'ajout d'un contact en double devrait générer une erreur")
	}
}

func TestRechercherContact(t *testing.T) {
	annuaire := NewAnnuaire("test_annuaire.json")
	
	// Ajouter des contacts de test
	annuaire.Ajouter("Dupont", "Jean", "0123456789")
	annuaire.Ajouter("Martin", "Pierre", "0987654321")
	annuaire.Ajouter("Durand", "Marie", "0555123456")
	
	// Test de recherche existante
	resultats, err := annuaire.Rechercher("Dupont")
	if err != nil {
		t.Errorf("Erreur lors de la recherche: %v", err)
	}
	if len(resultats) != 1 {
		t.Errorf("Nombre de résultats = %d, expected 1", len(resultats))
	}
	if resultats[0].Nom != "Dupont" {
		t.Errorf("Contact trouvé = %s, expected Dupont", resultats[0].Nom)
	}
	
	// Test de recherche inexistante
	_, err = annuaire.Rechercher("Inexistant")
	if err == nil {
		t.Error("La recherche d'un contact inexistant devrait générer une erreur")
	}
}

func TestSupprimerContact(t *testing.T) {
	annuaire := NewAnnuaire("test_annuaire.json")
	
	// Ajouter un contact
	annuaire.Ajouter("Dupont", "Jean", "0123456789")
	
	// Supprimer le contact
	err := annuaire.Supprimer("Dupont", "Jean")
	if err != nil {
		t.Errorf("Erreur lors de la suppression: %v", err)
	}
	
	if annuaire.NombreContacts() != 0 {
		t.Errorf("Nombre de contacts après suppression = %d, expected 0", annuaire.NombreContacts())
	}
	
	// Test de suppression d'un contact inexistant
	err = annuaire.Supprimer("Inexistant", "Personne")
	if err == nil {
		t.Error("La suppression d'un contact inexistant devrait générer une erreur")
	}
}

func TestSauvegarderEtChargerJSON(t *testing.T) {
	fichierTest := "test_save_load.json"
	defer os.Remove(fichierTest) // Nettoyer après le test
	
	// Créer un annuaire et ajouter des contacts
	annuaire1 := NewAnnuaire(fichierTest)
	annuaire1.Ajouter("Dupont", "Jean", "0123456789")
	annuaire1.Ajouter("Martin", "Pierre", "0987654321")
	
	// Sauvegarder
	err := annuaire1.SauvegarderJSON()
	if err != nil {
		t.Errorf("Erreur lors de la sauvegarde: %v", err)
	}
	
	// Charger dans un nouvel annuaire
	annuaire2 := NewAnnuaire(fichierTest)
	err = annuaire2.ChargerJSON()
	if err != nil {
		t.Errorf("Erreur lors du chargement: %v", err)
	}
	
	// Vérifier que les données sont identiques
	if annuaire2.NombreContacts() != annuaire1.NombreContacts() {
		t.Errorf("Nombre de contacts chargés = %d, expected %d", 
			annuaire2.NombreContacts(), annuaire1.NombreContacts())
	}
}
