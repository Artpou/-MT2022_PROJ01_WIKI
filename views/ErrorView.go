package views

import (
	"fmt"
)

func FieldRequiered(field string) string {
	return fmt.Sprintf(`{"%s": ["this field is requiered"]}`, field)
}
