import Admin from "../pages/admin";
import Home from "../pages/home";

import Demo1 from "../pages/routerDemo/demo1";

interface router {
    path:string,
    component:any,
    children?:Array<router>
}

const routers:Array<router> = [
    {
        path:'/',
        component:Admin,
        children:[
            {
                path:'/demo1',
                component:Demo1
            }
        ]
    },
    {
        path:'/home',
        component:Home
    }
]
export default routers