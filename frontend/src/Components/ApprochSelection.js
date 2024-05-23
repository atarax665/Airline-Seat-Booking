import React, { useState } from "react";
import Approach from "./Approach";

const ApproachSelection = () => {
  const approaches = ["1", "2", "3", "4", "5"]; // Identifiers for each approach
  const [selectedApproach, setSelectedApproach] = useState("");

  const handleApproachClick = (number) => {
    if (selectedApproach === number) {
      setSelectedApproach(null);
    } else {
      setSelectedApproach(number);
    }
  };

  const applyApproach = () => {
    alert(`Approach ${selectedApproach} applied!`);
  };

  const getContentForApproach = (number) => {
    switch (number) {
      case "1":
        return "Content for Approach 1";
      case "2":
        return "Content for Approach 2";
      case "3":
        return "Content for Approach 3";
      case "4":
        return "Content for Approach 4";
      case "5":
        return "Content for Approach 5";
      default:
        return "Select an approach";
    }
  };

  return (
    <div className="p-4 bg-gray-100">
      <div className="flex justify-center mb-4">
        {approaches.map((number) => (
          <Approach
            key={number}
            name={`Approach ${number}`}
            number={number}
            isSelected={selectedApproach === number}
            onClick={handleApproachClick}
          />
        ))}
      </div>
      <div className="mt-4 p-4 bg-white border rounded shadow">
        {getContentForApproach(selectedApproach)}
      </div>
      <button
        className="mt-4 bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded"
        onClick={applyApproach}
      >
        Apply Approach
      </button>
    </div>
  );
};

export default ApproachSelection;
