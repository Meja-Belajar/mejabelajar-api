export interface RegisterUserRequest {
  user_name: string;
  email: string;
  phone_number: string;
  bod: Date;
  password: string;
  confirm_password: string;
  created_by: string;
}

export interface RegisterUserErrorValidation{
  name: string;
  email: string;
  phone_number: string;
  bod: string;
  password: string;
  confirm_password: string;
  created_by: string;
}

export interface LoginUserRequest {
  email: string;
  password: string;
}

export interface LoginUserErrorValidation {
  email: string;
  password: string;
}