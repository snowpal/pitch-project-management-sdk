package setupnewuser

import (
	"fmt"

	"github.com/snowpal/pitch-project-management-sdk/lib/helpers/recipes"
	"github.com/snowpal/pitch-project-management-sdk/lib/structs/response"

	log "github.com/sirupsen/logrus"
)

func RegisterNewUser() (string, error) {
	var err error
	var user response.User

	for i := 1; ; i += 1 {
		email := fmt.Sprintf("apiuser_rec%d_pm@yopmail.com", i)
		log.Info(fmt.Sprintf("Register new user with %s", email))
		user, err = recipes.RegisterUser(email)
		if err != nil {
			log.Info(fmt.Sprintf("%s is already registered.", email))
		} else {
			return user.Email, nil
		}
	}
}
