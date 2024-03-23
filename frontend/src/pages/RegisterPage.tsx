import React, { useContext, useEffect, useReducer, useState } from 'react'
import { motion }from 'framer-motion'
import { exit, animate, initial } from '@assets/PageTransition'
import { Button, Input } from '@nextui-org/react'
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import { faEye, faEyeSlash } from '@fortawesome/free-solid-svg-icons'
import { Link, useNavigate } from 'react-router-dom'
import { registerService } from '@src/apis/services/userService'
import '@assets/global.css';
import Logo from '@src/components/Logo'
import { RegisterUserSchema } from '@src/models/zod/user_zod'
import { RegisterUserErrorValidation, RegisterUserRequest } from '@src/models/requests/user_request'
import { RegisterUserResponse } from '@src/models/responses/user_response'
import { maxDateUtil } from '@src/utils/dateUtil'
import { register } from 'module'

const FormReducer = (state: RegisterUserRequest, action: any) => {
  return {
    ...state,
    [action.name]: action.value
  };
};

const Register = () => {
  const [isVisible, setIsVisible] = useState<boolean>(false);
  const [loading, setLoading] = useState<boolean>(false);
  
  const [formData, setFormData] = useReducer(FormReducer, {} as RegisterUserRequest);
  const [formDataError, setFormDataError] = useState<RegisterUserErrorValidation>({} as RegisterUserErrorValidation);

  const [warn, setWarn] = useState<string>('');
  
  const navigate = useNavigate();
  
  const handleRegister = async () => {
    try {
      setLoading(true);

      const registerResponse = await registerService({ 
        user_name: formData.user_name, 
        email: formData.email, 
        password: formData.password,
        phone_number: formData.phone_number, 
        bod: formData.bod, 
        confirm_password: formData.confirm_password, 
        created_by: formData.created_by
      });

      // error validation already handled by registerService
      navigate('/login');

    } catch (error) {
      if(error instanceof Error) {
        window.scrollTo(0, 0);
        setWarn(error.toString());
      }
    } finally {
      setLoading(false);
    }
  }

  const handleSubmit: React.FormEventHandler<HTMLFormElement> = (e) => {
    e.preventDefault();
    console.log("in");
    const parsedUser = RegisterUserSchema.safeParse(formData);

    if (!parsedUser.success) {
      setLoading(true);
      const error = parsedUser.error;
      let newErrors = {};
      for (const issue of error.issues) {
        newErrors = {
          ...newErrors,
          [issue.path[0]]: issue.message,
        };
      }
      console.log(newErrors);
      setFormDataError(newErrors as RegisterUserErrorValidation);
      setLoading(false);
    } else {
      setFormDataError({} as RegisterUserErrorValidation);
      handleRegister();
    }
  }

  return (  
    <>
      <motion.div
        className='w-full flex items-center flex-col pb-36'
        initial={ initial }
        animate={ animate }
        exit={ exit }
      > 
        <nav className='w-full h-16 mt-2 flex justify-between items-center p-3 sm:p-7 mb-10'>
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
                name='name'
                type='name' 
                variant='bordered'
                className='mt-3 lato-regular' 
                label='Name'
                value={formData.user_name}
                errorMessage={formDataError.user_name && formDataError.user_name}
                onChange={(e) => {
                  setFormData({type: 'change', value: e.target.value, name: 'user_name'})
                  setFormData({type: 'change', value: e.target.value, name: 'created_by'})
                }}
              />
              <Input 
                name='email'
                type='email' 
                variant='bordered'
                className='lato-regular mt-3' 
                label='Email'
                value={formData.email}
                errorMessage={formDataError.email && formDataError.email}
                onChange={(e) => setFormData({type: 'change', value: e.target.value, name: 'email'})}
              />
              <Input 
                name='phone_number'
                type='text' 
                variant='bordered'
                className='lato-regular mt-3' 
                label='Phone Number'
                value={formData.phone_number}
                errorMessage={formDataError.phone_number && formDataError.phone_number}
                onChange={(e) => setFormData({type: 'change', value: e.target.value, name: 'phone_number'})}
              />
              <Input 
                name='bod'
                type='date' 
                variant='bordered'
                className='lato-regular mt-3' 
                classNames={{
                  label: "-mt-4 text-xs"
                }}
                label='Date of Birth'
                value={formData?.bod?.toString()} 
                key='outside'
                errorMessage={formDataError.bod && formDataError.bod}
                onChange={(e) => setFormData({type: 'change', value: e.target.value, name: 'bod'})}
                max={maxDateUtil()}
              />
              <Input 
                name='password'
                type={ isVisible ? "text" : "password"}
                variant='bordered' 
                label='Password'
                className='mt-3 lato-regular'
                value={formData.password}
                errorMessage={formDataError.password && formDataError.password}
                onChange={(e) => setFormData({type: 'change', value: e.target.value, name: 'password'})}
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
              <Input 
                name='confirmpassword'
                type={ isVisible ? "text" : "password"}
                variant='bordered' 
                label='Confirm Password'
                className='mt-3 lato-regular'
                value={formData.confirm_password}
                errorMessage={formDataError.confirm_password && formDataError.confirm_password}
                onChange={(e) => setFormData({type: 'change', value: e.target.value, name: 'confirm_password'})}
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