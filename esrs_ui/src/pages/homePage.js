import React, {useContext} from 'react';
import { Link } from 'react-router-dom';
import {AuthContext} from "../hooks/AuthContext";

const HomePage = () => {
    const { logout } = useContext(AuthContext);
    return (
    <div>
      <h1>Home Page</h1>
      <Link to="/login">Login</Link> | <Link to="/dashboard">Dashboard</Link>
        <button onClick={logout}>Logout</button>
    </div>
    );
};

export default HomePage;
