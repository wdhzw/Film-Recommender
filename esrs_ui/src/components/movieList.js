import React from 'react';
import { List, Card } from 'antd';
import './movieList.css';

const data = [
  {
    title: 'To Mom (and Dad) with love',
    url: 'https://image.tmdb.org/t/p/original/sChfCU3PDV3N6nYessVPWeWkUBc.jpg'
  },
  {
    title: 'The book of love',
    url: 'https://image.tmdb.org/t/p/original/hwP0GEP0zy8ar965Xaht19SmMd3.jpg'
  },
  {
    title: 'Endless Love',
    url: 'https://image.tmdb.org/t/p/original/z7FZP6uivgVc4t0mnmia0B8YygW.jpg'
  },
    {
    title: 'Love Strange Love',
    url: 'https://image.tmdb.org/t/p/original/9CNnxpI6H8ynyOlACRc25vqgJBY.jpg'
  },
    {
    title: 'Sorry If I Call You Love',
    url: 'https://image.tmdb.org/t/p/original/pnSXPKQPjVi87YEeRYlbg5aUaGs.jpg'
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
            cover={
            <img
                style={{
                    height: '300px',
                    width: '250px'
                }}
                alt={item.title}
                src={item.url}
            />}
          >
            {item.title}
          </Card>
        </List.Item>
      )}
    />
  );
};

export default MoiveList;
