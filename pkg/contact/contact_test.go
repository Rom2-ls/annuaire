package contact

import (
	"testing"
)

func TestNewContact(t *testing.T) {
	tests := []struct {
		nom, prenom, telephone string
		expected               *Contact
	}{
		{"Dupont", "Jean", "0123456789", &Contact{"Dupont", "Jean", "0123456789"}},
		{" Martin ", " Pierre ", " 0987654321 ", &Contact{"Martin", "Pierre", "0987654321"}},
	}

	for _, test := range tests {
		result := NewContact(test.nom, test.prenom, test.telephone)
		if result.Nom != test.expected.Nom || result.Prenom != test.expected.Prenom || result.Telephone != test.expected.Telephone {
			t.Errorf("NewContact(%s, %s, %s) = %v, expected %v", 
				test.nom, test.prenom, test.telephone, result, test.expected)
		}
	}
}

func TestNomComplet(t *testing.T) {
	contact := NewContact("Dupont", "Jean", "0123456789")
	expected := "Jean Dupont"
	result := contact.NomComplet()
	
	if result != expected {
		t.Errorf("NomComplet() = %s, expected %s", result, expected)
	}
}

func TestEstValide(t *testing.T) {
	tests := []struct {
		contact  *Contact
		expected bool
	}{
		{NewContact("Dupont", "Jean", "0123456789"), true},
		{NewContact("", "Jean", "0123456789"), false},
		{NewContact("Dupont", "", "0123456789"), false},
		{NewContact("Dupont", "Jean", ""), false},
		{NewContact("", "", ""), false},
	}

	for _, test := range tests {
		result := test.contact.EstValide()
		if result != test.expected {
			t.Errorf("EstValide() pour %v = %t, expected %t", test.contact, result, test.expected)
		}
	}
}

func TestToJSONAndFromJSON(t *testing.T) {
	originalContact := NewContact("Dupont", "Jean", "0123456789")
	
	// Test ToJSON
	jsonData, err := originalContact.ToJSON()
	if err != nil {
		t.Errorf("ToJSON() error = %v", err)
		return
	}
	
	// Test FromJSON
	restoredContact, err := FromJSON(jsonData)
	if err != nil {
		t.Errorf("FromJSON() error = %v", err)
		return
	}
	
	if restoredContact.Nom != originalContact.Nom || 
	   restoredContact.Prenom != originalContact.Prenom || 
	   restoredContact.Telephone != originalContact.Telephone {
		t.Errorf("JSON round trip failed. Original: %v, Restored: %v", originalContact, restoredContact)
	}
}
