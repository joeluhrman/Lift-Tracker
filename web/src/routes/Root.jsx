import { Outlet } from "react-router-dom"
import NavHeader from "../components/NavHeader"

export default function Root() {
    return(<>
        <NavHeader/>
        <Outlet/>
    </>)
}