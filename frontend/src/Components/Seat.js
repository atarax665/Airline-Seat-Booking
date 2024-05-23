// Seat.js
import React from "react";

const Seat = ({ row, number, isSelected, onClick, booked }) => {
  return (
    <div
      className={`w-12 h-12 border-2 flex items-center justify-center cursor-pointer m-1
        ${isSelected ? "bg-green-500 text-white" : "bg-gray-200"}
      `}
      onClick={() => (booked ? null : onClick(row, number))}
    >
      {`${row}${number}`}
    </div>
  );
};

export default Seat;
