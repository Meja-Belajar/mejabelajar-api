import { useNavigate } from 'react-router-dom';
import logo from '@public/logo.svg';
import '@assets/global.css';

const Logo: React.FC = () => {
  const navigate = useNavigate();
  
  return (
    <>
      <div 
        className='flex items-center justify-center gap-2 cursor-pointer' 
        onClick={() => navigate('/')}
      >
        <img src={logo} alt="logo" className='w-6 mb-1' />
        <h1 className='special-font text-blue-accent-400'>MejaBelajar</h1>
      </div>
    </>
  )
}

export default Logo