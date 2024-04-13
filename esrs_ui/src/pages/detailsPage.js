import React from 'react';
import { useParams } from 'react-router-dom';
import {Flex, Descriptions, Rate, Card} from 'antd';
import TopBar from "../components/topBar";

const { Meta } = Card;
const desc = ['terrible', 'bad', 'normal', 'good', 'wonderful'];


function DetailsPage() {
    const { movieID } = useParams();

    const items = [
      {
        key: '1',
        label: 'Release Year',
        children: '2024',
      },
      {
        key: '2',
        label: 'Quality',
        children: 'HD',
      },
      {
        key: '3',
        label: 'Release Date',
        children: '04-11',
      },
      {
        key: '4',
        label: 'Director',
        children: 'Director',
      },
      {
        key: '5',
        label: 'stars',
        children: 'Actor list',
      },
        {
            key: '6',
            label: 'genre',
            children: 'horor'
        }
    ];

    return (
        <Flex gap="middle">
            <TopBar/>
            <Card
                style={{ width: 300 }}
                cover={
                  <img
                    alt="movie name"
                    src="https://gw.alipayobjects.com/zos/rmsportal/JiqGstEfoWAOHiTxclqi.png"
                  />
                }
              >
                <Meta
                  title="Movie Name"
                  description={<Rate value={3} tooltips={desc}/>}
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