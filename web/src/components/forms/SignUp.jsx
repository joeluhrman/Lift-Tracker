import React from "react"
import {
    Button,
    Form,
} from "react-bootstrap"
import { Navigate } from "react-router-dom"
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
    const [toLogin, setToLogin]         = React.useState(false)

    const handleChange = (event) => {
        setFormValue({ ...formValue, [event.target.name]: event.target.value });
    }

    const handleSubmit = async (event) => {
        event.preventDefault()
        event.stopPropagation()

        setValidated(true)

        const form = event.currentTarget
        if (form.checkValidity() === false) {
            return
        }   

        const [status, headers, data] = await userHandler.create(
            formValue.username, formValue.email, formValue.password
        )

        if (status === 202) {
            setToLogin(true)
        } else if (status === 409) {
            setSubmitError("That username or email is already taken.")
        } else if (status === 500) {
            setSubmitError("The server is not responding.")
        } else {
            setSubmitError("Unhandled error.")
        }
    }

    if (toLogin) {
        return (
            <Navigate to="/login"/>
        )
    }

    return (
        <Form noValidate validated={validated} onSubmit={handleSubmit}>
            <Form.Group className="mb-2">
                <Form.Label>Username</Form.Label>
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
            </Form.Group>
            <Form.Group className="mb-2">
                <Form.Label>Email</Form.Label>
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
            </Form.Group>
            <Form.Group className="mb-2">
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