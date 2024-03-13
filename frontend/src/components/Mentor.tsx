import { Card, CardBody, CardHeader, Image, Skeleton } from "@nextui-org/react";
import { useEffect, useState } from "react";
import { getMentor } from "@apis/services/user_service";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faArrowRight, faRightLong } from "@fortawesome/free-solid-svg-icons";
import { useNavigate } from "react-router-dom";

const Mentor = () => {
  const [mentorList, setMentorList] = useState<any>();
  const [loading, setLoading] = useState<boolean>(true);

  const navigate = useNavigate();
  
  useEffect(() => {
    setLoading(true);

    const fetchData = async () => {
      const data = await getMentor();
      
      setMentorList(data.mentors);
      console.log(data);
      setLoading(false);
    }
    
    fetchData();
  }, [])

  return (
    <>
      <div className='relative w-full p-3'>
        <div className='w-full p-3 open-sans-600 text-2xl flex items-center flex-row'>

          <h1 className='ml-3'>Find Your Mentor</h1>
          <FontAwesomeIcon icon={faArrowRight} className="z-[99] text-black ml-3 mt-1" fade/>
        </div>
      
        <div className="w-full p-4 flex flex-row overflow-x-auto mt-3 no-scrollbar scroll-smooth">
          <div className="relative flex flex-row gap-6 sm:gap-8">
            {
              loading ? (
                <>
                  <Card className="w-[200px] space-y-5 p-4 shadow-none border-2" radius="lg">
                    <div className="space-y-3">
                      <Skeleton className="w-3/5 rounded-lg">
                        <div className="h-3 w-4/5 rounded-lg bg-default-200"></div>
                      </Skeleton>
                      <Skeleton className="w-4/5 rounded-lg">
                        <div className="h-3 w-3/5 rounded-lg bg-default-200"></div>
                      </Skeleton>
                      <Skeleton className="w-3/5 rounded-lg">
                        <div className="h-3 w-3/5 rounded-lg bg-default-200"></div>
                      </Skeleton>
                      <Skeleton className="rounded-lg">
                        <div className="h-24 rounded-lg bg-default-300"></div>
                      </Skeleton>
                    </div>
                  </Card>
                </>
              ) : (
                <>
                  {
                    mentorList.map((mentor: any) => (
                      <Card 
                        key={mentor?.userid} 
                        className="py-4 w-56 hover:scale-110 cursor-pointer border-2 shadow-sm"
                      >
                        <CardHeader className="pb-0 pt-2 px-4 flex-col items-start flex-wrap">
                          <p className="text-tiny uppercase open-sans-600">{mentor?.account_detail?.binusian}</p>
                          <div className="max-w-[240px] line-clamp-1 text-default-500 text-ellipsis overflow-hidden">
                            <small>{mentor?.mentor_detail?.course_list}</small>
                          </div>
                          <h4 className="open-sans-600 text-large">{mentor?.account_detail?.name}</h4>
                        </CardHeader>
                        <CardBody className="overflow-visible py-2 flex items-center">
                          <Image
                            alt="Card background"
                            className="object-cover rounded-xl max-h-36"
                            src={mentor?.account_detail?.profile_picture}
                            width={270}
                            onClick={() => {
                              navigate(`/mentoring/${mentor?.userid}`);
                            }}
                          />
                        </CardBody>
                      </Card>
                    ))
                  }
                </>
              )
            }
            
            
          </div>
        </div>
        <div className="absolute top-16 right-1 bg-white-accent-1 blur-md w-10 h-80 z-[99]" />
      </div> 
    </>
  )
}

export default Mentor;
