import React from 'react';
import './App.css';
import Routes from './Routes';
import Nav from './components/nav';

function App() {
  return (
    <div className="App">
    	<Nav />
    	<Routes />
    </div>
  );
}

export default App;