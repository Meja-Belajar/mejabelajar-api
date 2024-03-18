import { useParams } from "react-router-dom";
import { initial, animate, exit } from '@src/assets/PageTransition';
import { motion } from "framer-motion";

const SearchResult = () => {

  const { query } = useParams();
  return (
    <>
      <motion.div
        initial={initial}
        animate={animate}
        exit={exit}
        className='bg-white-accent-1'
      >
        <div>
          <h1>Search Result {query}</h1>
        </div>
      </motion.div>
    </>

  )
}

export default SearchResult