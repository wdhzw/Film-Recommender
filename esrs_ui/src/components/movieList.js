import React, {useContext, useEffect, useState} from 'react';
import { List, Card } from 'antd';
import './movieList.css';
import {useNavigate} from "react-router-dom";
import {AuthContext} from "../hooks/AuthContext";
import {getRecommendations} from "../api";

// const data = [
//   {
//     title: 'To Mom (and Dad) with love',
//     url: 'https://image.tmdb.org/t/p/original/sChfCU3PDV3N6nYessVPWeWkUBc.jpg'
//
//   },
//   {
//     title: 'The book of love',
//     url: 'https://image.tmdb.org/t/p/original/hwP0GEP0zy8ar965Xaht19SmMd3.jpg'
//   },
//   {
//     title: 'Endless Love',
//     url: 'https://image.tmdb.org/t/p/original/z7FZP6uivgVc4t0mnmia0B8YygW.jpg'
//   },
//     {
//     title: 'Love Strange Love',
//     url: 'https://image.tmdb.org/t/p/original/9CNnxpI6H8ynyOlACRc25vqgJBY.jpg'
//   },
//     {
//     title: 'Sorry If I Call You Love',
//     url: 'https://image.tmdb.org/t/p/original/pnSXPKQPjVi87YEeRYlbg5aUaGs.jpg'
//   },
// ];

const MovieList = ({movies}) => {
    const navigate = useNavigate();

      return (
        <List
          grid={{ gutter: 16, columns: 5 }}
          dataSource={movies}
          renderItem={movie => (
            <List.Item>
              <Card
                hoverable
                onClick={() => {navigate('/movie/'+movie.movie_id)}}
                cover={
                <img
                    style={{
                        height: '300px',
                        width: '250px'
                    }}
                    alt={movie.title}
                    src={movie.poster_uri}
                />}
              >
                {movie.title.slice(0, 20)}
              </Card>
            </List.Item>
          )}
        />
      );
};

export default MovieList;
