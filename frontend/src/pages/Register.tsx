import React, { useState } from 'react'
import { motion }from 'framer-motion'
import { exit, animate, initial } from '../assets/PageTransition'
import { Button, Input } from '@nextui-org/react'
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import { faEye, faEyeSlash } from '@fortawesome/free-solid-svg-icons'
import { Link, useNavigate } from 'react-router-dom'
import { login } from '../../api/api'
import logo from '../../public/vite.svg'

const Register: React.FC = () => {
  const [isVisible, setIsVisible] = useState<boolean>(false);
  const [loading, setLoading] = useState<boolean>(false);
  
  const [name, setName] = useState<string>('');
  const [email, setEmail] = useState<string>('');
  const [password, setPassword] = useState<string>('');
  
  
  const [warn, setWarn] = useState<string>('');

  const navigate = useNavigate();

  const formSubmit = async (e: any) => {
    e.preventDefault();
    setLoading(true);

    const data = await login(email, password);
    console.log(data);

    if(data.status === 200){
      navigate('/')
    } else {
      setWarn('Invalid email or password');
      setLoading(false);
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
        <nav className='w-full h-16 border border-black flex justify-between items-center p-3 sm:p-7'>
          <div className='flex items-center justify-center gap-2'>
            <img src={logo} alt="logo" className='w-6 mb-1' />
            <h1 className='special-font text-yellow-800'>MejaBelajar</h1>
          </div>
          <div>
            <Link className='lato-regular p-3' to='/'>HOME</Link>
            <Link className='lato-regular p-3' to='/login'>LOGIN</Link>
          </div>
        </nav>

        <div className='w-full h-full flex items-center justify-center flex-col'>
          <form 
            className='w-[90%] md:w-1/3 bg-white rounded-lg p-5 drop-shadow-2xl '
            onSubmit={(e) => formSubmit(e)}
          >
            <div className='m-3'>
              <h1 className='lato-bold text-xl'>Hay 👋, let's become our family!</h1>
              <p className='lato-regular text-sm'>start your journey with our best perform mentor</p>
              <p className='mb-3 lato-bold text-red-600 mt-2'>{warn}</p>
            </div>
            <div className='m-3 mt-8'>
              <Input 
                type='email' 
                variant='bordered'
                className='lato-regular' 
                label='Email'
                value={email}
                onChange={(e) => setEmail(e.target.value)}
              />
              <Input 
                type='name' 
                variant='bordered'
                className='mt-3 lato-regular' 
                label='Name'
                value={name}
                onChange={(e) => setName(e.target.value)}
              />
              <Input 
                type={ isVisible ? "text" : "password"}
                variant='bordered' 
                label='Password'
                className='mt-3 lato-regular'
                value={password}
                onChange={(e) => setPassword(e.target.value)}
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
              <Button color='default' variant='solid' className='w-full lato-regular' type='submit' isLoading={loading}>Register</Button>
            </div>
          </form>
        </div>
      </motion.div>
    </>
  )
}

export default Register