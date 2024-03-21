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
		Phone:    		AddUserRequestDTO.PhoneNumber,
		IsActive:    AddUserRequestDTO.IsActive,
		CreatedBy:      AddUserRequestDTO.CreatedBy,
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
			IsActive:    	user.IsActive,
			CreatedBy:      user.CreatedBy,
			ProfilePicture: user.ProfilePicture,
			UpdatedBy:      user.UpdatedBy,
			CreatedAt:      user.CreatedAt,
			UpdatedAt:      user.UpdatedAt,
		},
	}
	return 200, output

}
