import React from "react"
import { Button } from "react-bootstrap"
import ExerciseTypeHandler from "../../handlers/ExerciseTypeHandler"

// Form meant to be configurable to be adding/editing a template
// or a log. Currently just works for templates.
export default function Workout() {
    const Exercise = () => {
        const [exercise, setExercise] = React.useState({
            exerciseTypeID: undefined,
            setgroupTemplates: []
        })
        const [sgElements, setSGElements] = React.useState()

        const Setgroup = () => {
            return (
                <p>Setgroup</p>
            )
        }

        const handleAddSetgroup = () => {
            const setgroups = exercise.setgroupTemplates
            setgroups.push({})
            setExercise({...exercise, setgroupTemplates: [...setgroups]})

            const elements = exercise.setgroupTemplates.map((sgTemp, key) => {
                return (
                    <Setgroup key={key}/>
                )
            })
            setSGElements(elements)
        }

        return (<>
            <p>Exercise</p>
            <Button>Add Setgroup</Button>
            { sgElements }
        </>)
    }

    const [workout, setWorkout] = React.useState({
        name: "",
        exerciseTemplates: [],
    })
    const [exerciseTypes, setExerciseTypes] = React.useState()
    const [exerciseElements, setExerciseElements] = React.useState()

    React.useEffect(function getExerciseTypes() {
        (async () => {
            const handler = new ExerciseTypeHandler()
            const [ , , data] = await handler.getAll()
            setExerciseTypes(data)
        })()
    }, [])

    const handleAddExercise = () => {
        const exercises = workout.exerciseTemplates
        exercises.push({})
        setWorkout({...workout, exerciseTemplate: [...exercises]})

        const elements = workout.exerciseTemplates.map((eTemp, key) => {
            return (
                <Exercise
                    key={key}
                />
            )
        })
        setExerciseElements(elements)
    }

    return (<>
        <p>Form</p>
        <Button onClick={handleAddExercise}>Add Exercise</Button>
        { exerciseElements }
    </>)
}