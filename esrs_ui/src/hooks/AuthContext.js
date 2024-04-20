import React, { createContext, useState, useEffect } from 'react';

export const AuthContext = createContext(null);

export const AuthProvider = ({ children }) => {
  const [authData, setAuthData] = useState(JSON.parse(localStorage.getItem('authData')) || {});

  const handleSetAuthData = (data) => {
    localStorage.setItem('authData', JSON.stringify(data));
    setAuthData(data);
  };

  const handleLogout = () => {
    localStorage.removeItem('authData');
    setAuthData({});
  };

  useEffect(() => {
    const data = localStorage.getItem('authData');
    if (data) {
      setAuthData(JSON.parse(data));
    }
  }, []);

 return (
    <AuthContext.Provider value={{ authData, setAuthData: handleSetAuthData, logout: handleLogout }}>
      {children}
    </AuthContext.Provider>
  );
};
