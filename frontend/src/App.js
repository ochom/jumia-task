import Axios from "axios";
import { useEffect, useState } from "react";
import DataTable from "react-data-table-component";
import "./App.css";

const API_ROUTE = process.env.REACT_APP_API_ROUTE;

const columns = [
  { name: "Country", selector: (row) => row.Country },
  { name: "State", selector: (row) => row.State, width: "80px" },
  { name: "Country code", selector: (row) => row.Code },
  { name: "Phone", selector: (row) => row.Phone },
];

function App() {
  const [country, setCounty] = useState("all");
  const [state, setState] = useState("all");
  const [countries, setCountries] = useState([]);
  const [rows, setRows] = useState([]);

  useEffect(() => {
    // get countries
    Axios({
      url: `${API_ROUTE}api/countries`,
      method: "GET",
      headers: {
        "Content-type": "application/json",
      },
    })
      .then((res) => {
        setCountries(res.data);
      })
      .catch((err) => {
        console.log(err);
      });

    // get  numbers
    Axios({
      url: `${API_ROUTE}api/numbers/${country}/${state}`,
      method: "GET",
      headers: {
        "Content-type": "application/json",
      },
    })
      .then((res) => {
        setRows(res.data);
      })
      .catch((err) => {
        console.log(err);
      });
  }, [country, state]);

  return (
    <div className="container">
      <div className="body">
        <h2>Phone Numbers</h2>
        <div className="controls">
          <div>
            <label>Country</label>
            <select value={country} onChange={(e) => setCounty(e.target.value)}>
              <option value="all">All</option>
              {countries.map((c) => (
                <option value={c.Code} key={c.Code}>
                  {c.Name}
                </option>
              ))}
            </select>
          </div>

          <div>
            <label>Numbers</label>
            <select value={state} onChange={(e) => setState(e.target.value)}>
              <option value="all">ALL</option>
              <option value="OK">OK</option>
              <option value="NOK">NOK</option>
            </select>
          </div>
        </div>
        <DataTable
          columns={columns}
          pagination
          data={rows.map(({ Country, Phone, State, Code }) => ({
            Country,
            Phone,
            State,
            Code,
          }))}
        />
      </div>
    </div>
  );
}

export default App;
