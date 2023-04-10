import Handler from "./Handler"

export default class UserHandler extends Handler {
    get = async() => {
        const res = await this.request("GET", "/api/v1/user")
        return res
    }

    create = async(username, email, password) => {
        const body = {
            username: username,
            email: email,
            password: password,
        }

        const res = await this.request("POST", "/api/v1/user", body)
        return res
    }

    login = async(username, password) => {
        const body = {
            username: username,
            password: password,
        }

        const res = await this.request("POST", "/api/v1/login", body)
        return res
    }

    logout = async() => {
        const res = await this.request("POST", "/api/v1/logout")
        return res
    }
}