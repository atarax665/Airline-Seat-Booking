// Seat.js
import React from "react";

const Approach = ({ name, number, isSelected, onClick }) => {
  return (
    <div
      className={`w-38 h-12 border-2 border-black flex items-center justify-center cursor-pointer m-1
        ${isSelected ? "bg-green-500 text-white" : "bg-gray-200"}
      `}
      onClick={() => onClick(number)}
    >
      {name}
    </div>
  );
};

export default Approach;
