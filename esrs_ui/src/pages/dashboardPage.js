import { UndoOutlined, EditOutlined } from '@ant-design/icons';
import React, {useContext, useEffect, useState} from 'react';
import { useNavigate } from 'react-router-dom';
import {Divider, Flex, Button} from 'antd';
import MovieList from "../components/movieList";
import SearchBar from "../components/searchBar";
import TopBar from "../components/topBar";
import {AuthContext} from "../hooks/AuthContext";
import {getRecommendations} from "../api";
import {next} from "lodash/seq";

const DashboardPage = () => {
    const { authData } = useContext(AuthContext);
    const [movieStorage, setMovieStorage] = useState([])
    const [movies, setMovies] = useState([])
    const [moreClickTime, setMoreClickTime] = useState(0)
    const navigate = useNavigate();

    useEffect(() => {
        fetchRecommendationMovieds()
    }, []);

    const fetchRecommendationMovieds = async () => {
        let pageID = Math.floor(movieStorage.length / 20) + 1
        let getRecommendationsRes = await getRecommendations(authData.userInfo.email, pageID)
        console.log(getRecommendationsRes)
        if (getRecommendationsRes !== null) {
            let tmp = movieStorage
            for (let i = 0; i < getRecommendationsRes.length; i++) {
                movieStorage.push(getRecommendationsRes[i])
            }
            setMovieStorage(tmp)
            setMovies(movieStorage.slice(moreClickTime * 5, (moreClickTime+1) * 5))
        }
    }

    const handleMore = () => {
        if (moreClickTime * 5 < movieStorage.length) {
            setMoreClickTime(moreClickTime + 1)
            setMovies(movieStorage.slice(moreClickTime * 5, (moreClickTime+1) * 5))
        } else {
            setMoreClickTime(moreClickTime + 1)
            fetchRecommendationMovieds()
        }
    }

    return (
      <Flex gap="middle" vertical>
          <TopBar username={authData.userInfo.user_name} currentPage={"Recommendation Center"}/>
          <SearchBar />
          <Divider/>
          <h3>Based on your prefer genres:</h3>
          <MovieList movies={movies} />
          <Flex gap="middle" align={"center"} justify={"center"}>
              <Button type="primary" onClick={handleMore} icon={<UndoOutlined />}>
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
