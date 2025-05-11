import { useState, useEffect } from 'react';
import './App.css';
import Hello from './components/Hello';

function App() {
  const [apiData, setApiData] = useState(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  useEffect(() => {
    async function fetchData() {
      try {
        const response = await fetch('/api/hello');
        
        if (!response.ok) {
          throw new Error(`HTTP error! status: ${response.status}`);
        }
        
        const data = await response.json();
        setApiData(data);
        setLoading(false);
      } catch (e) {
        console.error('Error fetching API data:', e);
        setError('Failed to connect to the API. Is the backend running?');
        setLoading(false);
      }
    }

    fetchData();
  }, []);

  return (
    <div className="app">
      <h1>{{ .Project }}</h1>
      <div className="card">
        <Hello />
        
        <div className="api-section">
          <h2>API Status</h2>
          {loading ? (
            <p>Loading API data...</p>
          ) : error ? (
            <p className="error">{error}</p>
          ) : (
            <div className="api-response">
              <p>{apiData.message}</p>
              <p className="timestamp">Server time: {new Date(apiData.time).toLocaleString()}</p>
            </div>
          )}
        </div>
      </div>
    </div>
  );
}

export default App;
