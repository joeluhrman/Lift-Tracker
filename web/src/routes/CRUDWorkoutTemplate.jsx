import React from "react"
import FormWorkout from "../components/forms/Workout"

export default function CRUDWorkoutTemplate(props) {
    let titleText
    if (props.type === "log" && props.variant === "create")
        titleText = "Log Workout"
    else if (props.type === "log" && props.variant === "update")
        titleText = "Edit Workout Log"
    else if (props.type === "template" && props.variant === "update")
        titleText = "Edit Workout Template"
    else   
        titleText = "Add Workout Template"

    return (
        <>
            <h2>{ titleText }</h2>
            <FormWorkout/>
        </>
    )
}