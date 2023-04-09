import React from "react"
import { 
    Navigate, 
    Outlet, 
    useLocation 
} from "react-router-dom"

import UserHandler from "../handlers/UserHandler"

export default function Auth() {
    const [currentUser, setCurrentUser] = React.useState()
    //const loc = useLocation

    React.useEffect(() => {
        const get = async() => {
            const handler = new UserHandler()
            const user = await handler.getUser()
            return user
        }

        const user = get()
        setCurrentUser(user)
    }, [])

    if (currentUser === undefined) return null

    return currentUser === null
        ? <Outlet/>
        : <Navigate to="/login" /*replace state={{ from: loc }}*//>
}