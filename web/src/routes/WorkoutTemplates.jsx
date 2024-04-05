import React from "react"
import { useNavigate } from "react-router-dom"
import {
    Button,
    Card,
    Container,
} from "react-bootstrap"
import WorkoutTemplateHandler from "../handlers/WorkoutTemplateHandler"
import ExerciseTypeHandler from "../handlers/ExerciseTypeHandler"

export default function WorkoutTemplates() {
    const [temps, setTemps] = React.useState()
    const [exerciseTypes, setExerciseTypes] = React.useState()
    const navigate = useNavigate()

    React.useEffect(() => {
        (async () => {
            const [status, headers, data] = await WorkoutTemplateHandler.getAll()
            setTemps(data)
        })()
    }, [])

    React.useEffect(() => {
        (async () => {
            const [status, headers, data] = await ExerciseTypeHandler.getAll()
            console.log(data)
            setExerciseTypes(data)
        })()
    }, [])

    const handleToAddWT = () => navigate("/add-workout-template")

    const renderWorkoutTemplates = () => {
        return temps.map((temp) => {
            const exerciseElements = temp.exercises.map((eTemp) => {
                const exerciseType = exerciseTypes
                    .find(eType => eType.id === eTemp.exerciseTypeID)
    
                const setgroupElements = eTemp.setgroups.map((sg) => {
                    return (<>
                            {sg.sets} x {sg.reps + " "} 
                    </>)
                })
    
                return (
                    <Container>
                        <p>{ exerciseType.name } { setgroupElements }</p>
                    </Container>
                )
            })
        
            return (
                <Card className="border border-2 w-50 align-items-center mb-2">
                    <h4>{ temp.name }</h4>
                    { exerciseElements }
                </Card>
            )    
        })
    }

    if (temps === undefined || exerciseTypes === undefined) 
        return <Container>Loading...</Container>

    const templateElements = temps === null
        ? "You have not created any workout templates."
        : renderWorkoutTemplates()

    return (
        <>
            <h2>Templates</h2>
            <Container fluid className="d-flex mb-2">
                <Button className="ms-auto" variant="outline-primary" 
                    size="md" onClick={handleToAddWT}>
                    Add Workout Template
                </Button>
            </Container>
            <Container className="row">
                {templateElements}
            </Container>
        </>
    )
}