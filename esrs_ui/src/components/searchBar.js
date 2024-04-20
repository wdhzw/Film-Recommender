import React, {useState} from 'react';
import {AutoComplete, Input} from 'antd';
import {searchMovie} from "../api";
import {debounce} from "lodash";
import {useNavigate} from "react-router-dom";
const { Search } = Input;
function SearchBar () {

    const navigate = useNavigate();
    const [options, setOptions] = useState([]);
    const searchResult = (contents) => {
        return contents.map((movie, idx) => {
            console.log(idx, movie)
            return {
                value: movie.title,
                label: (
                    <div>
                        <a onClick={() => navigate('/movie/' + movie.movie_id)}>{movie.title}</a>
                    </div>)
            }
        })
    }

    const handleSearch = debounce(async (value) => {
        let searchRes = await searchMovie(value)
        console.log(searchRes)
        if (searchRes !== null) {
            setOptions(value ? searchResult(searchRes.content) : []);
        }
    }, 500);

    const onSelect = (value) => {
        console.log('onSelect', value);
    };

    return (
        <AutoComplete
          popupMatchSelectWidth={false}
          style={{
            width: '70%', margin: '0 auto'
          }}
          options={options}
          onSelect={onSelect}
          onSearch={handleSearch}
          size="large"
        >
            <Input.Search size="large" placeholder="input here" enterButton />
        </AutoComplete>
    )
}

export default SearchBar;
