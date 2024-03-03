import React from 'react'
import image from '../assets/image/aaron-burden-LNwn_A9RGHo-unsplash.jpg'
import ss from '../assets/image/mkmui_the_teacher_stands_in_front_of_the_classroom_pointing_at__5af4cbd1-2169-49ba-9e76-127695c1f9a5.png'
import '../assets/global.css';

const Carousel = () => {
  return (
    <div className='relative'>
      <div className='w-full mt-3'>
        <div className='w-full p-3 lato-bold text-2xl'>
          <h1 className='ml-3'>Best Offer</h1>
        </div>

        <div className='w-full flex items-center justify-center'>

          <div className="w-[98vw] flex flex-row items-center mt-3 overflow-x-auto no-scrollbar overflow-hidden">
            <div className="relative flex flex-row animate-carousel">
              <div className='w-[98vw] max-h-[50vh] flex items-center object-cover sm:object-none'>
                <img src={ss} alt="img" className='w-full h-full sm:h-auto' />
              </div>
              <div className='w-[98vw] max-h-[50vh] flex items-center object-cover sm:object-none'>
                <img src={image} alt="img" className='w-full h-full sm:h-auto' />
              </div>
              <div className='w-[98vw] max-h-[50vh] flex items-center object-cover sm:object-none'>
                <img src={ss} alt="img" className='w-full h-full sm:h-auto' />
              </div>
            </div>
          </div>
        </div>

      </div> 
    </div>
  )
}

export default Carousel