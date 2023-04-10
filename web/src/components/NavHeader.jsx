import {
    Button,
    Container,
    Nav,
    Navbar,
} from "react-bootstrap"
import UserHandler from "../handlers/UserHandler"
const userHandler = new UserHandler()

export default function NavHeader() {
  return (
    <Navbar bg="light" expand="lg">
      <Container>
        <Navbar.Brand href="/">Lift-Tracker</Navbar.Brand>
        <Navbar.Toggle aria-controls="basic-navbar-nav" />
        <Navbar.Collapse id="basic-navbar-nav">
          <Nav className="me-auto">
            <Nav.Item>
              <Button onClick={userHandler.logout}>Logout</Button>
            </Nav.Item>
          </Nav>
        </Navbar.Collapse>
      </Container>
    </Navbar>
  );
}