import React, {useState, useEffect} from 'react';
import '../styles/Table.scss';

interface DataType {
  [key: string]: any;
}

const Table = () => {
  const [data, setData] = useState<DataType[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);
  const [filters, setFilters] = useState<{[key: string]: string[]}>({});
  const [expandedKeys, setExpandedKeys] = useState<string[]>([]);
  const [isFilterExpanded, setIsFilterExpanded] = useState(false);

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
    return <div className="noData">データがありません</div>;
  }

  const columns = Object.keys(data[0]);

  const filteredData = data.filter(item => {
    return columns.every(column => {
      if (!filters[column] || filters[column].length === 0) {
        return true;
      }
      return filters[column].includes(item[column]);
    });
  });

  const handleKeyClick = (column: string) => {
    setExpandedKeys(prevExpandedKeys => {
      if (prevExpandedKeys.includes(column)) {
        return prevExpandedKeys.filter(key => key !== column);
      } else {
        return [...prevExpandedKeys, column];
      }
    });
  };

  const getFilterOptionsHeight = (column: string) => {
    const numberOfOptions = [...new Set(data.map(item => item[column]))].length;
    const optionHeight = 35;
    const baseHeight = 38;
    return expandedKeys.includes(column) ? `${baseHeight + (numberOfOptions * optionHeight)}px` : `${baseHeight}px`;
  }

  return (
    <div className="filterContainer">
      <label onClick={() => setIsFilterExpanded(!isFilterExpanded)} className="filterHeader">
        フィルタ
      </label>
      {isFilterExpanded && (
        <div className="filterOptions">
          {columns.map(column => {
            return (
              <div
                key={column}
                className="filterKeyContainer"
                style={{maxHeight: getFilterOptionsHeight(column)}}
              >
                <label onClick={() => handleKeyClick(column)} className="filterKey">
                  {column}
                </label>
                {
                  [...new Set(data.map(item => column))].length > 0 && expandedKeys.includes(column) && (
                    <div>
                      {[...new Set(data.map(item => item[column]))].map(value => (
                        <label key={value} className="filterOptionLabel">
                          <input
                            type="checkbox"
                            value={value}
                            checked={filters[column]?.includes(value) || false}
                            onChange={e => {
                              const checked = !filters[column]?.includes(value);
                              setFilters(prevFilters => {
                                const columnFilters = prevFilters[column] || [];
                                return {
                                  ...prevFilters,
                                  [column]: checked
                                    ? [...columnFilters, value]
                                    : columnFilters.filter(v => v !== value),
                                };
                              });
                            }}
                          />
                          {value}
                        </label>
                      ))}
                    </div>
                  )
                }
              </div>
            );
          })}
        </div>
      )}
      <table className="tableContainer">
        <thead>
          <tr>
            {columns.map(column => (
              <th key={column} className="tableHeader">{column}</th>
            ))}
          </tr>
        </thead>
        <tbody>
          {filteredData.map((row, index) => (
            <tr key={index} className="tableRow">
              {columns.map(column => (
                <td key={`${index}-${column}`} className="tableCell">{row[column]?.toString()}</td>
              ))}
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
};

export default Table;
