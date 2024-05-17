package outputs

type BaseOutput struct {
	Code    int    `json:"code" default:"200"`
	Message string `json:"message" default:"Success: {message}"`
}

type BadRequestOutput struct {
	Code    int    `json:"code" default:"400"`
	Message string `json:"message" default:"Bad Request: {message}"`
}

type UnauthorizedOutput struct {
	Code    int    `json:"code" default:"401"`
	Message string `json:"message" default:"Unauthorized: {message}"`
}

type ForbiddenOutput struct {
	Code    int    `json:"code" default:"403"`
	Message string `json:"message" default:"Forbidden: {message}"`
}

type NotFoundOutput struct {
	Code    int    `json:"code" default:"404"`
	Message string `json:"message" default:"Not Found: {message}"`
}

type ConflictOutput struct {
	Code    int    `json:"code" default:"409"`
	Message string `json:"message" default:"Conflict: {message}"`
}

type RequestTimeoutOutput struct {
	Code    int    `json:"code" default:"408"`
	Message string `json:"message" default:"Request Timeout: {message}"`
}

type InternalServerErrorOutput struct {
	Code    int    `json:"code" default:"500"`
	Message string `json:"message" default:"Internal Server Error: {message}"`
}