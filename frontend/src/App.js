import React from "react"
import LoginPage from "./pages/Login"
import HomePage from "./pages/Home"
import RegisterPage from "./pages/Register"
import {
  BrowserRouter as Router,
  Switch,
  Route,
} from "react-router-dom";

function App() {
  return (
    <Router>
      <Switch>
        {/* root route */}
        <Route path="/" exact>
          <HomePage />
        </Route>
        <Route path="/register">
          <RegisterPage />
        </Route>
        <Route path="/login">
          <LoginPage />
        </Route>
      </Switch>
    </Router>
  );
}

export default App;
