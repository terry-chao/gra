import React from 'react';
import logo from './logo.svg';
import './App.css';
import Header from "./components/header";
// @ts-ignore
import { BrowserRouter,Route,Switch } from "react-router-dom";
import routers from "./router";

function App() {
  return (
      <BrowserRouter>
          {
              routers.map(router=>{
                  return (
                      <Route
                          path={router.path}
                          component = { router.component }
                      ></Route>
                  )
              })
          }
      </BrowserRouter>
  );
}

export default App;
