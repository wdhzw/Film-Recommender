import React, {useContext, useEffect, useState} from 'react';
import { Form, Input, Button, Modal, Typography, notification } from 'antd';
import { useNavigate } from 'react-router-dom';
import { AuthContext } from '../hooks/AuthContext';
import {confirmSignUp, getUserInfo, login, signUp} from '../api'

const { Title } = Typography;

function LoginPage() {

  const loginFormStyle = {
    padding: '40px',
    height: '250px',
    width: '400px',
    margin: '0 auto',
    border: '1px solid #f0f0f0',
    boxShadow: '0 4px 8px rgba(0, 0, 0, 0.1)',
    borderRadius: '12px',
    background: '#ffffff'
  }

  const [api, contextHolder] = notification.useNotification();
  const openNotificationWithIcon = (type, message, desc) => {
    api[type]({
      message: message,
      description: desc,
    });
  };

  const [username, setUsername] = useState("");
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [isModalVisible, setIsModalVisible] = useState(false);
  const [isVerifyCode, setIsVerifyCode] = useState(false)
  const [verifyCode, setVerifyCode] = useState('')
  const { setAuthData } = useContext(AuthContext);
  const navigate = useNavigate();


  useEffect(() => {
    const data = localStorage.getItem('authData');
    if (data) {
      navigate('/dashboard');
    }
  }, []);

  const showModal = () => {
    setIsModalVisible(true);
  };

  const handleOk = async () => {
    if (isVerifyCode === false) {
      let signUpRes = await signUp(username, email, password)
      if (signUpRes !== null && signUpRes.status_code === 0) {
        setIsVerifyCode(true)
        return
      } else {
        openNotificationWithIcon('error', 'Sign up failed!!', 'Password too simple!!!')
      }
    }
    if (isVerifyCode === true && verifyCode !== '') {
      let confirmSighUpRes = await confirmSignUp(username, email, verifyCode)
      if (confirmSighUpRes !== null && confirmSighUpRes.status_code === 0) {
        setIsModalVisible(false);
        openNotificationWithIcon('success', 'Sign up successfully!!', 'Pls refresh and login!!!')
      }
    }

  };

  const handleCancel = () => {
    setIsModalVisible(false);
  };

  const handleLogin = async (e) => {
    e.preventDefault()
    let loginRes = await login(username, email, password)
    if (loginRes !== null && loginRes.status_code === 0) {
      let userInfoRes = await getUserInfo(email)

      if (userInfoRes !== null && userInfoRes.status_code === 0) {
        let userInfo = userInfoRes.data
        setAuthData({ userInfo });
        navigate('/dashboard');
      }

    } else {
      openNotificationWithIcon('error', 'Login Failed!!', 'Pls correct your username/password!!!')
    }

  }

  return (
    <div style={loginFormStyle}>
      {contextHolder}
      <Form
        name="basic"
        initialValues={{ remember: true }}
        autoComplete="off"
        style={{width: '80%', margin: '0 auto'}}
      >
        <Title level={2}>Login</Title>
        <Form.Item
          label="User Name"
          name="username"
          rules={[{ required: true, message: 'Please input your username!' }]}
          value={username}
          onChange={(e) => setUsername(e.target.value)}
        >
          <Input />
        </Form.Item>

        <Form.Item
          label="Email"
          name="email"
          rules={[{ required: true, message: 'Please input your email!' }]}
          value={email}
          onChange={(e) => setEmail(e.target.value)}
        >
          <Input />
        </Form.Item>

        <Form.Item
          label="Password"
          name="password"
          rules={[{ required: true, message: 'Please input your password!' }]}
          value={password}
          onChange={(e) => setPassword(e.target.value)}
        >
          <Input.Password />
        </Form.Item>

        <Form.Item>
          <Button type="primary" htmlType="submit" onClick={handleLogin}>
            Go
          </Button>
        </Form.Item>
      </Form>
      <Button type="link" onClick={showModal}>
        Sign Up
      </Button>
      <Modal title="Sign Up" centered open={isModalVisible} onOk={handleOk} onCancel={handleCancel}>
        <Form
          name="register"
          initialValues={{ remember: true }}
          autoComplete="off"
          style={{width: '70%', margin: '0 auto'}}
        >
          <Form.Item
            label="User Name"
            name="username"
            rules={[{ required: true, message: 'Please input your username!' }]}
            onChange={(e) => setUsername(e.target.value)}
          >
            <Input />
          </Form.Item>
          <Form.Item
            label="Email"
            name="register_email"
            rules={[{ required: true, message: 'Please input your email!' }]}
            onChange={(e) => setEmail(e.target.value)}
          >
            <Input />
          </Form.Item>
          <Form.Item
            label="Password"
            name="register_password"
            rules={[{ required: true, message: 'Please input your password!' }]}
            onChange={(e) => setPassword(e.target.value)}
          >
            <Input.Password />
          </Form.Item>
          {
            isVerifyCode ?
                <Form.Item
                  label="Verify Code"
                  name="verifyCode"
                  rules={[{ required: true, message: 'Please input the code from your mailbox!' }]}
                  onChange={(e) => setVerifyCode(e.target.value)}
                >
                  <Input.Password />
                </Form.Item> : null
          }

        </Form>
      </Modal>
    </div>
  );
}

export default LoginPage;
