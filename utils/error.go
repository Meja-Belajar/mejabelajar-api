package utils

import "github.com/meja_belajar/models/outputs"

func HandleBadRequest(message string) (int, interface{}) {
	output := outputs.BadRequestOutput{
		Code:    400,
		Message: "Bad Request: " + message,
	}
	return 400, output
}

func HandleUnauthorized(message string) (int, interface{}) {
	output := outputs.UnauthorizedOutput{
		Code:    401,
		Message: "Unauthorized: " + message,
	}
	return 401, output
}

func HandleNotFound(entity string) (int, interface{}) {
	output := outputs.NotFoundOutput{
		Code:    404,
		Message: "Not Found: " + entity + " not found",
	}
	return 404, output
}

func HandleTimeout(err error) (int, interface{}) {
	output := outputs.RequestTimeoutOutput{
		Code:    408,
		Message: "Request Timeout",
	}
	return 408, output

}

func HandleConflict(message string) (int, interface{}) {
	output := outputs.ConflictOutput{
		Code:    409,
		Message: "Conflict: " + message,
	}
	return 409, output
}

func HandleInternalServerError(err error) (int, interface{}) {
	output := outputs.InternalServerErrorOutput{
		Code:    500,
		Message: "Internal Server Error: " + err.Error(),
	}
	return 500, output
}
