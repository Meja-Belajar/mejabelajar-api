import React, { useContext } from 'react'
import { motion } from 'framer-motion'
import Navigation from '@components/Navigation'
import { exit, animate, initial } from '@assets/PageTransition'
import '@assets/global.css';
import icon from '@public/logo.svg'
import { faArrowLeft, faArrowRight, faArrowUpRightFromSquare } from '@fortawesome/free-solid-svg-icons';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import BecomeMentor from '@components/BecomeMentor';
import Mentor from '@components/Mentor';
import { UserContext } from '@contexts/UserContext';
import Carousel from '@components/Carousel';
import { Button, Skeleton } from '@nextui-org/react';
import { Link, useNavigate } from 'react-router-dom';
import Guest from '@pages/landing/Guest';
import Footer from '@components/Footer';

const Landing = () => {
  const navigate = useNavigate();

  return <Guest />

  return (  
    <>
      <Navigation />
      <motion.div
        initial={initial}
        animate={animate}
        exit={exit}
        className='bg-white-accent-1'
        >

        <main className='mt-3'>
          <section className='border w-full p-5 pb-12 bg-gradient-to-r from-blue-accent-300 via-purple-500 to-pink-500 text-white'>
            <div className='text-xl md:p-3'>
              <h1 className='open-sans-600'>Hay, Friends 👋!</h1>
              <h3 className='text-sm mt-1'>What would you like to learn about today? </h3>
            </div>
          </section>

          <section className='black rounded-xl drop-shadow-lg w-full bg-white relative -top-3 flex flex-col sm:flex-row items-center justify-between'>
            <h1 className='p-3 open-sans-600 md:ml-10 mt-5 mb-3 sm:mb-0 sm:mt-0'>Best option<span className='text-blue-accent-300'> FOR YOU🫰 </span></h1>
            <div className='flex gap-3 flex-col p-3 sm:flex-row sm:justify-between sm:p-5 lg:w-3/4 items-center w-full md:mr-10'>

              <div className='pr-12 w-full p-4 open-sans-600 rounded-xl border-2 border-blue-accent-100 flex flex-row items-center gap-3 cursor-pointer transition ease-out hover:bg-blue-accent-300 hover:bg-opacity-50'>
                <img src={icon} alt="icon tutor" className='w-8'/>
                <h1>
                  Tutoring Class
                </h1>
                <FontAwesomeIcon icon={faArrowRight} fade className='sm:hidden absolute right-10'/>
              </div>
              <div className='pr-12 w-full p-4 open-sans-600 rounded-xl border-2 border-blue-accent-100 flex flex-row items-center gap-3 cursor-pointer transition ease-out hover:bg-blue-accent-300 hover:bg-opacity-50'>
                <img src={icon} alt="icon tutor" className='w-8'/>
                <h1>
                  Mentoring 1-to-1
                </h1>
                <FontAwesomeIcon icon={faArrowRight} fade className='sm:hidden absolute right-10'/>
              </div>
              <div className='pr-12 w-full p-4 open-sans-600 rounded-xl border-2 border-blue-accent-100 flex flex-row items-center gap-3 cursor-pointer transition ease-out hover:bg-blue-accent-300 hover:bg-opacity-50'>
                <img src={icon} alt="icon tutor" className='w-8'/>
                <h1>
                  Scheduled Mentoring
                </h1>
                <FontAwesomeIcon icon={faArrowRight} fade className='sm:hidden absolute right-10'/>

              </div>
            </div>
             
          </section>


          <Carousel />
          <BecomeMentor />
          <Mentor />

          <div className='mt-10 w-full flex items-center justify-center'>
            <div className='w-full m-5 sm:m-10 p-8 flex flex-col sm:flex-row justify-between shadow-md drop-shadow-lg sm:h-96 rounded-3xl sm:gap-10 bg-gradient-to-r from-purple-500 via-purple-400 to-blue-accent-300'>
              <div className='p-3 w-full h-44 sm:h-auto mb-10 sm:mb-0 sm:w-1/3 rounded-2xl bg-white-accent-1 drop-shadow-md shadow-md'>
                <Skeleton className="w-full flex flex-col rounded-lg">
                  <div className="h-[5vh] rounded-lg bg-default-200" />
                </Skeleton>
                <Skeleton className="w-3/4 flex flex-col mt-3 rounded-lg">
                  <div className="h-[5vh] rounded-lg bg-default-200" />
                </Skeleton>
                <Skeleton className="w-3/4 flex flex-col mt-3 rounded-lg">
                  <div className="h-[5vh] rounded-lg bg-default-200" />
                </Skeleton>
                <Skeleton className="w-1/2 flex flex-col mt-3 rounded-lg">
                  <div className="h-[5vh] rounded-lg bg-default-200" />
                </Skeleton>
              </div>
              <div className='p-1 sm:p-3 w-full text-white-accent-1 sm:w-2/3 flex justify-center flex-col sm:pr-10'>
                <h1 className='text-3xl p-3'>Enhance Academic Performance</h1>
                <h1 className='text-xl sm:text-justify p-2 pl-3'>Meja Belajar assists university students in enhancing their academic performance with the aid of our mentors. Whenever you're ready, our mentors are always prepared to provide assistance.</h1>
                <div className='mt-10 flex items-center'>
                  <Link to='/announcement' className='peer p-3 underline-offset-4 underline decoration-transparent transition ease-linear duration-300 hover:decoration-white-accent-1'>
                    learn more 
                  </Link>
                  <FontAwesomeIcon icon={faArrowUpRightFromSquare} fade className='peer-hover:scale-110'/>
                </div>
              </div>
            </div>
          </div>
          <div className='mt-32'>
            <div className='w-full flex items-center justify-center mb-14'>
              <h1 className='text-5xl text-purple-500 open-sans-300 text-center'>What are you waiting for?</h1>
            </div>
            <div className='pb-28 mt-10 p-2 pt-8 bg-gradient-to-t from-purple-500 via-purple-400 to-white-accent-1'>
              <div className='w-full  flex items-center justify-center'>
                <Button 
                  className='w-1/2 sm:w-1/6 h-14 rounded-xl text-purple-500 bg-white-accent-1 drop-shadow-lg shadow-sm'
                  endContent={<FontAwesomeIcon icon={faArrowRight} className='ml-1' fade />}
                  onClick={() => navigate('/register')}
                  >
                  Get started for free
                </Button>
              </div>
            </div>

          </div>
        </main>    


      </motion.div>
      <Footer/>
    </>
  )
}

export default Landing