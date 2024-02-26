import React, { useContext, useEffect, useState } from 'react'
import { motion }from 'framer-motion'
import { exit, animate, initial } from '../assets/PageTransition'
import { Button, Input } from '@nextui-org/react'
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import { faEye, faEyeSlash } from '@fortawesome/free-solid-svg-icons'
import { Link, useNavigate } from 'react-router-dom'
import { loginService } from '../../services/user_service'

import { UserContext } from '../context/UserContext'
import '../assets/global.css';
import Logo from '../utils/Logo'

const Login: React.FC = () => {
  const { login, setLogin } = useContext(UserContext);

  const [isVisible, setIsVisible] = useState<boolean>(false);
  const [email, setEmail] = useState<string>('');
  const [password, setPassword] = useState<string>('');
  const [loading, setLoading] = useState<boolean>(false);

  const [warn, setWarn] = useState<string>('');

  const navigate = useNavigate();

  useEffect(() => {
    // If user is already logged in, redirect to home
    if(login && login.status === 200) {
      navigate('/');
    }  
  }, [])

  const formSubmit = async (e: any) => {
    e.preventDefault();
    try {
      setLoading(true);

      const data = await loginService(email, password);

      console.log(data);

      if(data && setLogin && data.status === 200){
        setLogin(data);
        navigate('/')
      } else {
        setWarn('Invalid email or password');
        setLoading(false);
      }
    } catch (error) {
      console.error(error);
      setWarn('An error occurred during login');
      setLoading(false);
    }
  }

  return (  
    <>
      <motion.div
        className='w-full h-[100vh] flex items-center justify-center flex-col'
        initial={ initial }
        animate={ animate }
        exit={ exit }
      > 
        <nav className='w-full h-16 mt-2 absolute top-0 flex justify-between items-center p-3 sm:p-7'>
          <Logo />
          <div>
            <Link className='lato-regular p-3 transition ease-soft-spring hover:text-blue-accent-300' to='/'>HOME</Link>
            <Link className='lato-regular p-3 transition ease-soft-spring hover:text-blue-accent-300' to='/register'>REGISTER</Link>
          </div>
        </nav>

        <form 
          className='w-[90%] md:w-1/3 bg-white rounded-lg p-5 drop-shadow-2xl'
          onSubmit={(e) => formSubmit(e)}
        >
          <div className='m-3'>
            <h1 className='lato-bold text-xl'>Welcome back!</h1>
            <p className='lato-regular text-sm'>create your next courses with our best perform mentor</p>
            <p className='mb-3 lato-bold text-red-600 mt-2'>{warn}</p>
          </div>
          <div className='m-3 mt-8'>
            <Input 
              type='email' 
              variant='bordered'
              className='lato-regular' 
              label='Email / Username'
              value={email}
              onChange={(e) => setEmail(e.target.value)}
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
            <Link to="/" className='text-xs underline-offset-2 underline decoration-transparent hover:decoration-black'>Forget Password ?</Link>
          </div>

          <div className='m-3 flex items-center justify-center flex-col'>
            <Button color='default' variant='solid' className='bg-blue-accent-300 text-black w-full lato-regular' type='submit' isLoading={loading}>Login</Button>
          </div>
        </form>

      </motion.div>
    </>
  )
}

export default Login