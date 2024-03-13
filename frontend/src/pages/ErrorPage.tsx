import { motion } from 'framer-motion';
import { animate, exit, initial } from '@assets/PageTransition';
import img from '@assets/images/image-removebg-preview.png'
import { Button } from '@nextui-org/react';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faArrowLeft } from '@fortawesome/free-solid-svg-icons';

const ErrorPage: React.FC = () => {
  return (
    <>
      <motion.div
        initial={initial}
        animate={animate}
        exit={exit}
        className='mt-3 p-10 flex items-center justify-center bg-gradient-to-t from-purple-500 via-purple-400 to-pink-500'
      >
        <div className='w-3/4 h-[80vh] flex flex-col-reverse sm:flex-row items-center justify-between'>
          <div className='flex flex-col items-start text-white'>
            <h1 className='text-6xl open-sans-600 mb-10'>404 ERROR</h1>
            <h2 className='text-4xl open-sans-300'>Seem like you get lost</h2>
            <Button
              className='bg-white drop-shadow-md shadow-md text-black mt-10 w-1/2 sm:w-1/3'
              startContent={
                <FontAwesomeIcon icon={faArrowLeft} className='mr-3' fade/>
              }
              onClick={() => window.history.back()}
            >
              BACK
            </Button>
          </div>
          <div className='mt-10'>
            <img src={img} alt="img" width="200"/>
          </div>
        </div>
      </motion.div>
    </>
  )
}

export default ErrorPage