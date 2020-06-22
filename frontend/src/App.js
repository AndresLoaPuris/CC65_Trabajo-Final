

import React from "react";
import {
  BrowserRouter as Router,
  Switch,
  Route,
  Link,
  useRouteMatch,
  useParams
} from "react-router-dom";
import KMeans from './KMeans'
import KMeansDetalle from './KMeansDetalle'
import Home from './Home'
import KNN from './Knn'

export default function App() {
  return (
    <Router>
      <div>


        <nav style={{ marginTop: 10 }} class="navbar navbar-expand-lg">

        <div class="collapse navbar-collapse" >
        <ul class="navbar-nav">
          <li style={{ marginLeft: 10 }} class="nav-item active">
            <Link to="/" style={{ color: 'black' }}>Home</Link>
          </li>
          <li style={{ marginLeft: 40 }}  class="nav-item active">
            <Link to="/knn" style={{ color: 'black' }}> Knn </Link>
          </li>
          <li style={{ marginLeft: 40 }} class="nav-item active">
            <Link to="/kmeans" style={{ color: 'black' }}>K Means - DataSet</Link>
          </li>
          <li style={{ marginLeft: 40 }}  class="nav-item active">
            <Link to="/kmeansDetalle" style={{ color: 'black' }}>K Means - Grupos</Link>
          </li>

        </ul>
            </div>
        </nav>


        <Switch>
          <Route path="/kmeansDetalle">
          <div style={{ marginTop: 30 }} class="container-fluid">
          <KMeansDetalle />
          </div>
            
          </Route>

          <Route path="/knn">
          <div style={{ marginTop: 30 }} class="container-fluid">
          <KNN />
          </div>
            
          </Route>

          <Route path="/kmeans">

          <div style={{ marginTop: 30 }} class="container-fluid">
          <KMeans />
          </div>
          </Route>

          <Route path="/">
          <div style={{ marginTop: 30 }} class="container-fluid">
          <Home />
          </div>
          </Route>

        </Switch>
      </div>
    </Router>
  );

}