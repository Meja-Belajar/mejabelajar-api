import { initial, animate, exit } from '@src/assets/PageTransition';
import { motion } from 'framer-motion';

// landing page for mentor
const MentorLanding = () => {
  return (
    <>
      <motion.div
        initial={initial}
        animate={animate}
        exit={exit}
        className='bg-white-accent-1'
      >
        <div>
          <h1>MentorLanding</h1>
        </div>
      </motion.div>
    </>

  )
}

export default MentorLanding