import { UndoOutlined } from '@ant-design/icons';
import React, {useContext} from 'react';
import {Avatar, Divider, Flex, Dropdown, Button} from 'antd';
import MovieList from "../components/movieList";
import SearchBar from "../components/searchBar";
import TopBar from "../components/topBar";

const DashboardPage = () => {

    return (
      <Flex gap="middle" vertical>
          <TopBar/>
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
