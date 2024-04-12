import React from 'react';
import { List, Card } from 'antd';
import './movieList.css';

const data = [
  {
    title: 'Movie 1',
    url: '/logo192.png'
  },
  {
    title: 'Movie 2',
    url: '/logo192.png'
  },
  {
    title: 'Movie 3',
    url: '/logo192.png'
  },
    {
    title: 'Movie 4',
    url: '/logo192.png'
  },
    {
    title: 'Movie 5',
    url: '/logo192.png'
  },
];

const MoiveList = () => {
  return (
    <List
      grid={{ gutter: 16, columns: 5 }}
      dataSource={data}
      renderItem={item => (
        <List.Item>
          <Card
            hoverable
            cover={<img alt={item.title} src={item.url} />}
          >
            {item.title}
          </Card>
        </List.Item>
      )}
    />
  );
};

export default MoiveList;
