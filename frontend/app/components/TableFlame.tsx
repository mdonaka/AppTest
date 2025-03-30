import React from "react";
import {useTable} from "react-table";

const TableFlame = ({data}) => {

  const columns = React.useMemo(
    () => [
      {Header: "ID", accessor: "id"},
      {Header: "Name", accessor: "name"},
      {Header: "Flavor", accessor: "flavor"},
      {Header: "Family", accessor: "family"},
    ],
    []
  );

  const {getTableProps, getTableBodyProps, headerGroups, rows, prepareRow} =
    useTable({columns, data});

  return (
    <table {...getTableProps()} border="1">
      <thead>
        {headerGroups.map((headerGroup) => (
          <tr {...headerGroup.getHeaderGroupProps()}>
            {headerGroup.headers.map((column) => (
              <th {...column.getHeaderProps()}>{column.render("Header")}</th>
            ))}
          </tr>
        ))}
      </thead>
      <tbody {...getTableBodyProps()}>
        {rows.map((row) => {
          prepareRow(row);
          return (
            <tr {...row.getRowProps()}>
              {row.cells.map((cell) => (
                <td {...cell.getCellProps()}>{cell.render("Cell")}</td>
              ))}
            </tr>
          );
        })}
      </tbody>
    </table>
  )
};

export default TableFlame;
