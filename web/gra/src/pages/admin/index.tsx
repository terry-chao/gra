import React, {Component} from 'react';
import { Form, Input, Button, Checkbox } from 'antd';
import axios from 'axios'
axios.defaults.baseURL = 'http://localhost:8080';
axios.defaults.timeout = 8080;
class Admin extends Component <{
    navList ?: any
},
    {
        name ?: string
        password ?: string
    }
    > {

    constructor(props: any) {
        super(props);
        this.state = {
            name: '',
            password: ''
        };

    }

    render() {

        const onFinish = async (values: any) => {
            let data = {
                params:{
                    id: '2',
                    name: 'tony',
                    gender: 'a'
                }
            }
            axios.post('/auth', data).then(res => {
                console.log(res)
            }).catch(error => {
                console.error(error)
            })
        };

        const onFinishFailed = (errorInfo: any) => {
            console.log('Failed:', errorInfo);
        };

        return (
            <Form
                name="basic"
                labelCol={{ span: 8 }}
                wrapperCol={{ span: 16 }}
                initialValues={{ remember: true }}
                onFinish={onFinish}
                onFinishFailed={onFinishFailed}
                autoComplete="off"
            >
                <Form.Item
                    label="Username"
                    name="username"
                    rules={[{ required: true, message: 'Please input your username!' }]}
                >
                    <Input />
                </Form.Item>

                <Form.Item
                    label="Password"
                    name="password"
                    rules={[{ required: true, message: 'Please input your password!' }]}
                >
                    <Input.Password />
                </Form.Item>

                <Form.Item name="remember" valuePropName="checked" wrapperCol={{ offset: 8, span: 16 }}>
                    <Checkbox>Remember me</Checkbox>
                </Form.Item>

                <Form.Item wrapperCol={{ offset: 8, span: 16 }}>
                    <Button type="primary" htmlType="submit">
                        Submit
                    </Button>
                </Form.Item>
            </Form>
        );
    }
}

export default Admin;