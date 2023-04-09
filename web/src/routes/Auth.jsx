import React from "react"
import { 
    Navigate, 
    Outlet, 
    useLocation 
} from "react-router-dom"

import UserHandler from "../handlers/UserHandler"
const userHandler = new UserHandler()

export default function Auth() {
    const [currentUser, setCurrentUser] = React.useState()
    //const loc = useLocation

    React.useEffect(() => {
        const handleGetCurrentUser = async() => {
            const res = await userHandler.getUser()
            const user = (
                res.data !== undefined 
                ? res.data 
                : null
            )

            return user
        }

        const user = handleGetCurrentUser
        setCurrentUser(user)
    }, [])

    if (currentUser === undefined) return null

    return currentUser === null
        ? <Outlet/>
        : <Navigate to="/login" /*replace state={{ from: loc }}*//>
}