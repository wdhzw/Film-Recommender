import React, {useContext, useEffect, useState} from 'react';
import { Form, Input, Button, Modal, Typography } from 'antd';
import { useNavigate } from 'react-router-dom';
import { AuthContext } from '../hooks/AuthContext';

const { Title } = Typography;

function LoginPage() {

  const loginFormStyle = {
    padding: '40px',
    height: '250px',
    width: '400px',
    border: '1px solid #f0f0f0',
    boxShadow: '0 4px 8px rgba(0, 0, 0, 0.1)',
    borderRadius: '12px',
    background: '#ffffff'
  }

  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const [isModalVisible, setIsModalVisible] = useState(false);
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

  const handleOk = () => {
    setIsModalVisible(false);
  };

  const handleCancel = () => {
    setIsModalVisible(false);
  };

  const handleLogin = async (e) => {
    e.preventDefault()
    console.log(username)
    setAuthData({ username });
    navigate('/dashboard');
  }

  return (
    <div style={loginFormStyle}>
      <Form
        name="basic"
        initialValues={{ remember: true }}
        autoComplete="off"
      >
        <Title level={2}>Login</Title>
        <Form.Item
          label="Email"
          name="email"
          rules={[{ required: true, message: 'Please input your email!' }]}
          value={username}
          onChange={(e) => setUsername(e.target.value)}
        >
          <Input />
        </Form.Item>

        <Form.Item
          label="Pwd"
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
        >
          <Form.Item
            label="Email"
            name="register_email"
            rules={[{ required: true, message: 'Please input your email!' }]}
          >
            <Input />
          </Form.Item>
          <Form.Item
            label="Password"
            name="register_password"
            rules={[{ required: true, message: 'Please input your password!' }]}
          >
            <Input.Password />
          </Form.Item>
        </Form>
      </Modal>
    </div>
  );
}

export default LoginPage;
