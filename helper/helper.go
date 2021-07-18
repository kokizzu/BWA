package helper

import (
	"BWA/rpcp"

	"github.com/go-playground/validator/v10"
)

type Response struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

func (r Response) ToMetaProto(meta *rpcp.Meta) {
	meta.Message = r.Meta.Message
	meta.Code = int32(r.Meta.Code)
	meta.Status = r.Meta.Status
}

type Meta struct {
	Message string
	Code    int
	Status  string
}

func APIResponse(message string, code int, status string, data interface{}) Response {
	meta := Meta{
		Message: message,
		Code:    code,
		Status:  status,
	}

	NewRensponse := Response{
		Meta: meta,
		Data: data,
	}

	return NewRensponse
}

func FormatError(err error) []string {
	var errors []string

	for _, e := range err.(validator.ValidationErrors) {
		errors = append(errors, e.Error())
	}

	return errors
}
