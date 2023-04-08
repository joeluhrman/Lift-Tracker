import {
    Container,
    Nav,
    Navbar,
    NavDropdown,
} from "react-bootstrap"

export default function NavHeader() {
  return (
    <Navbar bg="light" expand="lg">
      <Container>
        <Navbar.Brand href="/">Lift-Tracker</Navbar.Brand>
        <Navbar.Toggle aria-controls="basic-navbar-nav" />
        <Navbar.Collapse id="basic-navbar-nav">
          <Nav className="me-auto">
            
          </Nav>
        </Navbar.Collapse>
      </Container>
    </Navbar>
  );
}