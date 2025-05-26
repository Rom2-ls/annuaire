package storage

import (
	"annuaire/pkg/contact"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// Annuaire représente l'annuaire de contacts
type Annuaire struct {
	Contacts map[string]*contact.Contact
	fichier  string
}

// NewAnnuaire crée un nouvel annuaire
func NewAnnuaire(fichier string) *Annuaire {
	return &Annuaire{
		Contacts: make(map[string]*contact.Contact),
		fichier:  fichier,
	}
}

// genererCle génère une clé unique pour un contact basée sur le nom complet
func (a *Annuaire) genererCle(nom, prenom string) string {
	return strings.ToLower(fmt.Sprintf("%s_%s", strings.TrimSpace(prenom), strings.TrimSpace(nom)))
}

// Ajouter ajoute un contact à l'annuaire
func (a *Annuaire) Ajouter(nom, prenom, telephone string) error {
	cle := a.genererCle(nom, prenom)
	
	// Vérifier si le contact existe déjà
	if _, existe := a.Contacts[cle]; existe {
		return fmt.Errorf("un contact avec le nom '%s %s' existe déjà", prenom, nom)
	}
	
	nouveauContact := contact.NewContact(nom, prenom, telephone)
	if !nouveauContact.EstValide() {
		return fmt.Errorf("contact invalide: tous les champs sont requis")
	}
	
	a.Contacts[cle] = nouveauContact
	return nil
}

// Rechercher cherche un contact par nom
func (a *Annuaire) Rechercher(nom string) ([]*contact.Contact, error) {
	var resultats []*contact.Contact
	nomLower := strings.ToLower(nom)
	
	for cle, contact := range a.Contacts {
		if strings.Contains(cle, nomLower) || 
		   strings.Contains(strings.ToLower(contact.Nom), nomLower) ||
		   strings.Contains(strings.ToLower(contact.Prenom), nomLower) {
			resultats = append(resultats, contact)
		}
	}
	
	if len(resultats) == 0 {
		return nil, fmt.Errorf("aucun contact trouvé pour '%s'", nom)
	}
	
	return resultats, nil
}

// Lister retourne tous les contacts
func (a *Annuaire) Lister() []*contact.Contact {
	var contacts []*contact.Contact
	for _, contact := range a.Contacts {
		contacts = append(contacts, contact)
	}
	return contacts
}

// Supprimer supprime un contact de l'annuaire
func (a *Annuaire) Supprimer(nom, prenom string) error {
	cle := a.genererCle(nom, prenom)
	
	if _, existe := a.Contacts[cle]; !existe {
		return fmt.Errorf("contact '%s %s' non trouvé", prenom, nom)
	}
	
	delete(a.Contacts, cle)
	return nil
}

// Modifier modifie un contact existant
func (a *Annuaire) Modifier(ancienNom, ancienPrenom, nouveauNom, nouveauPrenom, nouveauTelephone string) error {
	ancienneCle := a.genererCle(ancienNom, ancienPrenom)
	
	// Vérifier si l'ancien contact existe
	if _, existe := a.Contacts[ancienneCle]; !existe {
		return fmt.Errorf("contact '%s %s' non trouvé", ancienPrenom, ancienNom)
	}
	
	// Supprimer l'ancien contact
	delete(a.Contacts, ancienneCle)
	
	// Ajouter le nouveau contact
	return a.Ajouter(nouveauNom, nouveauPrenom, nouveauTelephone)
}

// SauvegarderJSON sauvegarde l'annuaire au format JSON
func (a *Annuaire) SauvegarderJSON() error {
	data, err := json.MarshalIndent(a.Contacts, "", "  ")
	if err != nil {
		return fmt.Errorf("erreur lors de la sérialisation JSON: %v", err)
	}
	
	return ioutil.WriteFile(a.fichier, data, 0644)
}

// ChargerJSON charge l'annuaire depuis un fichier JSON
func (a *Annuaire) ChargerJSON() error {
	if _, err := os.Stat(a.fichier); os.IsNotExist(err) {
		// Le fichier n'existe pas, c'est normal pour un nouvel annuaire
		return nil
	}
	
	data, err := ioutil.ReadFile(a.fichier)
	if err != nil {
		return fmt.Errorf("erreur lors de la lecture du fichier: %v", err)
	}
	
	return json.Unmarshal(data, &a.Contacts)
}

// NombreContacts retourne le nombre de contacts dans l'annuaire
func (a *Annuaire) NombreContacts() int {
	return len(a.Contacts)
}
