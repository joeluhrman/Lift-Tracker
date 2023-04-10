import React from "react"
import {
    Button,
    Form,
    InputGroup,
} from "react-bootstrap"
import { Link } from "react-router-dom"
import UserHandler from "../../handlers/UserHandler"

const userHandler = new UserHandler()

export default function SignUp() {
    const [validated, setValidated] = React.useState(false)
    const [formValue, setFormValue] = React.useState({
        username:        "",
        email:           "",
        password:        "",
        confirmPassword: "",
    })
    const [submitError, setSubmitError] = React.useState(null)

    const handleChange = (event) => {
        setFormValue({ ...formValue, [event.target.name]: event.target.value });
    }

    const handleSubmit = async(event) => {
        event.preventDefault()
        event.stopPropagation()

        setValidated(true)

        const form = event.currentTarget
        if (form.checkValidity() === false) {
            return
        }   

        const [successful, data] = await userHandler.createUser(
            formValue.username, formValue.email, formValue.password
        )
        console.log(data)
        console.log(successful)

        // success
        if(successful) {
            return
        } else {
            setSubmitError(data)
            console.log(submitError)
            return
        }
    }

    return (
        <Form noValidate validated={validated} onSubmit={handleSubmit}>
            <Form.Group>
                <Form.Label>Username</Form.Label>
                {/*<InputGroup hasValidation>*/}
                <Form.Control
                    required
                    name="username"
                    type="text"
                    placeholder="Username"
                    value={formValue.username}
                    minLength="3"
                    maxLength="20"
                    onChange={handleChange}
                />
                <Form.Control.Feedback type="valid">
                    Looks good!
                </Form.Control.Feedback>
                <Form.Control.Feedback type="invalid">
                    Must be 3-20 characters.
                </Form.Control.Feedback>
                {/*</InputGroup>*/}
            </Form.Group>
            <Form.Group>
                <Form.Label>Email</Form.Label>
                {/*<InputGroup hasValidation>*/}
                <Form.Control 
                    required
                    name="email"
                    type="email"
                    placeholder="Email"
                    value={formValue.email}
                    onChange={handleChange}
                />
                <Form.Control.Feedback type="valid">
                    Looks good!
                </Form.Control.Feedback>
                <Form.Control.Feedback type="invalid">
                    Please enter a valid email.
                </Form.Control.Feedback>
                {/*</InputGroup>*/}
            </Form.Group>
            <Form.Group>
                <Form.Label>Password</Form.Label>
                <Form.Control
                    required
                    name="password"
                    type="password"
                    placeholder="Password"
                    value={formValue.password}
                    onChange={handleChange}
                    minLength="8"
                    maxLength="25"
                />
                <Form.Control.Feedback type="valid">
                    Looks good!
                </Form.Control.Feedback>
                <Form.Control.Feedback type="invalid">
                    Must be 8-25 characters.
                </Form.Control.Feedback>
            </Form.Group>
            
            <Button type="submit">Sign Up</Button>
            <p className="text-danger">
                {submitError}
            </p>
        </Form>
    ) 
}