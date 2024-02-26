import React from 'react';
import Instagram from '../../public/logo-instagram-svgrepo-com.svg';
import Twitter from '../../public/twitter-svgrepo-com.svg';
import Facebook from '../../public/facebook-svgrepo-com.svg';

const Footer: React.FC = () => {
  return (
    <>
      <footer className='pb-10 p-2 mt-10 pt-8 bg-gradient-to-l from-blue-accent-300 via-purple-500 to-purple-400 text-white'>
        <div className='p-3 sm:p-5 flex flex-row flex-wrap items-start justify-between'>
          <div className='flex flex-start flex-col'>
            <h1 className='mb-2 lato-bold'>Company</h1>
            <ul className='lato-reguler opacity-80'>
              <li className='mb-2 ease-in-out duration-300 hover:underline'><a href="/error">About</a></li>
              <li className='mb-2 ease-in-out duration-300 hover:underline'><a href="/error">Overview</a></li>
              <li className='mb-2 ease-in-out duration-300 hover:underline'><a href="/error">For the License</a></li>
            </ul>
          </div>
          <div className=''>
            <h1 className='mb-2 lato-bold'>Communities</h1>
            <ul className='lato-reguler opacity-80'>
              <li className='mb-2 ease-in-out duration-300 hover:underline'><a href="/error">For Mentor</a></li>
              <li className='mb-2 ease-in-out duration-300 hover:underline'><a href="/error">Developer</a></li>
              <li className='mb-2 ease-in-out duration-300 hover:underline'><a href="/error">Advertising</a></li>
              <li className='mb-2 ease-in-out duration-300 hover:underline'><a href="/error">Investors</a></li>
              <li className='mb-2 ease-in-out duration-300 hover:underline'><a href="/error">Vendors</a></li>
            </ul>
          </div>
          <div>
            <h1 className='mb-2 lato-bold'>Useful links</h1>
            <ul className='lato-reguler opacity-80'>
              <li className='mb-2 ease-in-out duration-300 hover:underline'><a href="/error">Supports</a></li>
              <li className='mb-2 ease-in-out duration-300 hover:underline'><a href="/error">Free Mobile App</a></li>
            </ul>
          </div>
          <div className='flex flex-row gap-4 mt-10 sm:mt-0'>
            <div className='p-3 w-10 h-10 flex items-center justify-center rounded-full border border-black cursor-pointer transition ease-linear hover:border-purple-500'>
              <a href="#" target='_blank' rel='noreferrer' className='w-full'>
                <img src={Instagram} alt="instagram" className='w-full h-full' />
              </a>
            </div>
            <div className='p-3 w-10 h-10 flex items-center justify-center rounded-full border border-black cursor-pointer transition ease-linear hover:border-purple-500'>
              <a href="" target='_blank'  rel='noreferrer' className='w-full'>
                <img src={Twitter} alt="twitter" className='w-full h-full' />
              </a>
            </div>
            <div className='p-3 w-10 h-10 flex items-center justify-center rounded-full border border-black cursor-pointer transition ease-linear hover:border-purple-500'>
              <a href="" target='_blank'  rel='noreferrer' className='w-full'>
                <img src={Facebook} alt="facebook" className='w-full h-full' />
              </a>
            </div>
          </div>
        </div>
        <div className=''>
          <div className='flex items-center justify-center p-4'>
            <div className='border-b border-black w-full'></div>
          </div>
          <h1 className='mt-5 ml-4 opacity-80 lato-reguler'>© 2024 MejaBelajar</h1>
        </div>
      </footer>
    </>
  )
}

export default Footer