import {
    Card,
    Container,
} from "react-bootstrap"
import FormSignUp from "../components/forms/SignUp"

export default function SignUp() {
    return(
        <Container className="d-flex justify-content-center align-items-center">
            <Card className="w-25">
                <Card.Body>
                    <Card.Title className="text-center">
                        Sign Up
                    </Card.Title>

                    <FormSignUp/>

                    <Card.Footer>
                      {/*  Already have an account? <Link to="/login">Login</Link> */}
                    </Card.Footer>
                </Card.Body>
            </Card>
        </Container>
    )
}