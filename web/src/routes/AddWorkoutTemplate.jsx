import React from "react"
import {
    Container,
} from "react-bootstrap"
import FormAddWT from "../components/forms/AddWorkoutTemplate"

export default function AddWorkoutTemplate() {
    return (
        <Container>
            <h2>Add Workout Template</h2>
            <FormAddWT/>
        </Container>
    )
}