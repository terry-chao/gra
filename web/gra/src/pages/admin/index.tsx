import React, {Component} from 'react';

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

        this.handleNameChange = this.handleNameChange.bind(this);
        this.handlePasswordChange = this.handlePasswordChange.bind(this);
        this.handleSubmit = this.handleSubmit.bind(this);
    }

    handleNameChange(event: { target: { name: any; }; }) {
        this.setState({name: event.target.name});

        console.log(event.target.name ,this.state.name, '111111')
    }

    handlePasswordChange(event: { target: { password: any; }; }) {
        this.setState({password: event.target.password});
    }

    handleSubmit(event: { preventDefault: () => void; }) {
        alert('提交的名字: ' + this.state.name);
        alert('提交的密码: ' + this.state.password);
        event.preventDefault();
    }
    render() {
        return (
            <div>
                {/*<form onSubmit={this.handleSubmit}>*/}
                {/*    <label>*/}
                {/*        账号:*/}
                {/*        <input type="text" onChange={this.handleNameChange} /><br/>*/}
                {/*        密码:*/}
                {/*        <input type="text" /><br/>*/}
                {/*    </label>*/}
                {/*    <input type="submit" value="提交" />*/}
                {/*</form>*/}
                <form action="http://localhost:8080/user" method="post">
                    username:<input type="text" name="username" ></input><br/>
                    password:<input type="text" name="password"></input><br/>
                    <input type="submit" value="登录"></input>

                </form>

            </div>
        );
    }
}

export default Admin;