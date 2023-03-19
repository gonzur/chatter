const formatToTwelveHourDate = (date: Date) => {
  const meridian = date.getHours() >= 12 ? "pm" : "am";

  let minute = date.getMinutes().toString();
  if (minute.length < 2) minute = `0${minute}`;

  let hour =
    date.getHours() >= 12
      ? (date.getHours() - 12).toString()
      : date.getHours().toString();
  if (hour === "0") hour = "12";

  return `${hour}:${minute} ${meridian}`;
};

export default formatToTwelveHourDate;
