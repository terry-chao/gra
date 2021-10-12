import React from 'react';
import './App.css';
// @ts-ignore
import {BrowserRouter, Route} from "react-router-dom";
import routers from "./router";

function App() {
    return (
        <BrowserRouter>
            {
                routers.map(router => {
                    return (
                        <Route
                            path={router.path}
                            component={router.component}
                        />
                    )
                })
            }
        </BrowserRouter>
    );
}

export default App;
