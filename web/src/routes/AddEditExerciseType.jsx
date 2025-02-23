export default function AddEditExerciseType(props) {
    let titleText = props.variant === "add"
        ? "Add Exercise"
        : "Edit Exercise"

    return (
        <h2>{titleText}</h2>
    )
}