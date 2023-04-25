import React from "react"
import {
    Container,
} from "react-bootstrap"
import FormAddWT from "../components/forms/AddWorkoutTemplate"

export default function AddWorkoutTemplate() {
    return (
        <Container>
            <h1>Add Workout Template</h1>
            <FormAddWT/>
        </Container>
    )
}