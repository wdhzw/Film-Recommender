import React, {useContext} from 'react';
import {LogoutOutlined, SettingOutlined, UserOutlined, LeftOutlined} from "@ant-design/icons";
import {Avatar, Button, Dropdown} from "antd";
import {AuthContext} from "../hooks/AuthContext";
import {useLocation, useNavigate} from "react-router-dom";

const fixedAvator = {
  position: 'fixed',
  top: '10px',
  right: '40px',
  zIndex: '1000',
}

const fixedBack = {
    position: 'fixed',
    top: '30px',
    left: '40px',
    zIndex: '1000',
}

const fixedWelcome = {
    position: 'fixed',
    top: '10%',
    left: '50%',
    zIndex: '1000',
    textAlign: 'center',
    transform: 'translate(-50%, -50%)'
}

function TopBar({username, currentPage}) {
    const { logout } = useContext(AuthContext);
    const navigate = useNavigate();
    const location = useLocation();
    const items = [
        {
            key: '1',
            label: (
                <a rel="" href="/profile">
                  Profile
                </a>
            ),
            icon: <SettingOutlined/>
        },
        {
            key: '2',
            label: (
                <a target="_blank" onClick={logout}>
                  Logout
                </a>
            ),
            icon: <LogoutOutlined />,
        },
    ];

    const backToDashboard = (e) => {
        navigate('/dashboard')
    }

    return (
        <div>
            <Dropdown
                menu={{
                    items,
                }}
                placement="bottom"
                arrow
            >
                <Avatar style={fixedAvator} size={64} icon={<UserOutlined/>}/>
            </Dropdown>
            {
                location.pathname.includes('movie') || location.pathname.includes('profile') ?
                    <Button onClick={backToDashboard} style={fixedBack} type="text" icon={<LeftOutlined/>}>Back</Button>
                    : null
            }
            <h1 style={fixedWelcome}>
                Welcome {username}! <br/>
                {currentPage}
            </h1>
        </div>

    )
}

export default TopBar;