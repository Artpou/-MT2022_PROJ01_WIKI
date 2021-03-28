package views

import (
	"fmt"
)

func tmp(field string) string {
	return fmt.Sprintf(`{"%s": ["this field is requiered"]}`, field)
}
