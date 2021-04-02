package views

import (
	"fmt"
)

func FieldNotFound(field string) string {
	return fmt.Sprintf(`This %s cannot be found`, field)
}

func FieldRequiered(field string) string {
	return fmt.Sprintf(`The field '%s' is missing`, field)
}

func FieldForbidden(field string) string {
	return fmt.Sprintf(`Youd don't have the write to modify this %s`, field)
}
