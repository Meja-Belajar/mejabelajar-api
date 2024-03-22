import { z } from "zod";

// schema for user login
export const LoginUserSchema = z.object({
  email: z.string().email(),
  password: z.string()
      .min(8, { message: 'Password must be at least 8 characters' })
      .max(15, { message: 'Password must be at most 15 characters' })
})

// schema for user registration
export const RegisterUserSchema = z.object({ 
  user_name: z.string(),
  email: z.string().email(),
  phone_number: z.string(),
  bod: z.string(),
  password: z.string()
      .min(8, { message: 'Password must be at least 8 characters' })
      .max(15, { message: 'Password must be at most 15 characters' }),
  confirm_password: z.string()
      .min(8, { message: 'Password must be at least 8 characters' })
      .max(15, { message: 'Password must be at most 15 characters' }),
  created_by: z.string()
})
