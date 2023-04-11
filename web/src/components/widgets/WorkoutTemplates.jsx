import React from "react"
import {
    Button,
    Container,
    Modal
} from "react-bootstrap"
import WorkoutTemplateHandler from "../../handlers/WorkoutTemplateHandler"
const wtHandler = new WorkoutTemplateHandler()

export default function WorkoutTemplates() {
    const [temps, setTemps]                 = React.useState()
    const [showAddWTForm, setShowAddWTForm] = React.useState(false)

    React.useEffect(() => {
        (async () => {
            const [status, headers, data] = await wtHandler.getAll()
            console.log(data)
            setTemps(data)
        })()
    }, [])

    const handleShowAddWTForm = () => setShowAddWTForm(true)
    const handleCloseAddWTForm = () => setShowAddWTForm(false)

    if (temps === undefined) return <>Loading...</>

    return (
        <Container fluid className="d-flex">
            <Button className="ms-auto" variant="outline-primary" 
                size="md" onClick={handleShowAddWTForm}>
                Add Workout Template
            </Button>

            <Modal show={showAddWTForm} onHide={handleCloseAddWTForm}>
                <Modal.Header>
                    <Modal.Title>Add Workout Template</Modal.Title>
                </Modal.Header>
                <Modal.Body>
                    Form should go here
                </Modal.Body>
                <Modal.Footer>
                    <Button variant="secondary" onClick={handleCloseAddWTForm}>
                        Cancel
                    </Button>
                    <Button variant="primary">
                        Save Changes
                    </Button>
                </Modal.Footer>
            </Modal>
        </Container>
    )
}