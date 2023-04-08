import React from "react"
import { 
    Navigate, 
    Outlet, 
    useLocation 
} from "react-router-dom"

import { handleGetUser } from "../handlers/user"

export default function Auth() {
    const [currentUser, setCurrentUser] = React.useState()
    //const loc = useLocation

    React.useEffect(() => {
        const get = async() => {
            const user = await handleGetUser()
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