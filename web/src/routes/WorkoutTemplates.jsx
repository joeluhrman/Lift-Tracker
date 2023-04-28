import React from "react"
import { useNavigate } from "react-router-dom"
import {
    Button,
    Container,
    Modal
} from "react-bootstrap"
import WorkoutTemplateHandler from "../handlers/WorkoutTemplateHandler"
import ExerciseTypeHandler from "../handlers/ExerciseTypeHandler"

export default function WorkoutTemplates() {
    const [temps, setTemps] = React.useState()
    const [exerciseTypes, setExerciseTypes] = React.useState()
    const [tempElements, setTempElements] = React.useState()
    const navigate = useNavigate()

    React.useEffect(() => {
        (async () => {
            const wtHandler = new WorkoutTemplateHandler()
            const [status, headers, data] = await wtHandler.getAll()
            setTemps(data)
        })()
    }, [])

    React.useEffect(() => {
        (async () => {
            const eTypeHandler = new ExerciseTypeHandler()
            const [status, headers, data] = await eTypeHandler.getAll()
            setExerciseTypes(data)
        })()
    }, [])

    React.useEffect(() => {
        if (temps === undefined) return

        const elements = temps.map((temp) => {
            const exerciseElements = temp.exerciseTemplates.map((eTemp) => {
                const exerciseType = exerciseTypes
                    .find(eType => eType.id === eTemp.exerciseTypeID)

                const setgroupElements = eTemp.setgroupTemplates.map((sg) => {
                    return (
                        <Container>
                            <p>{sg.sets} x {sg.reps}</p>
                        </Container>
                    )
                })

                return (
                    <Container>
                        <p>{exerciseType.name}</p>
                        { setgroupElements }
                    </Container>
                )
            })

            return (
                <Container className="border border-2">
                    <p>{ temp.name }</p>
                    { exerciseElements }
                </Container>
            )    
        })

        setTempElements(elements)
    }, [temps])

    const handleToAddWT = () => navigate("/add-workout-template")

    if (temps === undefined || tempElements === undefined || exerciseTypes === undefined) 
        return <Container>Loading...</Container>

    return (
        <Container>
            <h2>Workout Templates</h2>
            <Container fluid className="d-flex mb-2">
                <Button className="ms-auto" variant="outline-primary" 
                    size="md" onClick={handleToAddWT}>
                    Add Workout Template
                </Button>
            </Container>
            {tempElements}
        </Container>
    )
}