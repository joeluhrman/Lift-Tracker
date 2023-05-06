import React from "react"
import { Button, Col, Container, Form, Row } from "react-bootstrap"
import { useNavigate } from "react-router-dom"
import WorkoutTemplateHandler from "../../handlers/WorkoutTemplateHandler"
import ExerciseTypeHandler from "../../handlers/ExerciseTypeHandler"

// Working on removing nested everything
export default function Workout(props) {
    const navigate = useNavigate()

    const newEmptySetgroup = () => {
        return {
            sets: 0,
            reps: 0
        }
    }

    const newEmptyExercise = () => {
        return {
            exerciseTypeID: 0,
            setgroups: [newEmptySetgroup()],
        }
    }

    const Setgroup = (props) => {
        return (<>
            <Container className="mb-2 d-inline-flex flex-row">
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
            </Container>
        </>)
    }

    // very inefficient and verbose
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

        const handleAddSetgroup = () => {
            const work = workout
            workout.exercises[props.index].setgroups.push(newEmptySetgroup())
            setWorkout({...work}) 
        }

        const ExerciseSelect = (props) => {
            const options = exerciseTypes.map((eType) => {
                return <option key={Math.random()} value={eType.id}> {eType.name} </option>
            })

            const handleChange = (e) => {
                const work = workout
                work.exercises[props.exerciseIndex].exerciseTypeID = e.target.value
                setWorkout({...work})

                console.log(workout.exercises[props.exerciseIndex].exerciseTypeID)
            }

            return (
                <Form.Select
                    value={workout.exercises[props.exerciseIndex].exerciseTypeID}
                    onChange={handleChange}
                    className="w-25"
                >
                    <option>Exercise</option>
                    { options }
                </Form.Select>
            )
        }

        return (<>
            <Button className="float-end" size="sm" onClick={handleAddSetgroup}>Add Setgroup</Button>
            <Form.Group className="mb-2" as={Row}>
                <Form.Label column><h5> Exercise { props.index + 1 } </h5></Form.Label>
                <Col sm="10">
                    <ExerciseSelect exerciseIndex={props.index}/>
                </Col>
            </Form.Group>
            <Container className="mb-2">
                { setgroupElements }
            </Container>
        </>)
    }

    const [workout, setWorkout] = React.useState({
        name: "",
        exercises: [newEmptyExercise()],
    })
    const [exerciseTypes, setExerciseTypes] = React.useState()
    
    const exerciseElements = workout.exercises.map((eTemp, index) => {
        return (
            <Exercise
                key={Math.random()}
                index={index}
            />
        )
    })

    React.useEffect(function getExerciseTypes() {
        (async () => {
            const handler = new ExerciseTypeHandler()
            const [ , , data] = await handler.getAll()
            console.log("ETYPES", data)
            setExerciseTypes(data)
        })()
    }, [])

    const handleAddExercise = () => {
        const work = workout
        work.exercises.push(newEmptyExercise())
        setWorkout({...work})
    }

    const handleSubmit = async (e) => {
        e.preventDefault()
        e.stopPropagation()

        console.log("WORKOUT ", workout)

        const handler = new WorkoutTemplateHandler()
        const [status, , data] = await handler.create(workout)
        
        if (status === 201) {
            navigate("/workout-templates")
        }
    }

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