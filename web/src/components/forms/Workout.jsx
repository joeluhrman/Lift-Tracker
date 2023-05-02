import React from "react"
import { Button, Container, Form } from "react-bootstrap"
import WorkoutTemplateHandler from "../../handlers/WorkoutTemplateHandler"
import ExerciseTypeHandler from "../../handlers/ExerciseTypeHandler"

// Currently works for adding a workout template,
// except exercise type selector not finished.
//
// Eventually want to expand this form to work for
// adding/editing a template OR a log, depending on
// what props are passed. 
export default function Workout() {
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
                <Form.Label><h6> Setgroup { props.index + 1 } </h6></Form.Label>
                <Container className="mb-3">
                    <Form.Group className="mb-3">
                        <Form.Label> Sets </Form.Label>
                        <Form.Control
                            required
                            name="sets"
                            type="number"
                            value={workout.exerciseTemplates[props.exerciseIndex]
                                .setgroupTemplates[props.index].sets}
                            onChange={handleChange}
                        />
                        <Form.Label> Reps </Form.Label>
                        <Form.Control
                            required
                            name="reps"
                            type="number"
                            value={workout.exerciseTemplates[props.exerciseIndex]
                                .setgroupTemplates[props.index].reps}
                            onChange={handleChange}
                        />
                    </Form.Group>
                </Container>
            </>)
        }
        
        const setgroupElements = workout.exerciseTemplates[props.index]
            .setgroupTemplates.map((sgTemp, key) => {
            return (
                <Setgroup 
                    key={key}
                    index={key}
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
                return <option value={eType.id}> {eType.name} </option>
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
                >
                    <option>Exercise</option>
                    { options }
                </Form.Select>
            )
        }

        return (<>
            <Form.Label><h5> Exercise { props.index + 1 } </h5></Form.Label>
            <Button size="sm" className="float-end" onClick={handleAddSetgroup}>Add Setgroup</Button>
            <Container className="mb-3">
                <ExerciseSelect exerciseIndex={props.index}/>
                { setgroupElements }
            </Container>
        </>)
    }

    const [workout, setWorkout] = React.useState({
        name: "",
        exerciseTemplates: [],
    })
    const [exerciseTypes, setExerciseTypes] = React.useState()
    
    const exerciseElements = workout.exerciseTemplates.map((eTemp, key) => {
        return (
            <Exercise
                key={key}
                index={key}
            />
        )
    })

    React.useEffect(function getExerciseTypes() {
        (async () => {
            const handler = new ExerciseTypeHandler()
            const [ , , data] = await handler.getAll()
            setExerciseTypes(data)
        })()
    }, [])

    const handleAddExercise = () => {
        const work = workout
        work.exerciseTemplates.push({
            exerciseTypeID: 0,
            setgroupTemplates: [],
        })
        setWorkout({...work})
    }

    const handleSubmit = async (e) => {
        e.preventDefault()
        e.stopPropagation()

        console.log("WORKOUT ", workout)

        const handler = new WorkoutTemplateHandler()
        const [status, , data] = await handler.create(workout)
        
        console.log("STATUS", status)
        console.log("DATA", data)
    }

    if (exerciseTypes === undefined) return null

    return (
        <Form className="border border-2" onSubmit={handleSubmit}>
            <Button className="float-end" onClick={handleAddExercise}> Add Exercise </Button>
            <Form.Label><h5> Name </h5></Form.Label>
            <Form.Control
                required
                name="name"
                type="text"
                value={workout.name}
                onChange={e => setWorkout({...workout, [e.target.name]: e.target.value})}
            />
            { exerciseElements }
            <Button className="float-end" type="submit">Save</Button>
        </Form>
    )
}