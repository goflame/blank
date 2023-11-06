package factories

import (
	"models"
	"sha256"
	"time"
)

type UserFactory struct{}

func (u *UserFactory) Generate(model *models.BaseInterface) models.BaseInterface {
	model.Username = "test"
	h := sha256.New()
	h.Write([]byte("password"))
	model.Password = h.Sum(nil)
	model.RegisteredAt = time.Now()
	// ...your logic here
	return model
}
