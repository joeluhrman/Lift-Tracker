import React from "react"
import {
    Button,
    Container,
    Nav,
    Navbar,
    NavDropdown,
} from "react-bootstrap"
import {
  useNavigate
} from "react-router-dom"
import logoUrl from "../assets/lt_nobackground.png"
import UserHandler from "../handlers/UserHandler"

export default function NavHeader(props) {
  const navigate = useNavigate()

  const handleLogout = async() => {
    const [status, headers, data] = await UserHandler.logout()
    if (status === 200) {
      navigate("/login")
    }
  }

  const navItems = props.user &&
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
      <NavDropdown className="ms-auto" title={props.user.username}>
        <NavDropdown.Item>
          Profile
        </NavDropdown.Item>
        <NavDropdown.Item>
          <Button onClick={handleLogout}>Logout</Button>
        </NavDropdown.Item>
      </NavDropdown>
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
          {navItems}
        </Navbar.Collapse>
      </Container>
    </Navbar>
  );
}