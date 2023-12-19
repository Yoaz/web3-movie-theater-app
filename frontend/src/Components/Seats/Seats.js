// Seats.js
import React, { useState, useEffect } from "react";
import "./Seats.scss";
import PropTypes from "prop-types";

const Seats = ({ seats }) => {
  const [seatData, setSeatData] = useState([]);

  useEffect(() => {
    setSeatData(seats);
  }, [seats]);

  const computeClassName = (isReserved) =>
    isReserved ? "reserved" : "available";

  const handleSeatClick = (rowIndex, seatIndex) => {
    // TODO: Update the reservation status for the selected seat
    const updatedSeats = seatData.map((row, currentRowIndex) =>
      currentRowIndex === rowIndex
        ? row.map((isReserved, currentSeatIndex) =>
            currentSeatIndex === seatIndex ? !isReserved : isReserved
          )
        : row
    );
    setSeatData(updatedSeats);

    // TODO: Send a request to the server to persist the reservation status
  };

  if (!seatData.length) {
    // Loading spinner while waiting for fetch
    return <div className="loading-spinner"></div>;
  }

  return (
    <div className="seats-container">
      {seatData.map((row, rowIndex) =>
        row.map((isReserved, seatIndex) => (
          <div
            key={`${rowIndex}-${seatIndex}`}
            className={`seat ${computeClassName(isReserved)}`}
            onClick={() => handleSeatClick(rowIndex, seatIndex)}
          >
            {rowIndex + 1} - {seatIndex + 1}
          </div>
        ))
      )}
    </div>
  );
};

// Validation
Seats.propTypes = {
  seats: PropTypes.array.isRequired,
};

export default Seats;
