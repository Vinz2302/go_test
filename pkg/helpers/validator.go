package helpers

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

type IValidationHelper interface {
	UserExists(fl validator.FieldLevel) bool
}

type validationHelper struct {
	IValidationHelper
}

func initValidationHelper(IValidationHelper IValidationHelper) *validationHelper {
	return &validationHelper{
		IValidationHelper,
	}
}

/* func HelperInit(db *gorm.DB, conf *config.Conf) *validationHelper {
	var (
		helperDatabase       = helperDatabases.InitHelperDatabase(db)
		userRepository       = repository.InitUserRepository(db, helperDatabase, conf)
		roleRepository       = roleRepository.InitRoleRepository(db, helperDatabase, conf)
		userService          = userService.InitUserRepository(userRepository, roleRepository)
		helperInjectionObj   = InitHelperInjection(userService)
		initValidationHelper = InitValidationHelper(helperInjectionObj)
	)

	return initValidationHelper
} */

func ErrorMessage(err interface{}) []string {
	errorMessages := []string{}
	for _, e := range err.(validator.ValidationErrors) {
		fmt.Println(e.ActualTag())
		switch e.ActualTag() {
		case "Enum":
			replacer := *strings.NewReplacer("_", ",")
			errorMessage := fmt.Sprintf("Error on field %s, must be one of: %s", e.Field(), replacer.Replace(e.Param()))
			errorMessages = append(errorMessages, errorMessage)
		case "EnumVersionTwo":
			replacer := *strings.NewReplacer("&", ", ")
			errorMessage := e.Field() + " must be one of " + replacer.Replace(e.Param())
			errorMessages = append(errorMessages, errorMessage)
		case "UserExists":
			errorMessage := fmt.Sprintf("Error on field %s, condition: User with ID %v is not exists", e.Field(), e.Value())
			errorMessages = append(errorMessages, errorMessage)
		case "min":
			errorMessage := fmt.Sprintf("Error on field %s, condition: Should Be At Least %v Character", e.Field(), e.Param())
			errorMessages = append(errorMessages, errorMessage)
		case "e164":
			errorMessage := fmt.Sprintf("Error on field %s, condition: Must Use Country Code Like: +62", e.Field())
			errorMessages = append(errorMessages, errorMessage)
		case "email":
			errorMessage := fmt.Sprintf("Error on field %s, condition: Must Use The Correct Email Format", e.Field())
			errorMessages = append(errorMessages, errorMessage)
		case "gte":
			errorMessage := fmt.Sprintf("Error on field %s, condition: Must Grather Than Equals %v", e.Field(), e.Param())
			errorMessages = append(errorMessages, errorMessage)
		default:
			errorMessage := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}
	}
	return errorMessages
}
