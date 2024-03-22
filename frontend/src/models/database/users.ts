interface Users {
  id: string;
  username: string;
  university?: string;
  email: string;
  password: string;
  phone?: string;
  description?: string;
  profilePicture?: string;
  bod: Date;
  isActive: boolean;
  createdBy: string;
  updatedBy: string;
  createdAt: Date;
  updatedAt: Date;
  bookings?: Bookings[];
  mentor?: Mentors;
}