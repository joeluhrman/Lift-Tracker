import React from 'react'
import ReactDOM from 'react-dom/client'
import {
  createBrowserRouter,
  RouterProvider,
} from "react-router-dom"

import Nav from "./routes/Nav"
import Auth from "./routes/Auth"
import Error from "./routes/Error"
import Root from './routes/Root'
import Register from "./routes/Register"
import Dashboard from "./routes/Dashboard"
import WorkoutTemplates from "./routes/WorkoutTemplates"
import AddWorkoutTemplate from "./routes/AddWorkoutTemplate"

import 'bootstrap/dist/css/bootstrap.min.css'
import './index.css'

const router = createBrowserRouter([
  { 
    element: <Root/>, 
    errorElement: <Error/>, 
    children: [
      { 
        element: <Nav loggedIn={false}/>,
        children: [
          { path: "/login",     element: <Register variant="login"/>     },
          { path: "/signup",    element: <Register variant="signup" />     },
        ]
      },

      { 
        element: <Auth/>,
        children: [
          { 
            element: <Nav loggedIn={true}/>,
            children: [
              { path: "/",                     element: <Dashboard/>         },
              { path: "/dashboard",            element: <Dashboard/>         },
              { path: "/workout-templates",    element: <WorkoutTemplates/>  },
              { path: "/add-workout-template", element: <AddWorkoutTemplate/>}
            ]
          }
        ]
      }
    ]     
  },
])

ReactDOM.createRoot(document.getElementById('root')).render(
  <React.StrictMode>
    <RouterProvider router={router}/>
  </React.StrictMode>,
)
