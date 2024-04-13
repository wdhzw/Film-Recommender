import React, {useContext} from 'react';
import {LogoutOutlined, SettingOutlined, UserOutlined, LeftOutlined} from "@ant-design/icons";
import {Avatar, Dropdown} from "antd";
import {AuthContext} from "../hooks/AuthContext";

const fixedAvator = {
  position: 'fixed',
  top: '10px',
  right: '40px',
  zIndex: '1000',
}

function TopBar() {
    const { logout } = useContext(AuthContext);
    const items = [
        {
            key: '1',
            label: (
                <a target="_blank" rel="" href="">
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

    return (
        <div>
            <Dropdown
              menu={{
                  items,
              }}
              placement="bottom"
              arrow
            >
                <Avatar style={fixedAvator} size={64} icon={<UserOutlined />} />
            </Dropdown>
            <LeftOutlined style={{

            }}/>
        </div>

    )
}

export default TopBar;