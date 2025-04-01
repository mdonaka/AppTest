import ListFlame from '../components/ListFlame';
import {useState, useEffect} from 'react';

async function getTable() {
  const response = await fetch('http://localhost:8080/data');
  const data = await response.json();
  return data;
}

function ShowTable() {
  const [data, setData] = useState(null);

  useEffect(() => {
    async function fetchData() {
      const fetchedData = await getTable();
      setData(fetchedData);
    }
    fetchData();
  }, []);

  if (data === null) {
    return <h1>Loading...</h1>;
  }
  return (
    <div>
      <ListFlame data={data} />
    </div>
  )

}

export default function List() {
  return (
    <div>
      <ShowTable />
    </div>
  )
}
