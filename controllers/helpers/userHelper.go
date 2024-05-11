package helpers

import (
	"github.com/meja_belajar/configs"
	"github.com/meja_belajar/models/database"
	"github.com/meja_belajar/models/outputs"
	"github.com/meja_belajar/models/requests"
	"github.com/meja_belajar/models/responses"
	"github.com/meja_belajar/utils"
	"gorm.io/gorm"
)

func RegisterUser(AddUserRequestDTO requests.RegisterUserRequestDTO) (int, interface{}) {
	//validasi password dan confirm password sama
	if AddUserRequestDTO.Password != AddUserRequestDTO.ConfirmPassword {
		output := outputs.BadRequestOutput{
			Code:    400,
			Message: "Bad Request: Password and Confirm Password does not match",
		}
		return 400, output
	}

	//validasi hash password
	hashedPassword, err := utils.HashPassword(AddUserRequestDTO.Password)
	if err != nil {
		output := outputs.InternalServerErrorOutput{
			Code:    500,
			Message: "Internal Server Error: " + err.Error(),
		}
		return 500, output
	}

	db := configs.GetDB()
	var user database.Users
	err = db.Where("email = ?", AddUserRequestDTO.Email).First(&user).Error
	//validasi email sudah ada atau belum
	if err == nil {
		output := outputs.ConflictOutput{
			Code:    409,
			Message: "Conflict: Email already exists",
		}
		return 409, output
	}
	user = database.Users{
		Username: AddUserRequestDTO.UserName,
		Email:    AddUserRequestDTO.Email,
		Password: hashedPassword,
		Phone:    AddUserRequestDTO.PhoneNumber,
		Log: database.Log{
			IsActive:  AddUserRequestDTO.IsActive,  // Example value for IsActive
			CreatedBy: AddUserRequestDTO.CreatedBy, // Example value for CreatedBy
		},
		ProfilePicture: AddUserRequestDTO.ProfilePicture,
	}
	err = db.Create(&user).Error
	//validasi error saat insert data
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
			Email:          user.Email,
			PhoneNumber:    user.Phone,
			ProfilePicture: user.ProfilePicture,
			Log: responses.LogResponseDTO{
				IsActive:  user.Log.IsActive,
				CreatedBy: user.Log.CreatedBy,
				UpdatedBy: user.Log.UpdatedBy,
				CreatedAt: user.Log.CreatedAt,
				UpdatedAt: user.Log.UpdatedAt,
			},
		},
	}
	return 201, output
}

func LoginUser(LoginUserRequestDTO requests.LoginUserRequestDTO) (int, interface{}, string) {
	db := configs.GetDB()
	var user database.Users
	err := db.Where("email = ?", LoginUserRequestDTO.Email).First(&user).Error
	//validasi user tidak ditemukan
	if err == gorm.ErrRecordNotFound {
		output := outputs.NotFoundOutput{
			Code:    404,
			Message: "Not Found: User not found",
		}
		return 404, output, ""
	}

	res, err := utils.ComparePassword(LoginUserRequestDTO.Password, user.Password)
	//validasi error saat proses compare password
	if err != nil {
		output := outputs.InternalServerErrorOutput{
			Code:    500,
			Message: "Internal Server Error: " + err.Error(),
		}
		return 500, output, ""
	}
	//validasi password salah
	if !res {
		output := outputs.UnauthorizedOutput{
			Code:    401,
			Message: "Unauthorized: Wrong Password",
		}
		return 401, output, ""
	}
	tokenString, err := utils.CreateJWTToken(user.ID)
	//validasi error saat create token
	if err != nil {
		output := outputs.InternalServerErrorOutput{
			Code:    500,
			Message: "Fail to create token: " + err.Error(),
		}
		return 500, output, ""
	}
	output := outputs.LoginUserOutput{
		BaseOutput: outputs.BaseOutput{
			Code:    200,
			Message: "Success: Login Successful",
		},
		Data: responses.UserResponseDTO{
			ID:             user.ID,
			UserName:       user.Username,
			Email:          user.Email,
			PhoneNumber:    user.Phone,
			ProfilePicture: user.ProfilePicture,
			Log: responses.LogResponseDTO{
				IsActive:  user.Log.IsActive,
				CreatedBy: user.Log.CreatedBy,
				UpdatedBy: user.Log.UpdatedBy,
				CreatedAt: user.Log.CreatedAt,
				UpdatedAt: user.Log.UpdatedAt,
			},
		},
	}
	return 200, output, tokenString
}
