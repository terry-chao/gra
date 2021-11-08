import React, { Component } from 'react'
import Img from "../assets/images/background.png"

class BackgroundImg extends Component {
  render() {
    return (
      <div>
        <img src={Img} alt="logo"/>
      </div>
    )
  }
}

export default BackgroundImg