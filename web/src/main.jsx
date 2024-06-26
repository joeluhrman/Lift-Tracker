import React from 'react'
import ReactDOM from 'react-dom/client'
import {
  createBrowserRouter,
  Outlet,
  RouterProvider,
} from "react-router-dom"

import Nav from "./routes/Nav"
import Auth from "./routes/Auth"
import Error from "./routes/Error"
import Register from "./routes/Register"
import Dashboard from "./routes/Dashboard"
import WorkoutTemplates from "./routes/WorkoutTemplates"
import WorkoutHistory from "./routes/WorkoutHistory"
import CRUDWorkout from "./routes/CRUDWorkout"

import 'bootstrap/dist/css/bootstrap.min.css'
import './index.css'

const router = createBrowserRouter([
  { 
    element: <Outlet/>, 
    errorElement: <Error/>, 
    children: [
      { 
        element: <Nav/>,
        children: [
          { path: "/login",     element: <Register variant="login"/>     },
          { path: "/signup",    element: <Register variant="signup" />     },
        ]
      },

      { 
        element: <Auth/>,
        children: [
          { 
            element: <Nav/>,
            children: [
              { path: "/",                      element: <Dashboard/>          },
              { path: "/dashboard",             element: <Dashboard/>          },
              { path: "/workout-templates",     element: <WorkoutTemplates/>   },
              { path: "/workout-history",       element: <WorkoutHistory/>     },

              { path: "/add-workout-template",  element: <CRUDWorkout type="template" variant="create"/>},
              { path: "/edit-workout-template", element: <CRUDWorkout type="template" variant="update"/>},
              { path: "/add-workout-log",       element: <CRUDWorkout type="log"      variant="create"/>},
              { path: "/edit-workout-log",      element: <CRUDWorkout type="log"      variant="update"/>}
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
