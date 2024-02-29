import React from 'react'
import image from '../assets/image/aaron-burden-LNwn_A9RGHo-unsplash.jpg'

const Offers = () => {
  return (
    <>
      <div className='w-full p-3'>
        <div className='w-full p-3 lato-bold text-2xl'>
          <h1 className='ml-3'>Best Offer</h1>
        </div>
        <div className='w-full border mt-3'>
          <img src={image} alt="" className='w-full max-h-64 object-cover'/>
        </div>
      </div> 
    </>
  )
}

export default Offers