import React from "react"
import {
    Card,
    Form
} from "react-bootstrap"
import UserHandler from "../../handlers/UserHandler"

export default function SignUp() {
    const [validated,       setValidated]       = React.useState(false)
    const [username,        setUsername]        = React.useState("")
    const [email,           setEmail]           = React.useState("")
    const [password,        setPassword]        = React.useState("")
    const [confirmPassword, setConfirmPassword] = React.useState("")

    const handleSubmit = async() => {
        const handler = new UserHandler()
        const status = await handler.createUser()
    }

    return (
        <Card>
            <Card.Body>
                <Card.Title className="text-center">
                    Sign Up
                </Card.Title>

                <Form>

                </Form>

                <Card.Footer>
                    Already have an account? Login
                </Card.Footer>
            </Card.Body>
        </Card>
    ) 
}