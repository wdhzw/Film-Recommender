import './App.css';
import LoginPage from "./pages/loginPage";
import {BrowserRouter as Router, Routes, Route} from "react-router-dom";
import HomePage from "./pages/homePage";
import PrivateRoute from "./pages/privateRoute";
import DashboardPage from "./pages/dashboardPage";
import DetailsPage from "./pages/detailsPage";

function App() {

  return (
    <Router>
        <div className="App">
            <Routes>
                <Route path="/" element={<HomePage />} />
                <Route path="/login" element={<LoginPage />} />
                <Route
                    path="/dashboard"
                    element={
                        <PrivateRoute>
                            <DashboardPage />
                        </PrivateRoute>
                    }
                  />
                <Route
                    path='/movie/:movieID'
                    element={
                        <PrivateRoute>
                            <DetailsPage />
                        </PrivateRoute>
                    }
                />
            </Routes>
        </div>
    </Router>
  );
}

export default App;
