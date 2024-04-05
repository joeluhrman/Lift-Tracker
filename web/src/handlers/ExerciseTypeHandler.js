import Handler from "./Handler"

const getAll = async() => {
    const res = await Handler.request("GET", "/api/v1/exercise-type")
    return res
}

const exerciseTypeHandler = {
    getAll,
}

export default exerciseTypeHandler