// Movie.js
import React, { useState, useEffect } from "react";
import MovieCard from "../MovieCard/MovieCard";
import "./Movie.scss";
import { useParams } from "react-router-dom";
import { getMovieByID } from "../Utils/Fetch Data/fetchData";
import Seats from "../Seats/Seats";
import { formatDuration } from "../Utils/Helpers";

const Movie = () => {
  const { id } = useParams();
  const [movie, setMovie] = useState(null);

  useEffect(() => {
    const fetchData = async () => {
      try {
        const data = await getMovieByID(id);
        setMovie(data.data.data);
      } catch (error) {
        console.error("Error fetching movie:", error);
      }
    };

    fetchData();
  }, [id]);

  if (!movie) {
    // Loading spinner while waiting for fetch
    return <div className="loading-spinner"></div>;
  }

  return (
    <div className="movie-container">
      <MovieCard
        id={movie._id}
        name={movie.name}
        description={movie.description}
        duration={formatDuration(movie.duration)}
        date={movie.date}
      />
      <Seats seats={movie.available_seats} />
    </div>
  );
};

export default Movie;
