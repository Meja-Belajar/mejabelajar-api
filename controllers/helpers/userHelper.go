package helpers

import (
	"github.com/meja_belajar/configs"
	"github.com/meja_belajar/models/database"
	"github.com/meja_belajar/models/outputs"
	"github.com/meja_belajar/models/requests"
	"github.com/meja_belajar/models/responses"
	"github.com/meja_belajar/utils"
)

func RegisterUser(AddUserRequestDTO requests.RegisterUserRequestDTO) (int, interface{}) {
	if AddUserRequestDTO.Password != AddUserRequestDTO.ConfirmPassword {
		output := outputs.BadRequestOutput{
			Code:    400,
			Message: "Bad Request: Password and Confirm Password does not match",
		}
		return 400, output
	}

	hashedPassword, err := utils.HashPassword(AddUserRequestDTO.Password)
	if err != nil {
		output := outputs.InternalServerErrorOutput{
			Code:    500,
			Message: "Internal Server Error: " + err.Error(),
		}
		return 500, output
	}

	db := configs.GetDB()
	user := database.Users{
		Username:       AddUserRequestDTO.UserName,
		Email:          AddUserRequestDTO.Email,
		Password:       hashedPassword,
		Phone:          AddUserRequestDTO.PhoneNumber,
		Log: database.Log{
			IsActive:  AddUserRequestDTO.IsActive, // Example value for IsActive
			CreatedBy: AddUserRequestDTO.CreatedBy, // Example value for CreatedBy
		},
		ProfilePicture: AddUserRequestDTO.ProfilePicture,
	}
	err = db.Create(&user).Error
	if err != nil {
		output := outputs.InternalServerErrorOutput{
			Code:    500,
			Message: "Internal Server Error: " + err.Error(),
		}
		return 500, output
	}
	output := outputs.RegisterUserOutput{
		BaseOutput: outputs.BaseOutput{
			Code:    200,
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
	return 200, output

}

func LoginUser(LoginUserRequestDTO requests.LoginUserRequestDTO) (int, interface{}, string) {
	db := configs.GetDB()
	var user database.Users
	err := db.Where("email = ?", LoginUserRequestDTO.Email).First(&user).Error
	if err != nil {
		output := outputs.NotFoundOutput{
			Code:    404,
			Message: "Not Found: User not found",
		}
		return 404, output, ""
	}

	res, err := utils.ComparePassword(LoginUserRequestDTO.Password, user.Password)
	if err != nil {
		output := outputs.InternalServerErrorOutput{
			Code:    500,
			Message: "compare password fail Internal Server Error: " + err.Error(),
		}
		return 500, output, ""
	}
	if !res {
		output := outputs.UnauthorizedOutput{
			Code:    401,
			Message: "Unauthorized: Wrong Password",
		}
		return 401, output, ""
	}
	tokenString, err := utils.CreateJWTToken(user.ID)
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
	return 200, output, tokenString
}
