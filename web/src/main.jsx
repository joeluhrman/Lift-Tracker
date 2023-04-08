import React from 'react'
import ReactDOM from 'react-dom/client'
import {
  createBrowserRouter,
  RouterProvider,
} from "react-router-dom"

import Root from './routes/Root'
import Login from "./routes/Login"
import Dashboard from "./routes/Dashboard"
import SignUp from "./routes/SignUp"

import 'bootstrap/dist/css/bootstrap.min.css'
import './index.css'

const router = createBrowserRouter([
  { 
    path: "/", 
    element: <Root/>,
    children: [
      { path: "login",     element: <Login/>     },
      { path: "signup",    element: <SignUp/>    },
      { path: "dashboard", element: <Dashboard/> },
    ]     
  },
])

ReactDOM.createRoot(document.getElementById('root')).render(
  <React.StrictMode>
    <RouterProvider router={router}/>
  </React.StrictMode>,
)
