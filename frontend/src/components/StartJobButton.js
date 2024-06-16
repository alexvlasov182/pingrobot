import React from "react";

function StartJobButton({ startJob }) {
  return (
    <button onClick={() => startJob("https://google.com")}>
      Start Job
    </button>
  );
}

export default StartJobButton