import React from "react"
import { useNavigate } from "react-router-dom"
import { Button, Container, Table } from "react-bootstrap"
import ExerciseTypeHandler from "../handlers/ExerciseTypeHandler"

export default function Exercises() {
    const [exerciseTypes, setExerciseTypes] = React.useState(null)
    const navigate = useNavigate()

    React.useEffect(() => {
        (async () => {
            const [,, data] = await ExerciseTypeHandler.getAll()
            setExerciseTypes(data)
        })()
    }, [])

    const renderExerciseTypes = () => {
        return exerciseTypes.map((e) => {
            return (
                <tr>
                    <td>{e.name}</td>
                    <td>{e.pplTypes}</td>
                    <td>{e.muscleGroups}</td>
                    <td>{e.isDefault === true 
                        ? "DEFAULT" 
                        : "Delete/Edit icon placeholder"}</td>
                </tr>
            )
        })
    }

    const handleToAddExerciseType = () => navigate("/add-exercise-type")

    return (
        <>
            <h2>Exercises</h2>
            <Container fluid className="d-flex mb-2">
                <Button className="ms-auto" variant="outline-primary" 
                        size="md" onClick={handleToAddExerciseType}>
                        Add Exercise
                </Button>
            </Container>
            <Table striped bordered hover>
                <thead>
                    <tr>
                        <th>Name</th>
                        <th>PPL Types</th>
                        <th>Muscle Groups</th>
                        <th></th>
                    </tr>
                </thead>
                <tbody>
                    { exerciseTypes === null ? null : renderExerciseTypes()} 
                </tbody>
            </Table>
        </>
    )
}