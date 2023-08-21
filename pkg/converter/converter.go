package converter

import "github.com/gofrs/uuid/v5"

func StringToUUID(s string) (uuid.UUID, error) {
	return uuid.FromString(s)
}
