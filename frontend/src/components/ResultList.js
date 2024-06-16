import React from 'react';


function ResultList({ results }) {
  return (
    <ul>
      {results.map(result => (
          <li key={result.URL}>
            {result.Info}
          </li>
      ))}
    </ul>
  );
}

export default ResultList