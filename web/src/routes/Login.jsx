import {
    Card,
    Container,
} from "react-bootstrap"
import { Link } from "react-router-dom"
import FormLogin from "../components/forms/Login"

export default function Login() {
    return(
        <Container className="d-flex justify-content-center align-items-center">
            <Card className="w-25">
                <Card.Body>
                    <Card.Title className="text-center">
                        Login
                    </Card.Title>

                    <FormLogin/>

                    <Card.Footer>
                        Don't have an account? <Link to="/signup">Sign up</Link>
                    </Card.Footer>
                </Card.Body>
            </Card>
        </Container>
    )
}