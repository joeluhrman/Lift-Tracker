import axios from "axios"

const handleGetUser = async() => {
    try {
        const res = await axios.get("/api/v1/user")
        return res.data
    } catch(err) {
        return null
    }
}

export { handleGetUser }