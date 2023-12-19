import { useState, useEffect, useRef } from "react";
import { getAllSchedules } from "../Utils/Fetch Data/fetchData";
import { Link } from "react-router-dom";
import MovieCard from "../MovieCard/MovieCard";
import "./Home.scss";
import { formatDuration } from "../Utils/Helpers";

const Home = () => {
  const [schedules, setSchedules] = useState(null);
  const [sortedByDate, setSortedByDate] = useState(true);
  const scheduleRowRef = useRef(null);

  // Sort schedule by date
  const sortSchedulesByDate = (schedules, ascending) => {
    const sortedSchedules = [...schedules];

    sortedSchedules.sort((a, b) => {
      const dateA = new Date(a.date);
      const dateB = new Date(b.date);

      return ascending ? dateA - dateB : dateB - dateA;
    });

    return sortedSchedules;
  };

  useEffect(() => {
    const fetchData = async () => {
      try {
        const data = await getAllSchedules();
        const sortedSchedules = sortSchedulesByDate(
          data.data.data,
          sortedByDate
        );
        setSchedules(sortedSchedules);
      } catch (error) {
        console.error("Error fetching schedules:", error);
      }
    };

    fetchData();
  }, [sortedByDate]);

  const toggleSortOrder = () => {
    setSortedByDate((prevSortOrder) => !prevSortOrder);
  };

  if (!schedules) {
    return <div className="loading-spinner"></div>;
  }

  return (
    <div className="schedules-container">
      {schedules.map((schedule) => (
        <div key={schedule._id} className="schedule-row">
          <h2 onClick={toggleSortOrder} className="date-title">
            {new Date(schedule.date).toLocaleDateString()}
          </h2>
          <div className="scroll-container" ref={scheduleRowRef}>
            <div className="movies-container">
              {schedule.movies.map((movie) => (
                <Link key={movie._id} to={`/movie/${movie._id}`}>
                  <MovieCard
                    id={movie._id}
                    name={movie.name}
                    description={movie.description}
                    duration={formatDuration(movie.duration)}
                  />
                </Link>
              ))}
            </div>
          </div>
        </div>
      ))}
    </div>
  );
};

export default Home;
