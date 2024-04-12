import React, { useContext } from 'react';
import { Navigate } from 'react-router-dom';
import { AuthContext } from '../hooks/AuthContext';

const PrivateRoute = ({ children }) => {
  const { authData } = useContext(AuthContext);
  return authData && authData.username ? children : <Navigate to="/login" />;
};

export default PrivateRoute;
