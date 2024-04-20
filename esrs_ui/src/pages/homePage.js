import React, {useContext, useEffect} from 'react';
import {Link, useNavigate} from 'react-router-dom';
import {AuthContext} from "../hooks/AuthContext";

const HomePage = () => {
    const { logout } = useContext(AuthContext);
    const navigate = useNavigate();

    useEffect(() => {
        navigate("/login")
    }, []);

    return (
    <div></div>
    );
};

export default HomePage;
