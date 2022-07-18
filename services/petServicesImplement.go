package services

import (
	"context"
	"errors"
	"base-app/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
    "strconv"
)
type PetServiceImpl struct {
	petcollection *mongo.Collection
	ctx            context.Context
}

func NewPetService(petcollection *mongo.Collection, ctx context.Context) PetService{
	return &PetServiceImpl{
		petcollection: petcollection,
		ctx:            ctx,
	}
}

func (p *PetServiceImpl) CreatePet(pet *models.Pet) error {
	_, err := p.petcollection.InsertOne(p.ctx, pet)
	return err
}

func (p *PetServiceImpl) GetPet(id *string) (*models.Pet, error) {
	var pet *models.Pet
	number,_:= strconv.Atoi(*id)
	query := bson.D{bson.E{Key: "id", Value: number}}
	err := p.petcollection.FindOne(p.ctx, query).Decode(&pet)
	return pet, err
}

func (p *PetServiceImpl) GetAll() ([]*models.Pet, error) {
	var pets []*models.Pet
	cursor, err := p.petcollection.Find(p.ctx, bson.D{{}})
	if err != nil {
		return nil, err
	}
	for cursor.Next(p.ctx) {
		var pet models.Pet
		err := cursor.Decode(&pet)
		if err != nil {
			return nil, err
		}
		pets=append(pets, &pet)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	cursor.Close(p.ctx)

	if len(pets) == 0 {
		return nil, errors.New("documents not found")
	}
	return pets, nil
}

func (p *PetServiceImpl) UpdatePet(pet *models.Pet) error {
	filter := bson.D{primitive.E{Key: "id", Value: pet.ID}}
	update := bson.D{primitive.E{Key: "$set", Value: bson.D{primitive.E{Key: "name", Value:pet.Name}, primitive.E{Key: "tags", Value:pet.Tags}, primitive.E{Key: "category", Value: pet.Category},primitive.E{Key: "status", Value: pet.Status}}}}
	result, _ := p.petcollection.UpdateOne(p.ctx, filter, update)
	if result.MatchedCount != 1 {
		return errors.New("no matched document found for update")
	}
	return nil
}

func (p *PetServiceImpl)DeletePet(id *string) error {
	number,_:= strconv.Atoi(*id)
	filter := bson.D{primitive.E{Key: "id", Value:number}}
	result, _ := p.petcollection.DeleteOne(p.ctx, filter)
	if result.DeletedCount != 1 {
		return errors.New("no matched document found for delete")
	}
	return nil
}

