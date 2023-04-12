import { Outlet } from "react-router-dom"
import NavHeader from "../components/NavHeader"

export default function Nav(props) {
    return (<>
        <NavHeader loggedIn={props.loggedIn}/>
        <Outlet/>
    </>)
}