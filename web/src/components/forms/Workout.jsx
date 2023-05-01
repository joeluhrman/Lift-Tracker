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
        const Setgroup = () => {
            return (
                <p>Setgroup</p>
            )
        }
        
        const setgroupElements = workout.exerciseTemplates[props.index]
            .setgroupTemplates.map((sgTemp, key) => {
            return (
                <Setgroup 
                    key={key}
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
            <p>Exercise</p>
            <Button onClick={handleAddSetgroup}>Add Setgroup</Button>
            { setgroupElements }
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

    return (<>
        <p>Form</p>
        <Button onClick={handleAddExercise}>Add Exercise</Button>
        { exerciseElements }
    </>)
}