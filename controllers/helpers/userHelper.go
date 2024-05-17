package helpers

import (
	"context"
	"time"

	"github.com/meja_belajar/models/database"
	"github.com/meja_belajar/models/outputs"
	"github.com/meja_belajar/models/requests"
	"github.com/meja_belajar/models/responses"
	"github.com/meja_belajar/repositories"
	"github.com/meja_belajar/utils"
	"gorm.io/gorm"
)

func RegisterUser(AddUserRequestDTO requests.RegisterUserRequestDTO) (int, interface{}) {
	//validasi password dan confirm password
	if AddUserRequestDTO.Password != AddUserRequestDTO.ConfirmPassword {
		output := outputs.BadRequestOutput{
			Code:    400,
			Message: "Bad Request: Password and Confirm Password does not match",
		}
		return 400, output
	}

	//validasi hash password
	hashedPassword, err := utils.HashPassword(AddUserRequestDTO.Password)
	//validasi hash password
	if err != nil {
		output := outputs.InternalServerErrorOutput{
			Code:    500,
			Message: "Internal Server Error: " + err.Error(),
		}
		return 500, output
	}
	// timeout, err := time.ParseDuration(os.Getenv("TIMEOUT"))
	// if err != nil {
	// 	output := outputs.InternalServerErrorOutput{
	// 		Code:    500,
	// 		Message: "Internal Server Error: " + err.Error(),
	// 	}
	// 	return 500, output
	// }
	// ctx, cancel := context.WithTimeout(context.Background(), timeout)
	// defer cancel()
	// db := configs.GetDB().WithContext(ctx)
	// var oldUser database.Users
	// err = db.
	// 	Where("email = ?", AddUserRequestDTO.Email).
	// 	First(&oldUser).
	// 	Error

	var user database.Users
	user, err = repositories.FindUserByEmail(AddUserRequestDTO.Email)

	//validasi timeout
	if err == context.DeadlineExceeded {
		output := outputs.RequestTimeoutOutput{
			Code:    408,
			Message: "Request Timeout",
		}
		return 408, output
	}

	//validasi email belum terdaftar
	if err == nil {
		output := outputs.ConflictOutput{
			Code:    409,
			Message: "Conflict: email already used",
		}
		return 409, output
	}

	// // buat user baru
	bod, err := time.Parse("2006-01-02T15:04:05Z", AddUserRequestDTO.BOD)
	if err != nil {
		output := outputs.BadRequestOutput{
			Code:    400,
			Message: "Bad Request: Invalid date format",
		}
		return 400, output
	}

	user = database.Users{
		Username:       AddUserRequestDTO.UserName,
		University:     AddUserRequestDTO.University,
		Email:          AddUserRequestDTO.Email,
		Phone:          AddUserRequestDTO.PhoneNumber,
		BOD:            bod,
		Password:       hashedPassword,
		IsActive:       true,
		ProfilePicture: AddUserRequestDTO.ProfilePicture,
	}
	err = repositories.InsertUser(user)

	//validasi timeout
	if err == context.DeadlineExceeded {
		output := outputs.RequestTimeoutOutput{
			Code:    408,
			Message: "Request Timeout",
		}
		return 408, output
	}

	//validasi create user
	if err != nil {
		output := outputs.InternalServerErrorOutput{
			Code:    500,
			Message: "Internal Server Error: " + err.Error(),
		}
		return 500, output
	}

	output := outputs.RegisterUserOutput{
		BaseOutput: outputs.BaseOutput{
			Code:    201,
			Message: "Success: Account has been created",
		},
		Data: responses.UserResponseDTO{
			ID:             user.ID,
			UserName:       user.Username,
			University:     user.University,
			Email:          user.Email,
			PhoneNumber:    user.Phone,
			Description:    user.Description,
			ProfilePicture: user.ProfilePicture,
			BOD:            user.BOD,
			IsActive:       user.IsActive,
			IsMentor:       false,
		},
	}
	return 201, output
}

func LoginUser(LoginUserRequestDTO requests.LoginUserRequestDTO) (int, interface{}, string) {
	var user database.Users
	user, err := repositories.FindUserByEmail(LoginUserRequestDTO.Email)

	//validasi timeout
	if err == context.DeadlineExceeded {
		output := outputs.RequestTimeoutOutput{
			Code:    408,
			Message: "Request Timeout",
		}
		return 408, output, ""
	}

	//validasi jika user tidak ditemukan
	if err == gorm.ErrRecordNotFound {
		output := outputs.NotFoundOutput{
			Code:    404,
			Message: "Not Found: User not found",
		}
		return 404, output, ""
	}

	res, err := utils.ComparePassword(LoginUserRequestDTO.Password, user.Password)
	//validasi jika ada error saat compare password
	if err != nil {
		output := outputs.InternalServerErrorOutput{
			Code:    500,
			Message: "Internal Server Error: " + err.Error(),
		}
		return 500, output, ""
	}

	//validasi jika password salah
	if !res {
		output := outputs.UnauthorizedOutput{
			Code:    401,
			Message: "Unauthorized: Wrong Password",
		}
		return 401, output, ""
	}
	tokenString, err := utils.CreateJWTToken(user.ID)

	//validasi jika error saat create token
	if err != nil {
		output := outputs.InternalServerErrorOutput{
			Code:    500,
			Message: "Fail to create token: " + err.Error(),
		}
		return 500, output, ""
	}

	var mentor database.Mentors
	var IsMentor bool = false
	// err = db.Where("mentors.user_id = (?)", user.ID).Find(&mentor).Error

	// //validasi jika error saat mencari mentor
	// if err != nil {
	// 	output := outputs.InternalServerErrorOutput{
	// 		Code:    500,
	// 		Message: "Fail to find mentor: " + err.Error(),
	// 	}
	// 	return 500, output, ""
	// }
	mentor, err = repositories.FindMentorByUserID(user.ID.String())
	//validasi timeout
	if err == context.DeadlineExceeded {
		output := outputs.RequestTimeoutOutput{
			Code:    408,
			Message: "Request Timeout",
		}
		return 408, output, ""
	}

	if mentor.UserID == user.ID {
		IsMentor = true
	}

	output := outputs.LoginUserOutput{
		BaseOutput: outputs.BaseOutput{
			Code:    200,
			Message: "Success: Login Successful",
		},
		Data: responses.UserResponseDTO{
			ID:             user.ID,
			UserName:       user.Username,
			University:     user.University,
			Email:          user.Email,
			PhoneNumber:    user.Phone,
			Description:    user.Description,
			ProfilePicture: user.ProfilePicture,
			BOD:            user.BOD,
			IsActive:       user.IsActive,
			IsMentor:       IsMentor,
		},
	}
	return 200, output, tokenString
}

func GetUserByID(userID string) (int, interface{}) {
	// timeout, err := time.ParseDuration(os.Getenv("TIMEOUT"))
	// if err != nil {
	// 	output := outputs.InternalServerErrorOutput{
	// 		Code:    500,
	// 		Message: "Fail to find mentor: " + err.Error(),
	// 	}
	// 	return 500, output
	// }
	// ctx, cancel := context.WithTimeout(context.Background(), timeout)
	// defer cancel()

	// db := configs.GetDB().WithContext(ctx)
	// var user database.Users

	// //cari user menggunakan id
	// err = db.Table("users").Where("users.id = (?)", userID).First(&user).Error

	var user database.Users
	user, err := repositories.FindUserByUserID(userID)

	//validasi timeout
	if err == context.DeadlineExceeded {
		output := outputs.RequestTimeoutOutput{
			Code:    408,
			Message: "Request Timeout",
		}
		return 408, output
	}

	//validasi jika user tidak ditemukan
	if err == gorm.ErrRecordNotFound {
		output := outputs.NotFoundOutput{
			Code:    404,
			Message: "Not Found: User not found",
		}
		return 404, output
	}

	var mentor database.Mentors
	var IsMentor bool = false
	mentor, err = repositories.FindMentorByUserID(user.ID.String())

	//validasi timeout
	if err == context.DeadlineExceeded {
		output := outputs.RequestTimeoutOutput{
			Code:    408,
			Message: "Request Timeout",
		}
		return 408, output
	}

	if mentor.UserID == user.ID {
		IsMentor = true
	}

	output := outputs.GetUserByIDOutput{
		BaseOutput: outputs.BaseOutput{
			Code:    200,
			Message: "Success: Account found",
		},
		Data: responses.UserResponseDTO{
			ID:             user.ID,
			UserName:       user.Username,
			University:     user.University,
			Email:          user.Email,
			PhoneNumber:    user.Phone,
			Description:    user.Description,
			ProfilePicture: user.ProfilePicture,
			BOD:            user.BOD,
			IsActive:       user.IsActive,
			IsMentor:       IsMentor,
		},
	}
	return 200, output
}
