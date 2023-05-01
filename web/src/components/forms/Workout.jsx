import React from "react"
import { Button } from "react-bootstrap"
import ExerciseTypeHandler from "../../handlers/ExerciseTypeHandler"

// Form meant to be configurable to be adding/editing a template
// or a log. Currently just works for templates.
//
// CURRENT ISSUE: Adding a new exercise removes all existing
// set group elements (although the form's state is still intact)
export default function Workout() {
    const Exercise = (props) => {
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

            const exercises = workout.exerciseTemplates
            exercises[props.index] = exercise
            setWorkout({...workout, exerciseTemplates: [...exercises]})

            const sgElements = exercise.setgroupTemplates.map((sgTemp, key) => {
                return (
                    <Setgroup 
                        key={key}
                    />
                )
            })
            setSGElements(sgElements)
        }

        return (<>
            <p>Exercise</p>
            <Button onClick={handleAddSetgroup}>Add Setgroup</Button>
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
        setWorkout({...workout, exerciseTemplates: [...exercises]})

        const elements = workout.exerciseTemplates.map((eTemp, key) => {
            return (
                <Exercise
                    key={key}
                    index={key}
                />
            )
        })
        setExerciseElements(elements)
    }

    console.log(workout)

    return (<>
        <p>Form</p>
        <Button onClick={handleAddExercise}>Add Exercise</Button>
        { exerciseElements }
    </>)
}