import React from "react"
import {
    Button,
    Container,
    Form,
} from "react-bootstrap"

export default function AddWorkoutTemplate() {
    const [formValue, setFormValue] = React.useState({
        name: "",
    })

    const handleChange = (event) => {
        setFormValue({ ...formValue, [event.target.name]: event.target.value });
    }
    
    return (
        <Form noValidate>
            <Form.Group className="mb-2">
                <Form.Label>Template name</Form.Label>
                <Form.Control 
                    required
                    name="name"
                    type="text"
                    placeholder="Template name"
                    value={formValue.name}
                    minLength="1"
                    maxLength="25"
                    onChange={handleChange}
                />
            </Form.Group>
            <Form.Group className="mb-2">
                <Form.Label>Exercises</Form.Label>
                <Button className="float-end" size="sm">+</Button>
            </Form.Group>
        </Form>
    )
}