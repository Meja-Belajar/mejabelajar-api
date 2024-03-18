import { z } from "zod";

export const LoginUserSchema = z.object({
  email: z.string().email(),
  password: z.string()
      .min(8, { message: 'Password must be at least 8 characters' })
      .max(15, { message: 'Password must be at most 15 characters' })
})

export const RegisterUserSchema = z.object({ 
  name: z.string(),
  email: z.string().email(),
  password: z.string()
      .min(8, { message: 'Password must be at least 8 characters' })
      .max(15, { message: 'Password must be at most 15 characters' })
})
