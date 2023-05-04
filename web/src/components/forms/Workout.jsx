import React from "react"
import { Button, Col, Container, Form, Row } from "react-bootstrap"
import { useNavigate } from "react-router-dom"
import WorkoutTemplateHandler from "../../handlers/WorkoutTemplateHandler"
import ExerciseTypeHandler from "../../handlers/ExerciseTypeHandler"

// Currently works for adding a workout template,
// except exercise type selector not finished.
//
// Eventually want to expand this form to work for
// adding/editing a template OR a log, depending on
// what props are passed. 
//
// I'd also ideally like to figure out how to un-nest everything
// because it is hard to read.
export default function Workout() {
    const navigate = useNavigate()

    const Exercise = (props) => {
        const Setgroup = (props) => {
            // very inefficient and verbose
            const handleChange = (e) => {
                const work = workout
                const setgroup = work.exerciseTemplates[props.exerciseIndex]
                    .setgroupTemplates[props.index]

                const sg = {...setgroup, [e.target.name]: e.target.value}
                
                work.exerciseTemplates[props.exerciseIndex]
                    .setgroupTemplates[props.index] = sg
                
                setWorkout({...work})
            }

            return (<>
                <Container className="mb-2 d-inline-flex flex-row">
                    <Form.Control
                        required
                        name="sets"
                        type="number"
                        value={workout.exerciseTemplates[props.exerciseIndex]
                            .setgroupTemplates[props.index].sets}
                        onChange={handleChange}
                        className="w-25"
                    /> 
                    <Form.Label> X </Form.Label>
                    <Form.Control
                        required
                        name="reps"
                        type="number"
                        value={workout.exerciseTemplates[props.exerciseIndex]
                            .setgroupTemplates[props.index].reps}
                        onChange={handleChange}
                        className="w-25"
                    />
                </Container>
            </>)
        }
        
        const setgroupElements = workout.exerciseTemplates[props.index]
            .setgroupTemplates.map((sgTemp, index) => {
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
            workout.exerciseTemplates[props.index].setgroupTemplates.push({
                sets: 0,
                reps: 0,
            })
            setWorkout({...work}) 
        }

        const ExerciseSelect = (props) => {
            const options = exerciseTypes.map((eType) => {
                return <option key={Math.random()} value={eType.id}> {eType.name} </option>
            })

            const handleChange = (e) => {
                const work = workout
                work.exerciseTemplates[props.exerciseIndex].exerciseTypeID = e.target.value
                setWorkout({...work})

                console.log(workout.exerciseTemplates[props.exerciseIndex].exerciseTypeID)
            }

            return (
                <Form.Select
                    value={workout.exerciseTemplates[props.exerciseIndex].exerciseTypeID}
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
        exerciseTemplates: [{
            exerciseTypeID: 0,
            setgroupTemplates: [{
                sets: 0,
                reps: 0,
            }]
        }],
    })
    const [exerciseTypes, setExerciseTypes] = React.useState()
    
    const exerciseElements = workout.exerciseTemplates.map((eTemp, index) => {
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
        work.exerciseTemplates.push({
            exerciseTypeID: 0,
            setgroupTemplates: [{
                sets: 0,
                reps: 0
            }],
        })
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