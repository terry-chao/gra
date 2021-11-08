import React, { Component } from 'react'
import Img from "../assets/images/background.png"
import css from "./index.module.css"

class BackgroundImg extends Component {
  render() {
    return (
      <div>
        <img src={Img} className={css.pages_img} alt="logo"/>
      </div>
    )
  }
}

export default BackgroundImg