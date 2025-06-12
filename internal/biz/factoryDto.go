package biz

import (
	"encoding/base64"
	"net/http"
)

// factory
type DtoAdapter struct {
	converter map[string]IDto
}

func (d *DtoAdapter) SetDtoAdapter(kind string, dto IDto) {
	d.converter[kind] = dto
}

func (d *DtoAdapter) GetDtoAdapter(kind string) IDto {
	return d.converter[kind]
}

func NewDtoApdater() *DtoAdapter {
	return &DtoAdapter{map[string]IDto{}}
}

func DecodeBase64(input string) (*[]byte, error) {
	decoded, err := base64.StdEncoding.DecodeString(input)
	if err != nil {
		return nil, GetError(http.StatusInternalServerError, err)
	}

	return &decoded, nil
}

func EncodeBase64(input []byte) string {
	return base64.StdEncoding.EncodeToString(input)
}

func decodeAvatar(avatar *string) ([]byte, error) {
	if avatar != nil && *avatar != "" {
		avatarDecoded, err := DecodeBase64(*avatar)
		if err != nil {
			return nil, GetError(http.StatusUnprocessableEntity, err)
		}
		return *avatarDecoded, nil
	}

	return nil, nil
}

func encodeAvatar(avatar []byte) string {
	if avatar != nil {
		avatarDecoded := EncodeBase64(avatar)

		return avatarDecoded
	}

	return ""
}
