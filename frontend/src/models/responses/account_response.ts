export interface RegisterUserResponse {
  id: number;
  name: string;
  email: string;
  createdAt: string;
  updatedAt: string;
}

export interface LoginUserResponse {
  token: string;
  user: {
    id: number;
    name: string;
    email: string;
    createdAt: string;
    updatedAt: string;
  }
}