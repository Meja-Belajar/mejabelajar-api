// sub interface 
interface Social {
  linkedin: string;
  whatsapp: string;
}

interface AccountDetail {
  name: string;
  profile_picture: string;
  bod: string; 
  social: Social;
  description: string;
}

interface Course {
  name: string;
  price_per_hours: number;
}

interface TeachHistory {
  name: string;
  total_hours: string;
}

interface MentorDetail {
  course: Course[];
  available: boolean;
  teach_history: TeachHistory[]; 
  review: string; 
  total_teaching_hours: string; 
  teaching_frequent: string;
}

// export 
export interface User {
  userid: string;
  account_type: string;
  email: string;
  account_detail: AccountDetail;
  mentor_detail: MentorDetail;
}

export interface Live {
  userid: string;
  thumbnail: string;
  startTime: string;
  expired: boolean;
  name: string;
  title: string;
  total_watchers: number;
}

export interface RegisterUser {
  name: string;
  password: string;
  email: string;
}

export interface UserLogin {
  status: number;
  user: User;
}