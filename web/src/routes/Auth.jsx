import React from "react"
import { 
    Navigate, 
    Outlet, 
} from "react-router-dom"

import UserHandler from "../handlers/UserHandler"

export default function Auth() {
    const [status, setStatus] = React.useState()
    const [user, setUser] = React.useState()

    React.useEffect(() => {
        (async () => {
            const [stat, headers, data] = await UserHandler.get()
            setStatus(stat)
            setUser(data)
        })()
    }, [])

    if (user === undefined) return null

    return (status === 200
        ? <Outlet context={user}/>
        : <Navigate to="/login"/>)
}