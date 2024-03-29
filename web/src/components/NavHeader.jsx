import React from "react"
import {
    Button,
    Container,
    Nav,
    Navbar,
} from "react-bootstrap"
import {
  useNavigate
} from "react-router-dom"
import logoUrl from "../assets/lt_nobackground.png"
import UserHandler from "../handlers/UserHandler"
const userHandler = new UserHandler()

export default function NavHeader(props) {
  const navigate = useNavigate()

  const handleLogout = async() => {
    const [status, headers, data] = await userHandler.logout()
    if (status === 200) {
      navigate("/login")
    }
  }

  const navItems = 
    <Nav className="container-fluid">
      <Nav.Item>
        <Nav.Link href="/workout-history">
          History
        </Nav.Link>
      </Nav.Item>
      <Nav.Item>
        <Nav.Link href="/workout-templates">
          Templates
        </Nav.Link>
      </Nav.Item>
      <Nav.Item className="ms-auto">
        <Button onClick={handleLogout}>Logout</Button>
      </Nav.Item>
    </Nav>

  return (
    <Navbar bg="light" expand="lg" className="border-bottom" fixed="top">
      <Container fluid>
        <Navbar.Brand href="/">
          <img style={{paddingRight: "5%"}}
            src={logoUrl}
            width="50"
            height="35"
            className="d-inline-block align-top"
          />
          Lift-Tracker
        </Navbar.Brand>
        <Navbar.Toggle aria-controls="basic-navbar-nav" />
        <Navbar.Collapse id="basic-navbar-nav">
          {props.loggedIn && navItems}
        </Navbar.Collapse>
      </Container>
    </Navbar>
  );
}