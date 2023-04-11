import React from 'react'
import ReactDOM from 'react-dom/client'
import {
  createBrowserRouter,
  RouterProvider,
} from "react-router-dom"

import Auth from "./routes/Auth"
import Error from "./routes/Error"
import Root from './routes/Root'
import Login from "./routes/Login"
import Dashboard from "./routes/Dashboard"
import SignUp from "./routes/SignUp"
import WorkoutTemplates from "./routes/WorkoutTemplates"

import 'bootstrap/dist/css/bootstrap.min.css'
import './index.css'

const router = createBrowserRouter([
  { 
    element: <Root/>, 
    errorElement: <Error/>, 
    children: [
      { path: "/login",     element: <Login/>     },
      { path: "/signup",    element: <SignUp/>    },

      { 
        element: <Auth/>,
        children: [
          { path: "/",                  element: <Dashboard/>       },
          { path: "/dashboard",         element: <Dashboard/>       },
          { path: "/workout-templates", element: <WorkoutTemplates/>},
        ]
      }
    ]     
  },
])

// also should figure out how to conditionally render 
// the logout button on the navbar

ReactDOM.createRoot(document.getElementById('root')).render(
  <React.StrictMode>
    <RouterProvider router={router}/>
  </React.StrictMode>,
)
