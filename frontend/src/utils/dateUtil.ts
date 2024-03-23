export const maxDateUtil = () => {
  const today = new Date();

  // Calculate the date 15 years ago
  const minDate = new Date(today.getFullYear() - 15, today.getMonth(), today.getDate());

  // Format the minDate as YYYY-MM-DD
  const minDateString = minDate.toISOString().split('T')[0];

  return minDateString
}