import { Card, CardBody, CardHeader, Image, Spacer } from "@nextui-org/react";
import image from '../assets/image/aaron-burden-LNwn_A9RGHo-unsplash.jpg';

const Mentor = () => {
  return (
    <>
      <div className='w-full p-3'>
        <div className='w-full p-3 lato-bold text-2xl'>

          <h1 className='ml-3'>Find Your Mentor</h1>
        </div>
      
        <div className="w-full p-3 flex flex-row overflow-x-auto mt-3 no-scrollbar scroll-smooth">
          <div className="relative flex flex-row gap-6 sm:gap-8">
            <Card className="py-4 w-56">
              <CardHeader className="pb-0 pt-2 px-4 flex-col items-start flex-wrap">
                <p className="text-tiny uppercase font-bold">BINUSIAN B26</p>
                <div className="max-w-[240px] text-default-500 text-ellipsis overflow-hidden">
                  <small>Software Engineering, Algorithm </small>
                </div>
                <h4 className="font-bold text-large">Nama Mentor</h4>
              </CardHeader>
              <CardBody className="overflow-visible py-2">
                <Image
                  alt="Card background"
                  className="object-cover rounded-xl"
                  src={image}
                  width={270}
                />
              </CardBody>
            </Card>
            
            <Card className="py-4 w-56">
              <CardHeader className="pb-0 pt-2 px-4 flex-col items-start flex-wrap">
                <p className="text-tiny uppercase font-bold">BINUSIAN B26</p>
                <div className="max-w-[240px] text-default-500 text-ellipsis overflow-hidden">
                  <small>Software Engineering, Algorithm </small>
                </div>
                <h4 className="font-bold text-large">Nama Mentor</h4>
              </CardHeader>
              <CardBody className="overflow-visible py-2">
                <Image
                  alt="Card background"
                  className="object-cover rounded-xl"
                  src={image}
                  width={270}
                />
              </CardBody>
            </Card>

            <Card className="py-4 w-56">
              <CardHeader className="pb-0 pt-2 px-4 flex-col items-start flex-wrap">
                <p className="text-tiny uppercase font-bold">BINUSIAN B26</p>
                <small className="max-w-[240px] text-default-500 text-ellipsis overflow-hidden">Software Engineering, Algorithm </small>
                <h4 className="font-bold text-large">Nama Mentor</h4>
              </CardHeader>
              <CardBody className="overflow-visible py-2">
                <Image
                  alt="Card background"
                  className="object-cover rounded-xl"
                  src={image}
                  width={270}
                />
              </CardBody>
            </Card>
            <Card className="py-4 w-56">
              <CardHeader className="pb-0 pt-2 px-4 flex-col items-start flex-wrap">
                <p className="text-tiny uppercase font-bold">BINUSIAN B26</p>
                <small className="max-w-[240px] text-default-500 text-ellipsis overflow-hidden">Software Engineering, Algorithm </small>
                <h4 className="font-bold text-large">Nama Mentor</h4>
              </CardHeader>
              <CardBody className="overflow-visible py-2">
                <Image
                  alt="Card background"
                  className="object-cover rounded-xl"
                  src={image}
                  width={270}
                />
              </CardBody>
            </Card>
            <Card className="py-4 w-56">
              <CardHeader className="pb-0 pt-2 px-4 flex-col items-start flex-wrap">
                <p className="text-tiny uppercase font-bold">BINUSIAN B26</p>
                <small className="max-w-[240px] text-default-500 text-ellipsis overflow-hidden">Software Engineering, Algorithm </small>
                <h4 className="font-bold text-large">Nama Mentor</h4>
              </CardHeader>
              <CardBody className="overflow-visible py-2">
                <Image
                  alt="Card background"
                  className="object-cover rounded-xl"
                  src={image}
                  width={270}
                />
              </CardBody>
            </Card>
            <Card className="py-4 w-56">
              <CardHeader className="pb-0 pt-2 px-4 flex-col items-start flex-wrap">
                <p className="text-tiny uppercase font-bold">BINUSIAN B26</p>
                <small className="max-w-[240px] text-default-500 text-ellipsis overflow-hidden">Software Engineering, Algorithm </small>
                <h4 className="font-bold text-large">Nama Mentor</h4>
              </CardHeader>
              <CardBody className="overflow-visible py-2">
                <Image
                  alt="Card background"
                  className="object-cover rounded-xl"
                  src={image}
                  width={270}
                />
              </CardBody>
            </Card>

          </div>
        </div>
      </div> 
    </>
  )
}

export default Mentor;
