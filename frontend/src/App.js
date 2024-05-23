// App.js
import React from "react";
import ReactDOM from "react-dom/client";
import SeatSelection from "./Components/SeatSelection";
import ApproachSelection from "./Components/ApprochSelection";
import GetUser from "./Components/GetUser";

const App = () => {
  return (
    <div className="min-h-screen flex items-center justify-center bg-blue-50">
      <div>
        <h1 className="text-2xl font-bold text-center">
          Flight Seat Selection
        </h1>
        <ApproachSelection />
        <GetUser />
        <SeatSelection />
      </div>
    </div>
  );
};

// create root using createRoot
const root = ReactDOM.createRoot(document.getElementById("root"));
// passing react element inside root
root.render(<App />);
