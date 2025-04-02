import React, {useState, useEffect} from 'react';
import '../styles/Table.scss';

interface DataType {
  [key: string]: any;
}

const Table = () => {
  const [data, setData] = useState<DataType[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    const fetchData = async () => {
      try {
        const response = await fetch('http://localhost:8080/data');
        if (!response.ok) {
          throw new Error(`HTTP error! status: ${response.status}`);
        }
        const json = await response.json();
        setData(json);
      } catch (e: any) {
        setError(e.message);
      } finally {
        setLoading(false);
      }
    };

    fetchData();
  }, []);

  if (loading) {
    return <div className="loading">Loading...</div>;
  }

  if (error) {
    return <div className="error">Error: {error}</div>;
  }

  if (!data || data.length === 0) {
    return <div className="noData">No data available.</div>;
  }

  const columns = Object.keys(data[0]);

  return (
    <table className="tableContainer">
      <thead>
        <tr>
          {columns.map(column => (
            <th key={column} className="tableHeader">{column}</th>
          ))}
        </tr>
      </thead>
      <tbody>
        {data.map((row, index) => (
          <tr key={index} className="tableRow">
            {columns.map(column => (
              <td key={`${index}-${column}`} className="tableCell">{row[column]?.toString()}</td>
            ))}
          </tr>
        ))}
      </tbody>
    </table>
  );
};

export default Table;
