import React, { useState } from 'react';
import axios from 'axios';
import './App.css';

function App() {
    const [urls, setUrls] = useState('');
    const [results, setResults] = useState([]);

    const handleSubmit = async (event) => {
        event.preventDefault();
        try {
            await axios.post('http://localhost:8080/api/start', {
                urls: urls.split(',').map(url => url.trim())
            });
            const response = await axios.get('http://localhost:8080/api/results');
            setResults(response.data);
        } catch (error) {
            console.error('Error starting the ping process', error);
        }
    };

    return (
        <div className="App">
            <header className="App-header">
                <h1>URL Pinger</h1>
                <form onSubmit={handleSubmit}>
                    <input
                        type="text"
                        value={urls}
                        onChange={(e) => setUrls(e.target.value)}
                        placeholder="Enter URLs separated by commas"
                    />
                    <button type="submit">Start Ping</button>
                </form>
                <h2>Results</h2>
                <ul>
                    {results.map((result, index) => (
                        <li key={index}>
                            URL: {result.url} - Status: {result.status} - Response Time: {result.response_time}
                        </li>
                    ))}
                </ul>
            </header>
        </div>
    );
}

export default App;
