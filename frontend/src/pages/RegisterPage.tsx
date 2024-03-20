import React, { useContext, useEffect, useReducer, useState } from 'react'
import { motion }from 'framer-motion'
import { exit, animate, initial } from '@assets/PageTransition'
import { Button, Input } from '@nextui-org/react'
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import { faEye, faEyeSlash } from '@fortawesome/free-solid-svg-icons'
import { Link, useNavigate } from 'react-router-dom'
import { registerService } from '@src/apis/services/userService'
import logo from '@public/vite.svg'
import { UserContext } from '@contexts/UserContext'
import '@assets/global.css';
import Logo from '@src/components/Logo'
import { RegisterUserSchema } from '@src/models/zod/user'

interface FormFormat {
  name: string,
  email: string,
  password: string,
}

const FormReducer = (state: FormFormat, action: any) => {
  return {
    ...state,
    [action.name]: action.value
  };
};

const Register = () => {
  const [isVisible, setIsVisible] = useState<boolean>(false);
  const [loading, setLoading] = useState<boolean>(false);
  
  const [formData, setFormData] = useReducer(FormReducer, {} as FormFormat);
  const [formDataError, setFormDataError] = useState<FormFormat>({} as FormFormat);

  const [warn, setWarn] = useState<string>('');

  const navigate = useNavigate();
  
  const handleSubmit: React.FormEventHandler<HTMLFormElement> = (e) => {
    e.preventDefault();

    const handleRegister = async ({ name, email, password}: FormFormat) => {
      setLoading(true);

      const data = await registerService({ name, email, password });
      setLoading(false);

      if(data.status === 200){
        navigate('/')
      } else {
        setWarn('Invalid email or password');
      }
    }

    const parsedUser = RegisterUserSchema.safeParse(formData);
    console.log(parsedUser);

    if (!parsedUser.success) {
      const error = parsedUser.error;
      let newErrors = {};
      for (const issue of error.issues) {
        newErrors = {
          ...newErrors,
          [issue.path[0]]: issue.message,
        };
      }
      setFormDataError(newErrors as FormFormat);
    } else {
      handleRegister(formData);
      setFormDataError({} as FormFormat);
    }

  }

  return (  
    <>
      <motion.div
        className='w-full h-[100vh] flex items-center flex-col'
        initial={ initial }
        animate={ animate }
        exit={ exit }
      > 
        <nav className='w-full h-16 mt-2 flex justify-between items-center p-3 sm:p-7'>
          <Logo />
          <div>
            <Link className='lato-regular p-3 transition ease-soft-spring hover:text-blue-accent-300' to='/'>HOME</Link>
            <Link className='lato-regular p-3 transition ease-soft-spring hover:text-blue-accent-300' to='/login'>LOGIN</Link>
          </div>
        </nav>

        <div className='w-full h-full flex items-center justify-center flex-col'>
          <form 
            className='w-[90%] lg:w-1/3 bg-white rounded-lg p-5 drop-shadow-2xl '
            onSubmit={ handleSubmit }
          >
            <div className='m-3'>
              <h1 className='lato-bold text-xl'>Hay 👋, let's become our family!</h1>
              <p className='lato-regular text-sm'>start your journey with our best perform mentor</p>
              <p className='mb-3 lato-bold text-red-600 mt-2'>{warn}</p>
            </div>
            <div className='m-3 mt-8'>
              <Input 
                name='email'
                type='email' 
                variant='bordered'
                className='lato-regular' 
                label='Email'
                value={formData.email}
                onChange={setFormData}
              />
              {formDataError.email && <p className='text-sm p-1 mt-1 lato-regular text-red-600'>{formDataError.email}</p>}
              <Input 
                name='name'
                type='name' 
                variant='bordered'
                className='mt-3 lato-regular' 
                label='Name'
                value={formData.name}
                onChange={setFormData}
              />
              {formDataError.name && <p className='text-sm p-1 mt-1 lato-regular text-red-600'>{formDataError.name}</p>}
              <Input 
                name='password'
                type={ isVisible ? "text" : "password"}
                variant='bordered' 
                label='Password'
                className='mt-3 lato-regular'
                value={formData.password}
                onChange={setFormData}
                endContent={
                  <button className="focus:outline-none" type="button" onClick={() => setIsVisible(!isVisible)}>
                    { 
                      isVisible ? (
                        <FontAwesomeIcon icon={faEyeSlash} className='opacity-60'/>  
                      ) : (
                        <FontAwesomeIcon icon={faEye} className='opacity-60'/>  
                      ) 
                    }
                  </button>
                }      
              />
              {formDataError.password && <p className='text-sm p-1 mt-1 lato-regular text-red-600'>{formDataError.password}</p>}

            </div>

            <div className='m-3 pt-2 pb-2 flex items-end justify-end'>
              <Link to="/login" className='text-xs underline-offset-2 underline decoration-transparent hover:decoration-black'>Already have an account ?</Link>
            </div>

            <div className='m-3 flex items-center justify-center flex-col'>
              <Button type='submit' color='default' variant='solid' className='w-full lato-regular bg-blue-accent-300 text-black' isLoading={loading}>Register</Button>
            </div>
          </form>
        </div>
      </motion.div>
    </>
  )
}

export default Register