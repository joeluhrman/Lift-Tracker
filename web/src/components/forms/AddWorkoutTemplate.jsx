import React from "react"
import {
    Button,
    Container,
    Form,
} from "react-bootstrap"
import ExerciseTypeHandler from "../../handlers/ExerciseTypeHandler"

export default function AddWorkoutTemplate() {
    const [exerciseTypes, setExerciseTypes] = React.useState()
    const [exerciseTypeSelect, setExerciseTypeSelect] = React.useState()
    const [formValue, setFormValue] = React.useState({ name: "", exerciseTemplates: [], })
    const [exerciseElements, setExerciseElements] = React.useState([])

    const ExerciseFormGroup = (props) => {
        return (
            <Container className="border border-1 mb-3">
            <Form.Group>
                <Form.Label> {props.order} </Form.Label>
                <br/>
                { exerciseTypeSelect }
            </Form.Group> 
            </Container>
        )
    }

    const setGroupInputGroup = (props) => {

    }

    React.useEffect(() => {
        (async () => {
            const eTypeHandler = new ExerciseTypeHandler()
            const [status, headers, data] = await eTypeHandler.getAll()
            setExerciseTypes(data)
        })()
    }, [])

    React.useEffect(() => {
        const select = (
            <Form.Select>
                <option>Exercise</option>
            </Form.Select>
        )

        setExerciseTypeSelect(select)
    }, [exerciseTypes])

    const handleChange = (event) => {
        setFormValue({ ...formValue, [event.target.name]: event.target.value });
    }

    const handleAddExercise = () => {      
        var exerciseTemplates = formValue.exerciseTemplates
        exerciseTemplates.push({})
        setFormValue({...formValue, exerciseTemplates: [...exerciseTemplates]})

        var elements = exerciseElements
        elements.push(
            <ExerciseFormGroup 
                key={elements.length} 
                order={elements.length + 1}
            />)
        setExerciseElements([...elements])
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
                    
                        { exerciseElements }
                    
                    <Button className="float-end" size="sm" onClick={handleAddExercise}>+</Button>
                </Container>
            </Form.Group>
        </Form>
    )
}