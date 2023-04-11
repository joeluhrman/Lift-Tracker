import React from "react"
import {
    Button,
    Container
} from "react-bootstrap"
import WorkoutTemplateHandler from "../../handlers/WorkoutTemplateHandler"
const wtHandler = new WorkoutTemplateHandler()

export default function WorkoutTemplates() {
    const [temps, setTemps] = React.useState()

    React.useEffect(() => {
        (async () => {
            const [status, headers, data] = await wtHandler.getAll()
            console.log(data)
            setTemps(data)
        })()
    }, [])

    if (temps === undefined) return <>Loading...</>

    return (
        <Container fluid>
            <Button className="ms-auto" variant="outline-primary">
                Add Workout Template
            </Button>
        </Container>
    )
}