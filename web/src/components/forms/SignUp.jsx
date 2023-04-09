import React from "react"
import {
    Button,
    Card,
    Form,
    InputGroup,
} from "react-bootstrap"
import UserHandler from "../../handlers/UserHandler"

const userHandler = new UserHandler()

export default function SignUp() {
    const [validated, setValidated] = React.useState(false)
    const [formValue, setFormValue] = React.useState({
        username: "",
        email: "",
        password: "",
        confirmPassword: ""
    })

    const handleChange = (event) => {
        setFormValue({ ...formValue, [event.target.name]: event.target.value });
    }

    const handleSubmit = async(event) => {
        event.preventDefault()

        const form = event.currentTarget
        if (form.checkValidity() === false) {
            event.stopPropagation()
        }   
        setValidated(true)

        const res = await userHandler.createUser(
            formValue.username, formValue.email, formValue.password
        )
        console.log(res)
    }

    return (
        <Card className="w-25">
            <Card.Body>
                <Card.Title className="text-center">
                    Sign Up
                </Card.Title>

                <Form noValidate validated={validated} onSubmit={handleSubmit}>
                    <Form.Group>
                        <Form.Label>Username</Form.Label>
                        <InputGroup hasValidation>
                            <Form.Control
                                required
                                name="username"
                                type="text"
                                placholder="Username"
                                value={formValue.username}
                                onChange={handleChange}
                            />
                            <Form.Control.Feedback type="valid">
                                Looks good!
                            </Form.Control.Feedback>
                            <Form.Control.Feedback type="invalid">
                                Please provide a username.
                            </Form.Control.Feedback>
                        </InputGroup>
                    </Form.Group>

                    <Button type="submit">Sign Up</Button>
                </Form>

                <Card.Footer>
                    Already have an account? Login
                </Card.Footer>
            </Card.Body>
        </Card>
    ) 
}