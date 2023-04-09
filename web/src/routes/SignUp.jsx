import {
    Container,
} from "react-bootstrap"
import FormSignUp from "../components/forms/SignUp"

export default function SignUp() {
    return(
        <Container className="d-flex justify-content-center align-items-center">
            <FormSignUp/>
        </Container>
    )
}