package faker

import (
	"math/rand"
	"technical-test/internal/domain/entity"
	"time"

	"github.com/bxcodec/faker/v4"
)

func StudentFaker() *entity.Student {
	name := faker.FirstName()
	age := rand.Intn(18-6) + 6
	grade := rand.Intn(12-1) + 1
	school := []string{"", "international school", "universitas indonesia", "budi mulia", "SMK strada",
		"balai keselamatan", "SMAN 17", "SMK BUDI OETOMO", "STAN", "SMPN 17", "Katholik School"}
	index := rand.Intn(len(school) - 1)
	return &entity.Student{
		Name:      name,
		Age:       age,
		School:    school[index],
		Grade:     grade,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
