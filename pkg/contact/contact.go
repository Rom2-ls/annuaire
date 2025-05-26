package contact

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Contact représente un contact dans l'annuaire
type Contact struct {
	Nom      string `json:"nom"`
	Prenom   string `json:"prenom"`
	Telephone string `json:"telephone"`
}

// NewContact crée un nouveau contact
func NewContact(nom, prenom, telephone string) *Contact {
	return &Contact{
		Nom:       strings.TrimSpace(nom),
		Prenom:    strings.TrimSpace(prenom),
		Telephone: strings.TrimSpace(telephone),
	}
}

// NomComplet retourne le nom complet du contact
func (c *Contact) NomComplet() string {
	return fmt.Sprintf("%s %s", c.Prenom, c.Nom)
}

// String retourne une représentation string du contact
func (c *Contact) String() string {
	return fmt.Sprintf("Nom: %s, Prénom: %s, Téléphone: %s", c.Nom, c.Prenom, c.Telephone)
}

// ToJSON convertit le contact en JSON
func (c *Contact) ToJSON() ([]byte, error) {
	return json.Marshal(c)
}

// FromJSON crée un contact depuis JSON
func FromJSON(data []byte) (*Contact, error) {
	var contact Contact
	err := json.Unmarshal(data, &contact)
	if err != nil {
		return nil, err
	}
	return &contact, nil
}

// EstValide vérifie si le contact est valide
func (c *Contact) EstValide() bool {
	return c.Nom != "" && c.Prenom != "" && c.Telephone != ""
}
