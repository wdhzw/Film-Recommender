import React, {useContext, useEffect, useState} from 'react';
import {Descriptions, Flex, Button, Form, Input, Modal, notification} from "antd";
import TopBar from "../components/topBar";
import {AuthContext} from "../hooks/AuthContext";
import {addGenre, getUserInfo} from "../api";

function ProfilePage() {
    const { authData } = useContext(AuthContext);
    const [newGenre, setNewGenre] = useState('')
    const [isModalVisible, setIsModalVisible] = useState(false);
    const [api, contextHolder] = notification.useNotification();
    const [userInfo, setUserInfo] = useState({
        user_name: '',
        email: '',
        preferred_genre: [],
    })
    const openNotificationWithIcon = (type, message, desc) => {
        api[type]({
            message: message,
            description: desc,
        });
    };

    const showModal = () => {
        setIsModalVisible(true);
    };

    const handleOk = async () => {
        setIsModalVisible(false);
        let addGenreRes = await addGenre(authData.userInfo.user_name, authData.userInfo.email, newGenre)
        if (addGenreRes !== null && addGenreRes.status_code === 0) {
            openNotificationWithIcon('success', 'Add genre successfully!!!', null)
            setUserInfo(prevState => ({
                ...prevState,
                preferred_genre: [...prevState.preferred_genre, newGenre]
            }))
        } else {
            openNotificationWithIcon('error', 'Add genre failed!!!', 'This genre may exists!!!')
        }
    };

    const handleCancel = () => {
        setIsModalVisible(false);
    };

    useEffect(() => {
        fetchUserInfo()
    }, []);

    const fetchUserInfo = async () => {
        let userInfoRes = await getUserInfo(authData.userInfo.email)
        if (userInfoRes !== null && userInfoRes.status_code === 0) {
            if (userInfoRes.data.preferred_genre !== undefined) {
                setUserInfo({
                    user_name: userInfoRes.data.user_name,
                    email: userInfoRes.data.email,
                    preferred_genre: userInfoRes.data.preferred_genre,
                })
            } else {
                setUserInfo({
                    user_name: userInfoRes.data.user_name,
                    email: userInfoRes.data.email,
                    preferred_genre: [],
                })
            }

        }
    }

    const items = [
      {
        key: '1',
        label: 'User Name',
        children: userInfo.user_name,
      },
      {
        key: '2',
        label: 'Email',
        children: userInfo.email,
      },
      {
        key: '3',
        label: 'User preferences',
        children: (
            <Flex justify="center" align="center" wrap="wrap">
                {userInfo.preferred_genre.map((genre, index) => (
                    <div key={index} style={{margin: "auto 2px"}}>
                        <Button danger>{genre} </Button>
                    </div>
                ))}
            </Flex>
        ),
      },
    ];

    return (
        <Flex gap="middle" vertical>
            {contextHolder}
            <TopBar username={authData.userInfo.user_name} currentPage={"Profile"}/>
            <Descriptions
                title="User Profile"
                bordered
                items={items}
                column={1}
                style={{
                    background: "white",
                    width: '800px'
            }}
            />
            <Button type='primary' style={{width: '60%', margin: '0 auto'}} onClick={showModal}>Add Genre</Button>
            <Modal title="Add new genre" centered open={isModalVisible} onOk={handleOk} onCancel={handleCancel}>
                <Form
                  name="Add preference"
                  initialValues={{ remember: true }}
                  autoComplete="off"
                  style={{width: '70%', margin: '0 auto'}}
                >
                  <Form.Item
                    label="genre"
                    name="newGenre"
                    rules={[{ required: true, message: 'Please input your genre!' }]}
                    onChange={(e) => setNewGenre(e.target.value)}
                  >
                    <Input />
                  </Form.Item>
                </Form>
              </Modal>
        </Flex>
    )
}

export default ProfilePage;