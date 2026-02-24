package text

import (
	"strings"

	"github.com/mt1976/frantic-core/idHelpers"
)

func BuildSignature(in string) string {
	id := idHelpers.Encode(strings.ToUpper(idHelpers.SanitizeID(in)))
	return id
}
