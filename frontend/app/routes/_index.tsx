import {useState, useEffect} from 'react';
import type {MetaFunction} from "@remix-run/node";
import Table from '../components/Table';

export const meta: MetaFunction = () => {
  return [
    {title: "New Remix App"},
    {name: "description", content: "Welcome to Remix!"},
  ];
};


export async function get() {
  const response = await fetch('http://localhost:8080/data');
  const data = await response.json();
  return data;
}

export default function Index() {
  const [data, setData] = useState(null);

  useEffect(() => {
    async function fetchData() {
      const fetchedData = await get();
      setData(fetchedData);
    }
    fetchData();
  }, []);

  if (data === null) {
    return <h1>Loading...</h1>;
  }

  return (
    <Table data={data} />
  )
}
