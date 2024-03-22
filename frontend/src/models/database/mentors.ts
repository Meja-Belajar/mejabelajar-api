interface Mentors {
  id: string;
  revenue: number;
  rating: number;
  totalTeachingHours: number;
  teachingFrequency: number;
  isActive: boolean;
  createdBy: string;
  updatedBy: string;
  createdAt: Date;
  updatedAt: Date;
  courses: Courses[];
  mentorReviews: MentorReviews[];
  bookings: Bookings[];
  userId: string;
}
