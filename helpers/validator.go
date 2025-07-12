package helpers

import (
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func TranslateErrorMessage(err error, obj any) map[string]string {
	errorsMap := make(map[string]string)

	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		objType := reflect.TypeOf(obj)
		if objType.Kind() == reflect.Ptr {
			objType = objType.Elem()
		}

		for _, fieldError := range validationErrors {
			fieldName := fieldError.StructField()

			if field, found := objType.FieldByName(fieldName); found {
				jsonTag := field.Tag.Get("json")
				if jsonTag == "" || jsonTag == "-" {
					jsonTag = strings.ToLower(fieldName)
				}

				readableField := strings.ReplaceAll(jsonTag, "_", " ")

				switch fieldError.Tag() {
				case "required":
					errorsMap[jsonTag] = readableField + " is required"
				case "min":
					errorsMap[jsonTag] = readableField + " must be at least " + fieldError.Param() + " characters"
				case "max":
					errorsMap[jsonTag] = readableField + " must be at most " + fieldError.Param() + " characters"
				case "email":
					errorsMap[jsonTag] = readableField + " must be a valid email address"
				case "eqfield":
					targetField := getJSONTag(objType, fieldError.Param())
					readableTarget := strings.ReplaceAll(targetField, "_", " ")
					errorsMap[jsonTag] = readableField + " must be equal to " + readableTarget
				default:
					errorsMap[jsonTag] = readableField + " is invalid"
				}

			}
		}
	}

	if err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			if strings.Contains(err.Error(), "email") {
				errorsMap["email"] = "Email already exists"
			}
		} else if err == gorm.ErrRecordNotFound {
			errorsMap["general"] = "Record not found"
		}
	}

	return errorsMap
}

func IsDuplicateEntryError(err error) bool {
	return err != nil && strings.Contains(err.Error(), "Duplicate entry")
}

func getJSONTag(t reflect.Type, fieldName string) string {
	if field, ok := t.FieldByName(fieldName); ok {
		jsonTag := field.Tag.Get("json")
		if jsonTag != "" && jsonTag != "-" {
			return jsonTag
		}
	}
	return strings.ToLower(fieldName)
}
