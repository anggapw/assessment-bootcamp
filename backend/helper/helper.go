package helper

import (
	"errors"
	"math"
	"strconv"

	"github.com/go-playground/validator/v10"
)

type Response struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

type Meta struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
}

func SplitErrorInformation(err error) []string {
	var errors []string

	for _, e := range err.(validator.ValidationErrors) {
		errors = append(errors, e.Error())
	}

	return errors
}

func APIResponse(message string, code int, status string, data interface{}) Response {
	var meta = Meta{
		Message: message,
		Code:    code,
		Status:  status,
	}

	var response = Response{
		Meta: meta,
		Data: data,
	}

	return response
}

func ValidateID(ID string) error {
	if num, err := strconv.Atoi(ID); err != nil || num == 0 || math.Signbit(float64(num)) == true {
		return errors.New("input id user harus benar")
	}

	return nil
}
