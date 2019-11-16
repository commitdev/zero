import React from 'react';
import './App.css';
import Divider from '@material-ui/core/Divider';


import Providers from './Providers.js';
import Frontend from './Frontend.js';
import Services from './Services.js';

function App() {
  return (
    <div className="App">
      <h1>Commit0 - Create a new project</h1>
      <Providers />
      <Frontend />
      <Services />
      <Divider variant="bottom"/>
    </div>
  );
}

export default App;
