import React from "react"
import { 
    Navigate, 
    Outlet, 
} from "react-router-dom"

import UserHandler from "../handlers/UserHandler"
const userHandler = new UserHandler()

export default function Auth() {
    const [status, setStatus] = React.useState(undefined)

    React.useEffect(() => {
        (async () => {
            const [stat, headers, data] = await userHandler.get()
            setStatus(stat)
        })()
    }, [])

    if (status === undefined) return null

    return (status === 200
        ? <Outlet/>
        : <Navigate to="/login"/>)
}