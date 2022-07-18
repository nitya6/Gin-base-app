package services
import "base-app/models"
type PetService interface {
	CreatePet(*models.Pet) error
	GetPet(*string) (*models.Pet, error)
	GetAll() ([]*models.Pet, error)
	UpdatePet(*models.Pet) error
	DeletePet(*string) error
}

