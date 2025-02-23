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
import Exercises from './routes/Exercises'
import AddEditWorkout from "./routes/AddEditWorkout"
import AddEditExerciseType from './routes/AddEditExerciseType'
import Profile from './routes/Profile'

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
              { path: "/",                  element: <Dashboard/>},
              { path: "/dashboard",         element: <Dashboard/>},
              { path: "/profile",           element: <Profile/>},
              { path: "/workout-templates", element: <WorkoutTemplates/>},
              { path: "/workout-history",   element: <WorkoutHistory/>},
              { path: "/exercises",         element: <Exercises/>},

              { path: "/add-workout-template",  element: <AddEditWorkout type="template" variant="add"/>},
              { path: "/edit-workout-template", element: <AddEditWorkout  type="template" variant="edit"/>},
              { path: "/add-workout-log",       element: <AddEditWorkout  type="log"      variant="add"/>},
              { path: "/edit-workout-log",      element: <AddEditWorkout  type="log"      variant="edit"/>},

              { path: "/add-exercise-type", element: <AddEditExerciseType variant="add"/>}
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
