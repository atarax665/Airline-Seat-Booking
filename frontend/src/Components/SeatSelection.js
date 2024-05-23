import React, { useState, useEffect } from "react";
import Seat from "./Seat";
import axios from "axios";

const SeatSelection = () => {
  const rows = [
    1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20,
  ];
  const seatLabels = ["A", "B", "C", "", "D", "E", "F"]; // Include an empty string for the aisle

  const [selectedSeats, setSelectedSeats] = useState([]);
  const [bookedSeats, setBookedSeats] = useState([]);

  useEffect(() => {
    axios
      .get("http://localhost:8080/v1/get?flightId=1")
      .then((response) => {
        const filledSeats = response.data.bookedSeats;
        const seatIds = filledSeats.map((seat) => seat.name.replace("-", ""));
        setBookedSeats(seatIds);
      })
      .catch((error) => console.error("Error:", error));
  }, []);

  const handleSeatClick = (row, label) => {
    if (label === "") return; // Don't do anything if the aisle (empty space) is clicked
    const seatId = `${row}${label}`;
    if (selectedSeats.includes(seatId)) {
      setSelectedSeats(selectedSeats.filter((seat) => seat !== seatId));
    } else {
      setSelectedSeats([...selectedSeats, seatId]);
    }
  };

  const renderSeatRows = () => {
    return rows.map((row) => (
      <div key={row} className="flex justify-center mb-2">
        {seatLabels.map((label, index) =>
          label === "" ? (
            <div key={index} style={{ width: "24px", height: "12px" }} /> // Visual spacer for the aisle
          ) : (
            <Seat
              key={`${row}${label}`}
              row={row}
              number={label}
              isSelected={selectedSeats.includes(`${row}${label}`)}
              onClick={() => handleSeatClick(row, label)}
              booked={bookedSeats.includes(`${row}${label}`)}
            />
          )
        )}
      </div>
    ));
  };

  return (
    <div className="bg-gray-100 mt-0">
      <h2 className="text-2xl font-bold mb-4 text-center">Select Your Seat</h2>
      <div className="mt-0">{renderSeatRows()}</div>
    </div>
  );
};

export default SeatSelection;
