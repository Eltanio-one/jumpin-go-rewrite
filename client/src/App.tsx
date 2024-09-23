import './App.css';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import HomePage from './HomePage';
import Login from './Login';
import { PrivateRoute } from './components/PrivateRoute';

function App() {
  return (
    <div className="App">
      <header className="App-header">
        <Router>
          <Routes>
            <Route path='/' element={<Login />}></Route>
            <Route path="/home" element={<PrivateRoute><HomePage /></PrivateRoute>}></Route>
          </Routes>
        </Router>
      </header>
    </div>
  );
}

export default App;
