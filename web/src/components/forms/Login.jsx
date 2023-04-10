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

    const handleSubmit = async(event) => {
        event.preventDefault()
        event.stopPropagation()

        const [status, headers, data] = await userHandler.login(
            formValue.username, formValue.password
        )

        console.log(status, data)
    }

    if (toDashboard) return <Navigate to="/dashboard"/>

    return (
        <Form noValidate onSubmit={handleSubmit}>
            <Button type="submit">Login</Button>
        </Form>
    )
}