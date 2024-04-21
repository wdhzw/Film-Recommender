import React, {useEffect, useLayoutEffect, useState} from 'react';
import { useParams } from 'react-router-dom';
import {Flex, Descriptions, Rate, Card, Image} from 'antd';
import TopBar from "../components/topBar";
import {getMovieDetail} from "../api";
const { Meta } = Card;
const desc = ['terrible', 'bad', 'normal', 'good', 'wonderful'];

function DetailsPage() {
    const { movieID } = useParams();
    const [ movie, setMovie ] = useState({
        movie_id: null,
        director: '',
        keywords: [],
        actors: [],
        genres: [],
        language: [],
        overview: [],
        popularity: 0,
        rate: 0,
        title: '',
        release_date: '',
        writers: [],
        poster_uri: '',
        production_countries: [],
    })
    const [imgKey, setImgKey] = useState(0)

    useLayoutEffect(() => {
        fetchDetail()
    }, []);

    const fetchDetail = async() => {
        let movieDetailRes = await getMovieDetail(movieID)
        console.log(movieDetailRes)
        if (movieDetailRes !== null &&movieDetailRes.status_code === 0) {

            setMovie({
                movie_id: movieDetailRes.content.movie_id,
                director: movieDetailRes.content.director,
                keywords: movieDetailRes.content.keywords,
                actors: movieDetailRes.content.actors,
                genres: movieDetailRes.content.genres,
                language: movieDetailRes.content.language,
                overview: movieDetailRes.content.overview,
                popularity: movieDetailRes.content.popularity,
                rate:  movieDetailRes.content.rate,
                title:  movieDetailRes.content.title,
                release_date:  movieDetailRes.content.release_date,
                writers:  movieDetailRes.content.writers,
                poster_url:  movieDetailRes.content.poster_uri,
                production_countries: movieDetailRes.content.production_countries
            })

            const imgElement = document.getElementById('poster');
            console.log(imgElement)
            imgElement.src = movieDetailRes.content.poster_uri
            console.log(imgElement)
        }
    }

    const items = [
      {
        key: '1',
        label: 'Release Date',
        children: movie.release_date.slice(0, 10),
      },
      {
        key: '2',
        label: 'Production Countries',
        children: movie.production_countries,
      },
      {
        key: '3',
        label: 'Language',
        children: movie.language,
      },
      {
        key: '4',
        label: 'Director',
        children: movie.director,
      },
      {
        key: '5',
        label: 'stars',
        children: movie.actors.map(actor => actor.name).join(', '),
      },
        {
            key: '6',
            label: 'genre',
            children: movie.genres.join(', ')
        },
        {
            key: '7',
            label: 'Overview',
            children: movie.overview
        }
    ];

    return (
        <Flex gap="middle">
            <TopBar/>
            <Card
                style={{width: 300}}
                cover={
                    <img
                        id="poster"
                        key={imgKey}
                        src=''
                    />
                }
            >
                <Meta
                    title={movie.title}
                    description={<Rate count={10} value={movie.rate} tooltips={desc}/>}
                    style={{textAlign: 'left'}}
                />
            </Card>

            <Flex gap="middle" vertical>
                <Card
                    bordered={false}
                    style={{
                        width: 600,
                    }}
                >
                    <Descriptions title="Movie Info" items={items} column={1}/>
                </Card>
            </Flex>

        </Flex>
    )
}

export default DetailsPage;