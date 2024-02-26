import React from 'react'
import { motion } from 'framer-motion'
import Navigation from '../components/Navigation'
import { exit, animate, initial } from '../assets/PageTransition'
import '../assets/global.css';
import icon from '../../public/vite.svg'
import { faArrowLeft, faArrowRight } from '@fortawesome/free-solid-svg-icons';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';

const Landing = () => {
  return (
    <>
      <motion.div
        initial={initial}
        animate={animate}
        exit={exit}
      >
        <Navigation />  

        <main className='mt-3'>
          <section className='border w-full p-5 pb-12 bg-gradient-to-r from-indigo-500 via-purple-500 to-pink-500 text-white'>
            <div className='text-xl md:p-3'>
              <h1 className='lato-bold'>Halo friends</h1>
              <h3>Lorem ipsum Lorem ipsum </h3>
            </div>
          </section>

          <section className='black rounded-t-xl w-full bg-white relative -top-3 flex flex-col sm:flex-row items-center justify-between'>
            <h1 className='p-3 lato-black md:ml-10'>Ruang <span className='text-blue-accent-200'>Belajar</span></h1>
            <div className='flex gap-3 flex-col p-3 sm:flex-row sm:justify-between sm:p-5 lg:w-1/2 items-center w-full md:mr-10'>

              <div className='pr-12 w-full p-4 lato-bold rounded-xl border-2 border-blue-accent-100 flex flex-row items-center gap-3 cursor-pointer transition ease-out hover:bg-blue-accent-300 hover:bg-opacity-50'>
                <img src={icon} alt="icon tutor" />
                <h1>
                  Scheduled Tutor
                </h1>
                <FontAwesomeIcon icon={faArrowRight} fade className='sm:hidden absolute right-10'/>
              </div>
              <div className='pr-12 w-full p-4 lato-bold rounded-xl border-2 border-blue-accent-100 flex flex-row items-center gap-3 cursor-pointer transition ease-out hover:bg-blue-accent-300 hover:bg-opacity-50'>
                <img src={icon} alt="icon tutor" />
                <h1>
                  Scheduled Tutor
                </h1>
                <FontAwesomeIcon icon={faArrowRight} fade className='sm:hidden absolute right-10'/>
              </div>
              <div className='pr-12 w-full p-4 lato-bold rounded-xl border-2 border-blue-accent-100 flex flex-row items-center gap-3 cursor-pointer transition ease-out hover:bg-blue-accent-300 hover:bg-opacity-50'>
                <img src={icon} alt="icon tutor" />
                <h1>
                  Scheduled Tutor
                </h1>
                <FontAwesomeIcon icon={faArrowRight} fade className='sm:hidden absolute right-10'/>

              </div>
            </div>
          </section>
        </main>    
      </motion.div>
    </>
  )
}

export default Landing