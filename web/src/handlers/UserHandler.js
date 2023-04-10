import axios from "axios"
import Handler from "./Handler"

// not sure if it makes sense to have this as a class or just
// a colletion of functions.
//
// All handlers should return a tuple w/ whether the
// request was successful and the body/data/message
export default class UserHandler extends Handler {
    getUser = async() => {
        try {
            const res = await axios.get("/api/v1/user")
            return res
        } catch(err) {
            return err
        }
    }

    // should return [status, body]
    createUser = async(username, email, password) => {
        const body = {
            username: username,
            email: email,
            password: password
        }

        const res = await this.request("POST", "/api/v1/user", body)
        return res
    }
}