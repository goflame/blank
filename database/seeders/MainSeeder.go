package seeders

import (
	"factories"
	"models"
)

func MainSeeder() {
	factories.SetSeed("YOUR RANDOM SEED")

	models.User.Seeder(func(user *models.User) *models.User {
		factories.Factory.Generate(user)
		return user
	}, 15)
}
