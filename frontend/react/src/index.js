import React from 'react';
import ReactDOM from 'react-dom';
import {BrowserRouter as Router,Route} from 'react-router-dom'
import {Button} from 'reactstrap'
import 'bootstrap/dist/css/bootstrap.min.css'
import './index.css';
import UserAttachmentPage from './UserAttachmentPage';
import * as serviceWorker from './serviceWorker';

ReactDOM.render(
  <React.StrictMode>
    <Router>
{/*       <Route exact={true} path="/" render={() => (
        <div>
          <h1> Fileshare App</h1>
        </div>
      )}/> */}
      <Route path="/" component={UserAttachmentPage}/>
      {/* <Route render={ ()=> (
        <h1> 404 - not found</h1>
      )} /> */}
    </Router>
    
  </React.StrictMode>,
  document.getElementById('root')
);

// If you want your app to work offline and load faster, you can change
// unregister() to register() below. Note this comes with some pitfalls.
// Learn more about service workers: https://bit.ly/CRA-PWA
serviceWorker.unregister();
