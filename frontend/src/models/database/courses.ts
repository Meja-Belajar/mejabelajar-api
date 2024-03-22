interface Courses {
  id: string;
  name: string;
  detail: string;
  isAvailable: boolean;
  createdBy: string;
  updatedBy: string;
  createdAt: Date;
  updatedAt: Date;
  mentors: Mentors[];
  booking: Bookings;
}
