import Handler from "./Handler"

const getAll = async() => {
    const res = await Handler.request("GET", "/api/v1/workout-template")
    return res
}

const create = async(workoutTemplate) => {
    const res = await Handler.request("POST", "/api/v1/workout-template", workoutTemplate)
    return res
}

const workoutTemplateHandler = {
    getAll,
    create,
}

export default workoutTemplateHandler