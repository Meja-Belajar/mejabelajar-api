import React, { useContext } from 'react'
import { motion } from 'framer-motion'
import Navigation from '../../components/Navigation'
import { exit, animate, initial } from '../../assets/PageTransition'
import '../../assets/global.css';
import icon from '../../../public/vite.svg'
import { faArrowLeft, faArrowRight } from '@fortawesome/free-solid-svg-icons';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import BecomeMentor from '../../components/BecomeMentor';
import Mentor from '../../components/Mentor';
import { UserContext } from '../../context/UserContext';
import Carousel from '../../components/Carousel';

const Landing = () => {

  return (
    <>
      <motion.div
        initial={initial}
        animate={animate}
        exit={exit}
      >

        <main className='mt-3'>
          <section className='border w-full p-5 pb-12 bg-gradient-to-r from-blue-accent-300 via-purple-500 to-pink-500 text-white'>
            <div className='text-xl md:p-3'>
              <h1 className='lato-bold'>Hay, Friends 👋!</h1>
              <h3 className='text-sm mt-2'>What would you like to learn about today? </h3>
            </div>
          </section>

          <section className='black rounded-xl drop-shadow-lg w-full bg-white relative -top-3 flex flex-col sm:flex-row items-center justify-between'>
            <h1 className='p-3 lato-black md:ml-10 mt-5 mb-3 sm:mb-0 sm:mt-0'>Best option<span className='text-blue-accent-300'> FOR YOU🫰 </span></h1>
            <div className='flex gap-3 flex-col p-3 sm:flex-row sm:justify-between sm:p-5 lg:w-3/4 items-center w-full md:mr-10'>

              <div className='pr-12 w-full p-4 lato-bold rounded-xl border-2 border-blue-accent-100 flex flex-row items-center gap-3 cursor-pointer transition ease-out hover:bg-blue-accent-300 hover:bg-opacity-50'>
                <img src={icon} alt="icon tutor" className='w-8'/>
                <h1>
                  Tutoring Class
                </h1>
                <FontAwesomeIcon icon={faArrowRight} fade className='sm:hidden absolute right-10'/>
              </div>
              <div className='pr-12 w-full p-4 lato-bold rounded-xl border-2 border-blue-accent-100 flex flex-row items-center gap-3 cursor-pointer transition ease-out hover:bg-blue-accent-300 hover:bg-opacity-50'>
                <img src={icon} alt="icon tutor" className='w-8'/>
                <h1>
                  Mentoring 1-to-1
                </h1>
                <FontAwesomeIcon icon={faArrowRight} fade className='sm:hidden absolute right-10'/>
              </div>
              <div className='pr-12 w-full p-4 lato-bold rounded-xl border-2 border-blue-accent-100 flex flex-row items-center gap-3 cursor-pointer transition ease-out hover:bg-blue-accent-300 hover:bg-opacity-50'>
                <img src={icon} alt="icon tutor" className='w-8'/>
                <h1>
                  Scheduled Mentoring
                </h1>
                <FontAwesomeIcon icon={faArrowRight} fade className='sm:hidden absolute right-10'/>

              </div>
            </div>
             
          </section>

          <section className='w-full'>

          </section>
          <Carousel />
          <BecomeMentor />
          <Mentor />

          <div>
            
          </div>
        </main>    


      </motion.div>
    </>
  )
}

export default Landing