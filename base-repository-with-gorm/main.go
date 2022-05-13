package main

import (
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Profile struct {
	ID         string
	FirstName  string
	MiddleName string
	LastName   string
}

func (p *Profile) BeforeCreate(_ *gorm.DB) (err error) {
	p.ID = uuid.New().String()
	return nil
}

func (p Profile) TableName() string {
	return "profiles"
}

// ProfilesRepository shows that you can wrap a base repository to expand functionality.
type ProfilesRepository struct {
	*BaseRepo[Profile, string]
}

func NewProfilesRepository() ProfilesRepository {
	return ProfilesRepository{
		BaseRepo: NewRepo[Profile, string](),
	}
}

// main contains an example of how you can write one base repo and have the base-level functionality for all CRUD actions.
func main() {
	profileRepo := NewProfilesRepository()

	createResult, err := profileRepo.Create(Profile{
		FirstName:  "Mary",
		MiddleName: "O.",
		LastName:   "Smith",
	})
	if err != nil {
		panic(err)
	}

	tx, _ := profileRepo.BeginTx()
	if _, err := tx.GetByID(createResult.ID); err != nil {
		panic(err)
	}
	if _, err = tx.Update(createResult.ID, Profile{
		FirstName: "Test",
	}); err != nil {
		panic(err)
	}
	_ = tx.CommitTx()

	// This works too instead, would change the final result first name back to "Mary"
	//_ = tx.RollbackTx()

	getResult, err := profileRepo.GetByID(createResult.ID)
	if err != nil {
		panic(err)
	}

	s, _ := json.MarshalIndent(getResult, "", "\t")
	fmt.Printf("Final result: %+v\n", string(s))
}
