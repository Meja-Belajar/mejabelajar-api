import React from 'react';
import '../assets/global.css';

import logo from '../../public/vite.svg';
import { AnimatePresence, motion } from 'framer-motion';

const Loading: React.FC = () => {
  
  return (
    <AnimatePresence>
      <motion.div
        initial={{ opacity: 0 }}
        animate={{ opacity: 1 }}
        exit={{ opacity: 0 }}
      >
        <div className="w-full h-[100vh] flex flex-col items-center justify-center bg-gradient-to-r from-blue-accent-300 via-purple-500 to-pink-500">
          <motion.div 
            className='border aspect-square p-10 pl-11 pr-11 rounded-full border-none'
            initial={{ 
              backgroundImage: 'linear-gradient(to right, #ec4899, #3b82f6)',
            }}
            animate={{ 
              backgroundImage: 'linear-gradient(to right, #3b82f6, #ec4899)',
            }}
            transition={{ repeat: Infinity, duration: 2, repeatType: 'reverse' }}
          >
            <img src={logo} alt="" width={100}/>
          </motion.div>
        </div>
      </motion.div>
    </AnimatePresence>
  );
};

export default Loading;