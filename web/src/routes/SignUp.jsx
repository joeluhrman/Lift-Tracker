import {
    Container,
} from "react-bootstrap"
import FormSignUp from "../components/forms/SignUp"

export default function SignUp() {
    return(
        <Container className="align-items-center justify-content-center">
            <FormSignUp/>
        </Container>
    )
}