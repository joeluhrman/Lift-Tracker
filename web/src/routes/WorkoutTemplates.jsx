import React from "react"
import { useNavigate } from "react-router-dom"
import {
    Button,
    Container,
    Modal
} from "react-bootstrap"
import AddWTForm from "../components/forms/AddWorkoutTemplate"
import WorkoutTemplateHandler from "../handlers/WorkoutTemplateHandler"
const wtHandler = new WorkoutTemplateHandler()

export default function WorkoutTemplates() {
    const [temps, setTemps] = React.useState()
    const navigate = useNavigate()

    React.useEffect(() => {
        (async () => {
            const [status, headers, data] = await wtHandler.getAll()
            console.log(data)
            setTemps(data)
        })()
    }, [])

    const handleToAddWT = () => navigate("/add-workout-template")

    if (temps === undefined) return <Container>Loading...</Container>

    return (
        <Container>
            <Container fluid className="d-flex">
                <Button className="ms-auto" variant="outline-primary" 
                    size="md" onClick={handleToAddWT}>
                    Add Workout Template
                </Button>
            </Container>
        </Container>
    )
}