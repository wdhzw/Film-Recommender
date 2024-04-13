import { UndoOutlined, EditOutlined } from '@ant-design/icons';
import React, {useContext} from 'react';
import { useNavigate } from 'react-router-dom';
import {Divider, Flex, Button} from 'antd';
import MovieList from "../components/movieList";
import SearchBar from "../components/searchBar";
import TopBar from "../components/topBar";
import {AuthContext} from "../hooks/AuthContext";

const DashboardPage = () => {
    const { authData } = useContext(AuthContext);
    const navigate = useNavigate();

    return (
      <Flex gap="middle" vertical>
          <TopBar username={authData.userInfo.user_name} currentPage={"Recommendation Center"}/>
          <SearchBar />
          <Divider/>
          <h3>Based on your prefer genres:</h3>
          <MovieList/>
          <Flex gap="middle" align={"center"} justify={"center"}>
              <Button type="primary" icon={<UndoOutlined />}>
                More
              </Button>
              <Button type="default" onClick={() => navigate('/profile')} icon={<EditOutlined />}>
                  Change My Prefer Genres
              </Button>
          </Flex>
      </Flex>
      );
};

export default DashboardPage;
