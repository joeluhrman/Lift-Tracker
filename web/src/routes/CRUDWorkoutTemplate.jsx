import React from "react"
import {
    Container,
} from "react-bootstrap"
import FormWorkout from "../components/forms/Workout"

export default function CRUDWorkoutTemplate(props) {
    // log vs temp --> variant
    // create vs update --> type

    return (
        <>
            <h2>Add Workout Template</h2>
            <FormWorkout/>
        </>
    )
}