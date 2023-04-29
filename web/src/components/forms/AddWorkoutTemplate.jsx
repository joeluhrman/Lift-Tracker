import React from "react"
import {
    Button,
    Container,
    Form,
} from "react-bootstrap"
import ExerciseTypeHandler from "../../handlers/ExerciseTypeHandler"

export default function AddWorkoutTemplate() {
    const [exerciseTypes, setExerciseTypes] = React.useState()
    const [formValue, setFormValue] = React.useState({ name: "", exerciseTemplates: [], })
    const [exerciseFormGroups, setExerciseFormGroups] = React.useState([])

    const ExerciseFormGroup = (props) => {
        return (
            <Container className="border border-1 mb-3">
                <Form.Group>
                    <h6>Exercise {props.order}</h6>
                    <Container>
                        <ExerciseTypeSelect exerciseTypes={exerciseTypes}/>
                        <Form.Label>Set Groups</Form.Label>
                        <Button 
                            size="sm"
                            className="float-end"
                            onClick={(() => { handleAddSetgroup(props.index) })}>
                                +
                        </Button>
                    </Container>
                </Form.Group> 
            </Container>
        )
    }

    const SetGroupFormGroup = (props) => {
        <Container className="mb-3">
            <Form.Group>
                <Form.Label>Setgroup {props.order}</Form.Label>
                <Container>
                    <Form.Control

                    />
                    <Form.Control
                    
                    />
                </Container>
            </Form.Group>
        </Container>
    }

    const ExerciseTypeSelect = (props) => {
        const options = props.exerciseTypes.map((eType) => {
            return (
                <option>{eType.name}</option>
            )
        })

        return (
            <Form.Select>
                <option>Exercise</option>
                { options }
            </Form.Select>
        )
    }

    React.useEffect(() => {
        (async () => {
            const eTypeHandler = new ExerciseTypeHandler()
            const [status, headers, data] = await eTypeHandler.getAll()
            setExerciseTypes(data)
        })()
    }, [])

    const handleChange = (event) => {
        setFormValue({ ...formValue, [event.target.name]: event.target.value });
    }

    const handleAddExercise = () => {      
        var exerciseTemplates = formValue.exerciseTemplates
        exerciseTemplates.push({
            exerciseTypeID: 0,
            setgroupTemplates: []
        })
        setFormValue({...formValue, exerciseTemplates: [...exerciseTemplates]})

        var elements = exerciseFormGroups
        elements.push(
            <ExerciseFormGroup 
                key={elements.length} 
                index={elements.length}
                order={elements.length + 1}
            />)
        setExerciseFormGroups([...elements])
    }

    const handleAddSetgroup = (exerciseIndex) => {
        const exerciseTemplates = formValue.exerciseTemplates
        exerciseTemplates[exerciseIndex].setgroupTemplates.push({
            sets: 0,
            reps: 0,
        })
        setFormValue({...formValue, exerciseTemplates: [...exerciseTemplates]})
    }

    if (exerciseTypes === undefined) return <Container>Loading...</Container>

    console.log(formValue)

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
                <Button className="float-end" size="sm" onClick={handleAddExercise}>+</Button>
                <Container>
                    { exerciseFormGroups }
                </Container>
            </Form.Group>

            <Button className="float-end">Save</Button>
        </Form>
    )
}