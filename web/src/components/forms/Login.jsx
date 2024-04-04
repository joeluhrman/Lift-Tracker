import React from "react"
import { Navigate } from "react-router-dom"
import {
    Form,
    Button,
} from "react-bootstrap"
import UserHandler from "../../handlers/UserHandler"
const userHandler = new UserHandler()

export default function Login() {
    const [formValue, setFormValue] = React.useState({
        username: "",
        password: "",
    })
    const [submitError, setSubmitError] = React.useState(null)
    const [toDashboard, setToDashboard] = React.useState(false)

    const handleChange = (event) => {
        setFormValue({ ...formValue, [event.target.name]: event.target.value });
    }

    const handleSubmit = async (event) => {
        event.preventDefault()
        event.stopPropagation()

        const [status, headers, data] = await userHandler.login(
            formValue.username, formValue.password
        )

        if (status === 200) setToDashboard(true)
        else if (status === 401) setSubmitError("Username and/or password are invalid.")
        else if (status === 500) setSubmitError("The server is not responding.")
        else setSubmitError("Unhandled error.")
    }

    if (toDashboard) return <Navigate to="/dashboard"/>

    return (
        <Form noValidate onSubmit={handleSubmit}>
            <Form.Group className="mb-2">
                <Form.Label>Username</Form.Label>
                <Form.Control
                    required
                    name="username"
                    type="text"
                    placeholder="Username"
                    value={formValue.username}
                    maxLength="20"
                    onChange={handleChange}
                />
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
                    maxLength="25"
                />
            </Form.Group>

            <Button type="submit">Login</Button>
            <p className="text-danger">
                {submitError}
            </p>
        </Form>
    )
}