import {
    Button,
    Container,
    Nav,
    Navbar,
} from "react-bootstrap"
import logoUrl from "../assets/lt.png"
import UserHandler from "../handlers/UserHandler"
const userHandler = new UserHandler()

export default function NavHeader() {
  return (
    <Navbar bg="light" expand="lg" className="border-bottom" fixed="top">
      <Container>
        <Navbar.Brand href="/">
          <img style={{paddingRight: "5%"}}
            src={logoUrl}
            width="50"
            height="30"
            className="d-inline-block align-top"
          />
          Lift-Tracker
        </Navbar.Brand>
        <Navbar.Toggle aria-controls="basic-navbar-nav" />
        <Navbar.Collapse id="basic-navbar-nav">
          <Nav className="container-fluid">
            <Nav.Item className="ms-auto">
              <Button onClick={userHandler.logout}>Logout</Button>
            </Nav.Item>
          </Nav>
        </Navbar.Collapse>
      </Container>
    </Navbar>
  );
}