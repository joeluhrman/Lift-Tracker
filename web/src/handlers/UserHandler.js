import Handler from "./Handler"

const get = async() => {
    const res = await Handler.request("GET", "/api/v1/user")
    return res
}

const create = async(username, email, password) => {
    const body = {
        username: username,
        email: email,
        password: password,
    }

    const res = await Handler.request("POST", "/api/v1/user", body)
    return res
}

const login = async(username, password) => {
    const body = {
        username: username,
        password: password,
    }

    const res = await Handler.request("POST", "/api/v1/login", body)
    return res
}

const logout = async() => {
    const res = await Handler.request("POST", "/api/v1/logout")
    return res
}

const userHandler = {
    get,
    create,
    login,
    logout
}

export default userHandler