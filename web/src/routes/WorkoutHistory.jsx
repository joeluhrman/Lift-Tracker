import { useNavigate } from "react-router-dom"
import { Button, Container } from "react-bootstrap"

export default function WorkoutHistory() {
    const navigate = useNavigate()

    const handleToAddWL = () => navigate("/add-workout-log")

    return (<>
        <h2>History</h2>
        <Container fluid className="d-flex mb-2">
            <Button variant="outline-primary" className="ms-auto" onClick={handleToAddWL}>
                Log Workout
            </Button>
        </Container>
    </>)
}