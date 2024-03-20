import { initial, animate, exit } from '@src/assets/PageTransition';
import { motion } from 'framer-motion';

const SearchDefault = () => {
  return (
    <>
      <motion.div
        initial={initial}
        animate={animate}
        exit={exit}
        className='bg-white-accent-1'
      >
        <div>
          <h1>Search Root Page</h1>
        </div>
      </motion.div>
    </>

  )
}

export default SearchDefault