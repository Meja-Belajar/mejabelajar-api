import { BaseResponse } from "./base_response";

export interface RegisterUserResponse extends BaseResponse {
  data: {
    id: string;
    username: string;
    email: string;
    phone_number: string;
    description?: string;
    profile_picture: string;
    bod: string; 
    is_active: boolean;
    created_by: string; 
    updated_by: string;
    created_at: string; 
    updated_at: string;
  };
}

export interface LoginUserResponse extends BaseResponse {
  data: {
    id: string;
    username: string;
    email: string;
    phone_number: string;
    description?: string;
    profile_picture: string;
    bod: string; 
    is_active: boolean;
    created_by: string; 
    updated_by: string;
    created_at: string; 
    updated_at: string;
  };
}
