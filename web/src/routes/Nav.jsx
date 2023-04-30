import { Outlet } from "react-router-dom"
import { Container } from "react-bootstrap"
import NavHeader from "../components/NavHeader"

export default function Nav(props) {
    return (<>
        <NavHeader loggedIn={props.loggedIn}/>
        <Container>
            <Outlet/>
        </Container>
    </>)
}