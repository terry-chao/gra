import React, {Component} from 'react';
import LoginForm from "../../components/login_form";
import BackgroundImg from "../../components/background_img";
class Index extends Component {
    render() {
        return (
            <div>
                <BackgroundImg/>
                <LoginForm />
            </div>
        );
    }
}

export default Index;