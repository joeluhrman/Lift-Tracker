import React from "react"
import {
    Button,
    Container,
    Form,
} from "react-bootstrap"

export default function AddWorkoutTemplate() {
    const [formValue, setFormValue] = React.useState({
        name: "",
        exercise_templates: [],
    })
    const [exerciseElements, setExerciseElements] = React.useState([])

    const handleChange = (event) => {
        setFormValue({ ...formValue, [event.target.name]: event.target.value });
    }

    const handleAddExercise = () => {
        console.log("yep")
        var temp = exerciseElements
        temp.push(<Form.Group><Form.Label>Test element</Form.Label></Form.Group>)
        console.log(temp)
        setExerciseElements(temp)
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
                <Container>
                    <Button className="float-end" size="sm" onClick={handleAddExercise}>+</Button>
                    {exerciseElements}
                </Container>
            </Form.Group>
        </Form>
    )
}