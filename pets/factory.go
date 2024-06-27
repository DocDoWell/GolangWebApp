package pets

import "web-app/models"

func NewPet(species string) *models.Pet {
	pet := models.Pet{
		Species: species,
		Breed: "",
		MinWeight: 0,
		MaxWeight: 0,
		Description: " nothing entered yet",
		LifeSpan: 0,
	}
	return &pet
}