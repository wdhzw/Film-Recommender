import { UserOutlined, SettingOutlined, LogoutOutlined, UndoOutlined } from '@ant-design/icons';
import React, {useContext} from 'react';
import {Avatar, Divider, Flex, Dropdown, Button} from 'antd';
import {AuthContext} from "../hooks/AuthContext";
import MovieList from "../components/movieList";
import SearchBar from "../components/searchBar";

const fixedAvator = {
  position: 'fixed',
  top: '10px',
  right: '40px',
  zIndex: '1000'
}

const DashboardPage = () => {
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
      <Flex gap="middle" vertical>
          <Dropdown
              menu={{
                  items,
              }}
              placement="bottom"
              arrow
          >
              <Avatar style={fixedAvator} size={64} icon={<UserOutlined />} />
          </Dropdown>

          <h1>Recommendation Center</h1>
          <SearchBar />
          <Divider/>
          <MovieList/>
          <Button style={{width: '30%', margin: '0 auto'}} type="primary" icon={<UndoOutlined />}>
            More
          </Button>
      </Flex>
      );
};

export default DashboardPage;
