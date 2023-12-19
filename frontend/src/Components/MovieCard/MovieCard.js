// MovieCard.js
import React from "react";
import PropTypes from "prop-types";
import ClockIcon from "./ClockIcon"; // Import your SVG clock icon
import "./MovieCard.scss";

const MovieCard = ({ name, description, duration, date }) => {
  return (
    <div className="movie-card">
      <div className="movie-card-header">
        <h2 className="movie-title">{name}</h2>
        <p className="movie-date">{date}</p>
      </div>
      <p className="movie-description">{description}</p>
      <div className="movie-duration">
        <ClockIcon />
        <span>{duration}</span>
      </div>
    </div>
  );
};

// Validation
MovieCard.propTypes = {
  name: PropTypes.string.isRequired,
  description: PropTypes.string.isRequired,
  duration: PropTypes.string.isRequired,
  date: PropTypes.string.isRequired,
};

export default MovieCard;
