import { Outlet, useOutletContext } from "react-router-dom"
import { Container } from "react-bootstrap"
import NavHeader from "../components/NavHeader"

export default function Nav() {
    const user = useOutletContext()

    return (<>
        <NavHeader user={user}/>
        <Container>
            <Outlet/>
        </Container>
    </>)
}