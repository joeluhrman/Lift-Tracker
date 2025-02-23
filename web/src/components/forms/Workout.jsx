import React from "react"
import { Button, Col, Container, Form, Row } from "react-bootstrap"
import { Trash } from "react-bootstrap-icons"
import { useNavigate } from "react-router-dom"
import WorkoutTemplateHandler from "../../handlers/WorkoutTemplateHandler"
import ExerciseTypeHandler from "../../handlers/ExerciseTypeHandler"

export default function Workout(props) {
    // logs have a few more fields than templates
    // and submit to a different api endpoint
    const isLog = props.type === "log"

    const newEmptySetgroup = () => {
        const sg = {sets: 0, reps: 0}
        if (isLog) sg.weight = 0
        return sg
    }

    const newEmptyExercise = () => {
        const ex = {exerciseTypeID: 0, setgroups: [newEmptySetgroup()]}
        if (isLog) ex.notes = ""
        return ex
    }

    const newEmptyWorkout = () => {
        const w = {name: "", exercises: [newEmptyExercise()]}
        if (isLog) {
            w.notes = ""
            w.date = ""
        }
        return w
    }

    const [workout, setWorkout] = React.useState(newEmptyWorkout())
    const [exerciseTypes, setExerciseTypes] = React.useState()

    React.useEffect(function getExerciseTypes() {
        (async () => {
            const [ , , data] = await ExerciseTypeHandler.getAll()
            setExerciseTypes(data)
        })()
    }, [])

    const navigate = useNavigate()

    const handleAddSetgroup = (exerciseIndex) => {
        const work = workout
        workout.exercises[exerciseIndex].setgroups.push(newEmptySetgroup())
        setWorkout({...work}) 
    }

    const handleAddExercise = () => {
        const work = workout
        work.exercises.push(newEmptyExercise())
        setWorkout({...work})
    }

    const handleDeleteSetgroup = (exerciseIndex, setgroupIndex) => {
        const work = workout
        work.exercises[exerciseIndex].setgroups.splice(setgroupIndex, 1)
        setWorkout({...work})
    }

    const handleDeleteExercise = (exerciseIndex) => {
        const work = workout
        work.exercises.splice(exerciseIndex, 1)
        setWorkout({...work})
    }

    const Setgroup = (props) => {
        return (
            <Container className="mb-2 d-inline-flex flex-row">
                <Trash width="15" height="15" className="float-start clickable" 
                    onClick={() => handleDeleteSetgroup(props.exerciseIndex, props.index)}/>
                <Form.Control
                    required
                    name="sets"
                    type="number"
                    value={workout.exercises[props.exerciseIndex]
                        .setgroups[props.index].sets}
                    onChange={e => handleSetgroupChange(e, props.index, props.exerciseIndex)}
                    className="w-25"
                /> 
                <Form.Label> X </Form.Label>
                <Form.Control
                    required
                    name="reps"
                    type="number"
                    value={workout.exercises[props.exerciseIndex]
                        .setgroups[props.index].reps}
                    onChange={e => handleSetgroupChange(e, props.index, props.exerciseIndex)}
                    className="w-25"
                />
                { 
                    isLog && <>
                    <Form.Control
                        required
                        name="weight"
                        type="number"
                        value={workout.exercises[props.exerciseIndex]
                        .setgroups[props.index].weight}
                        onChange={e => handleSetgroupChange(e, props.index, props.exerciseIndex)}
                        className="w-25"
                    />
                    <Form.Label>lbs</Form.Label>
                    </>
                }
            </Container>
        )
    }

    const handleSetgroupChange = (e, setgroupIndex, exerciseIndex) => {
        const work = workout
        const setgroup = work.exercises[exerciseIndex]
            .setgroups[setgroupIndex]

        const sg = {...setgroup, [e.target.name]: e.target.value}
        
        work.exercises[exerciseIndex]
            .setgroups[setgroupIndex] = sg
        
        setWorkout({...work})
    }

    const Exercise = (props) => {
        const setgroupElements = workout.exercises[props.index]
            .setgroups.map((sgTemp, index) => {
            return (
                <Setgroup 
                    key={Math.random()}
                    index={index}
                    exerciseIndex={props.index}
                />
            )
        })

        return (
            <Container className="border border-2 mb-2">
                <Button className="float-end" size="sm" onClick={() => handleAddSetgroup(props.index)}>Add Setgroup</Button>
                <Trash className= "float-start clickable" height="20" width="20" onClick={() => handleDeleteExercise(props.index)}/>
                <Form.Group className="mb-2" as={Row}>
                    <Col sm="10">
                        <ExerciseSelect exerciseIndex={props.index}/>
                    </Col>
                </Form.Group>
                <Container className="mb-2">
                    { setgroupElements }
                    { 
                        isLog && <>
                        <Form.Label>Notes</Form.Label>
                        <Form.Control 
                            as="textarea"
                            name="notes"
                            type="text"
                            value={workout.exercises[props.index].notes}
                            className="w-50"
                        />
                        </>
                    }   
                </Container>
            </Container>
        )
    }

    const ExerciseSelect = (props) => {
        const options = exerciseTypes.map((eType) => {
            return <option key={Math.random()} value={eType.id}> {eType.name + (eType.isDefault ? " -- [DEFAULT]" : "")} </option>
        })      

        return (
            <Form.Select
                value={workout.exercises[props.exerciseIndex].exerciseTypeID}
                onChange={e => handleExerciseSelectChange(e, props.exerciseIndex)}
                className="w-25"
            >
                <option>Exercise</option>
                { options }
            </Form.Select>
        )
    }

    const handleExerciseSelectChange = (e, exerciseIndex) => {
        const work = workout
        work.exercises[exerciseIndex].exerciseTypeID = e.target.value
        setWorkout({...work})

        console.log(workout.exercises[exerciseIndex].exerciseTypeID)
    }

    const handleSubmit = async (e) => {
        e.preventDefault()
        e.stopPropagation()

        const [status,,] = await WorkoutTemplateHandler.create(workout)
        
        if (status === 201) {
            navigate("/workout-templates")
        }
    }
    
    const exerciseElements = workout.exercises.map((eTemp, index) => {
        return (<Exercise key={Math.random()} index={index}/>)
    })

    if (exerciseTypes === undefined) return null

    console.log(workout)

    return (
        <Form className="border border-2 p-3" onSubmit={handleSubmit}>
            <Button className="float-end" onClick={handleAddExercise}> Add Exercise </Button>
            <Form.Label><h5> Name </h5></Form.Label>
            <Form.Control
                required
                name="name"
                type="text"
                value={workout.name}
                onChange={e => setWorkout({...workout, [e.target.name]: e.target.value})}
                className="mb-2 w-25"
            />
            { exerciseElements }
            <Button type="submit">Save</Button>
        </Form>
    )
}