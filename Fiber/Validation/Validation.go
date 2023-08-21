package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type (
	User struct {
		Name string `validate:"required,min=5,max=20"` //required field, min 5 char long max 20
		Age  int    `validate:"required,tenner"`       //Requrired field, and client needs to implement our 'tenner tag format which we'll see later'
	}

	ErrorRespones struct {
		Error       bool
		FailedField string
		Tag         string
		Value       interface {
		}
	}

	XValidator struct {
		validator *validator.Validate
	}

	GlobalErrorHandlerResp struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
	}
)

// This is the validator instance
// For more information see: https://github.com/go-playground/validator
var validate = validator.New()

func (v XValidator) Validate(data interface{}) []ErrorRespones {
	validationErrors := []ErrorRespones{}
	errs := validate.Struct(data)
	if errs != nil {
		for _, err := range errs.(validator.ValidationErrors) {
			//In this case data object is actually holding the user struct
			var elem ErrorRespones
			elem.FailedField = err.Field()
			elem.Tag = err.Tag()
			elem.Value = err.Value()
			elem.Error = true

			validationErrors = append(validationErrors, elem)
		}
	}
	return validationErrors
}

func main() {
	myValidator := &XValidator{
		validator: validate,
	}
	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(fiber.StatusBadRequest).JSON(GlobalErrorHandlerResp{
				Success: false,
				Message: err.Error(),
			})
		},
	})

	//Custom struct validation tag format
	myValidator.validator.RegisterValidation("tenner", func(fl validator.FieldLevel) bool {
		//User,Age needs to fit our needs, 12-18 years old
		return fl.Field().Int() >= 12 && fl.Field().Int() <= 18
	})

	app.Get("/", func(c *fiber.Ctx) error {
		user := &User{
			Name: c.Query("name"),
			Age:  c.QueryInt("age"),
		}

		//Validation
		if errs := myValidator.Validate(user); len(errs) > 0 && errs[0].Error {
			errMsgs := make([]string, 0)
			for _, err := range errs {
				errMsgs = append(errMsgs, fmt.Sprintf(
					"[%s]: '%v' | Needs to implement '%s'",
					err.FailedField,
					err.Value,
					err.Tag,
				))
			}
			return &fiber.Error{
				Code:    fiber.ErrBadRequest.Code,
				Message: strings.Join(errMsgs, " and "),
			}
		}
		//Logiic, validated with success
		return c.SendString("Hello, World!")
	})
	log.Fatal(app.Listen(":3000"))
}
