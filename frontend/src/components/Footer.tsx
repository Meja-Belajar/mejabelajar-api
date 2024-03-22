import React from 'react';
import Instagram from '@assets/images/social/instagram.svg';
import Twitter from '@assets/images/social/twitter.svg';
import Facebook from '@assets/images/social/facebook.svg';

const Footer: React.FC = () => {
  return (
    <>
      <footer className='pb-10 p-2 pt-8 bg-gradient-to-t from-blue-accent-300 via-purple-400 to-purple-500'>
        <div className='p-6 sm:p-8 flex flex-col gap-10 sm:gap-0 sm:flex-row flex-wrap items-start justify-between drop-shadow-md shadow-md m-2 sm:m-10 bg-white-accent-1 rounded-xl text-black mb-10 sm:mb-5'>
          <div className='flex flex-start flex-col'>
            <h1 className='mb-2 open-sans-600'>Company</h1>
            <ul className='opacity-80'>
              <li className='mb-2 ease-in-out duration-300 hover:underline'><a href="/error">About</a></li>
              <li className='mb-2 ease-in-out duration-300 hover:underline'><a href="/error">Overview</a></li>
              <li className='mb-2 ease-in-out duration-300 hover:underline'><a href="/error">For the License</a></li>
            </ul>
          </div>
          <div className=''>
            <h1 className='mb-2 open-sans-600'>Communities</h1>
            <ul className='opacity-80'>
              <li className='mb-2 ease-in-out duration-300 hover:underline'><a href="/error">For Mentor</a></li>
              <li className='mb-2 ease-in-out duration-300 hover:underline'><a href="/error">Developer</a></li>
              <li className='mb-2 ease-in-out duration-300 hover:underline'><a href="/error">Advertising</a></li>
              <li className='mb-2 ease-in-out duration-300 hover:underline'><a href="/error">Investors</a></li>
              <li className='mb-2 ease-in-out duration-300 hover:underline'><a href="/error">Vendors</a></li>
            </ul>
          </div>
          <div>
            <h1 className='mb-2 open-sans-600'>Resources</h1>
            <ul className='opacity-80'>
              <li className='mb-2 ease-in-out duration-300 hover:underline'><a href="/error">Supports</a></li>
              <li className='mb-2 ease-in-out duration-300 hover:underline'><a href="/error">Free Mobile App</a></li>
            </ul>
          </div>
          <div>
            <h1 className='mb-2 open-sans-600'>Useful links</h1>
            <ul className='opacity-80'>
              <li className='mb-2 ease-in-out duration-300 hover:underline'><a href="/error">FaQ</a></li>
              <li className='mb-2 ease-in-out duration-300 hover:underline'><a href="/error">Free Information</a></li>
              <li className='mb-2 ease-in-out duration-300 hover:underline'><a href="/error">Sports</a></li>
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
        <div className='mt-10 sm:mt-10 sm:m-7'>

          <div className='mt-2 flex flex-col gap-0'>
            <h1 className='opacity-80 ml-3'>©Copyright 2024 Meja Belajar </h1>
            <small className='ml-3'>Jakarta, Indonesia 11480</small>
          </div>
        </div>
      </footer>
    </>
  )
}

export default Footer