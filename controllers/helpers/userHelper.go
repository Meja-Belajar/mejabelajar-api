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
		return utils.HandleBadRequest("Password and Confirm Password must be the same")
	}

	//validasi hash password
	hashedPassword, err := utils.HashPassword(AddUserRequestDTO.Password)
	//validasi hash password
	if err != nil {
		return utils.HandleInternalServerError(err)
	}

	var user database.Users
	user, err = repositories.FindUserByEmail(AddUserRequestDTO.Email)

	//validasi timeout
	if err == context.DeadlineExceeded {
		return utils.HandleTimeout(err)
	}

	//validasi email belum terdaftar
	if err == nil {
		return utils.HandleTimeout(err)
	}

	// buat user baru
	bod, err := time.Parse("2006-01-02T15:04:05Z", AddUserRequestDTO.BOD)
	if err != nil {
		return utils.HandleBadRequest("Invalid BOD format")
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
		return utils.HandleTimeout(err)
	}

	//validasi create user
	if err != nil {
		return utils.HandleInternalServerError(err)
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
		code, output := utils.HandleTimeout(err)
		return code, output, ""
	}

	//validasi jika user tidak ditemukan
	if err == gorm.ErrRecordNotFound {
		code, output := utils.HandleNotFound("User")
		return code, output, ""
	}

	res, err := utils.ComparePassword(LoginUserRequestDTO.Password, user.Password)
	//validasi jika ada error saat compare password
	if err != nil {
		code, output := utils.HandleInternalServerError(err)
		return code, output, ""
	}

	//validasi jika password salah
	if !res {
		code, output := utils.HandleBadRequest("Invalid Password")
		return code, output, ""
	}
	tokenString, err := utils.CreateJWTToken(user.ID)

	//validasi jika error saat create token
	if err != nil {
		code, output := utils.HandleInternalServerError(err)
		return code, output, ""
	}

	var mentor database.Mentors
	var IsMentor bool = false

	mentor, err = repositories.FindMentorByUserID(user.ID.String())
	//validasi timeout
	if err == context.DeadlineExceeded {
		code, output := utils.HandleTimeout(err)
		return code, output, ""
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
	var user database.Users
	user, err := repositories.FindUserByUserID(userID)

	//validasi timeout
	if err == context.DeadlineExceeded {
		return utils.HandleTimeout(err)
	}

	//validasi jika user tidak ditemukan
	if err == gorm.ErrRecordNotFound {
		return utils.HandleNotFound("User")
	}

	var mentor database.Mentors
	var IsMentor bool = false
	mentor, err = repositories.FindMentorByUserID(user.ID.String())

	//validasi timeout
	if err == context.DeadlineExceeded {
		return utils.HandleTimeout(err)
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
