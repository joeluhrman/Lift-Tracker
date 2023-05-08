import { Link } from "react-router-dom"
import { 
    Card,
    Col,
    Container
} from "react-bootstrap"
import FormLogin from "../components/forms/Login"
import FormSignUp from "../components/forms/SignUp"
import logoURL from "../assets/lt.png"

export default function Register(props) {
    let Form, titleText, Footer
    if (props.variant === "login") {
        Form = FormLogin
        titleText = "Login"
        Footer = <>
            Don't have an account? <Link to="/signup">Sign up</Link>
        </>
    } else {
        Form = FormSignUp
        titleText = "Sign Up"
        Footer = <>
            Already have an account? <Link to="/login">Login</Link>
        </>
    }

    return(
        <Container className="d-inline-flex flex-row justify-content-center align-items-center" style={{marginTop:"4%"}}>
            <Card className="d-inline-flex flex-row p-2 border border-2 w-75 align-items-center">
                <Col>
                    <Card.Body>
                        <Card.Title className="text-center">
                            { titleText }
                        </Card.Title>
                        <Form/>
                        <Card.Footer>
                            { Footer }
                        </Card.Footer>
                    </Card.Body>
                </Col>
                <Col>
                    <Card.Img src={logoURL}/>
                </Col>
            </Card>
        </Container>
    )
}