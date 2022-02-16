import Axios from "axios";
import { useEffect, useState } from "react";
import DataTable from "react-data-table-component";
import "./App.css";

const columns = [
  { name: "Country", selector: (row) => row.Country },
  { name: "Code", selector: (row) => row.Code },
  { name: "State", selector: (row) => row.State },
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
      url: `http://localhost:8000/api/countries`,
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
      url: `http://localhost:8000/api/numbers/${country}/${state}`,
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
