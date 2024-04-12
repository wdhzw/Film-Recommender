import React from 'react';
import { Input } from 'antd';
const { Search } = Input;
function SearchBar () {

    const onSearch = (value, _e, info) => console.log(info?.source, value);

    return (
        <Search
          placeholder="input movie name"
          allowClear
          enterButton="Search"
          size="large"
          onSearch={onSearch}
          style={{width: '70%', margin: '0 auto'}}
        />
    )
}

export default SearchBar;
