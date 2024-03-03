import { Card, CardBody, CardHeader, Image, Skeleton } from "@nextui-org/react";
import { useEffect, useState } from "react";
import { getMentor } from "../api/services/user_service";

const Mentor = () => {
  const [mentorList, setMentorList] = useState<any>();
  const [loading, setLoading] = useState<boolean>(true);

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
      <div className='w-full p-3'>
        <div className='w-full p-3 lato-bold text-2xl'>

          <h1 className='ml-3'>Find Your Mentor</h1>
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
                      <Card key={mentor?.userid} className="py-4 w-56 hover:scale-110 cursor-pointer border-2 shadow-none">
                        <CardHeader className="pb-0 pt-2 px-4 flex-col items-start flex-wrap">
                          <p className="text-tiny uppercase font-bold">{mentor?.account_detail?.binusian}</p>
                          <div className="max-w-[240px] line-clamp-1 text-default-500 text-ellipsis overflow-hidden">
                            <small>{mentor?.mentor_detail?.course_list}</small>
                          </div>
                          <h4 className="font-bold text-large">{mentor?.account_detail?.name}</h4>
                        </CardHeader>
                        <CardBody className="overflow-visible py-2 flex items-center">
                          <Image
                            alt="Card background"
                            className="object-cover rounded-xl max-h-36"
                            src={mentor?.account_detail?.profile_picture}
                            width={270}
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
      </div> 
    </>
  )
}

export default Mentor;
