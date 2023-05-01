import React from "react"
import { Button, Container, Form } from "react-bootstrap"
import ExerciseTypeHandler from "../../handlers/ExerciseTypeHandler"

export default function Workout() {
    const Exercise = (props) => {
        const Setgroup = (props) => {
            return (<>
                <h6> Setgroup { props.index + 1 } </h6>
                <Container className="mb-3">
                    <Form.Group className="mb-3">
                        <Form.Label> Sets </Form.Label>
                        <Form.Control
                        
                        />
                        <Form.Label> Reps </Form.Label>
                        <Form.Control
                        
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

        return (<>
            <Form.Label><h5> Exercise { props.index + 1 } </h5></Form.Label>
            <Button size="sm" className="float-end" onClick={handleAddSetgroup}>Add Setgroup</Button>
            <Container className="mb-3">
                <Form.Select>
                    <option>Exercise</option>
                </Form.Select>
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

    console.log(workout)

    return (
        <Container className="border border-2">
            <Button className="float-end" onClick={handleAddExercise}> Add Exercise </Button>
            <Form.Label><h5> Name </h5></Form.Label>
            <Form.Control
                required
                name="name"
                type="text"
                value={workout.name}
                minLength="1"
                maxLength="100"
                onChange={e => setWorkout({...workout, [e.target.name]: e.target.value})}
            />
            { exerciseElements }
        </Container>
    )
}