import axios from "axios"


// not sure if it makes sense to have this as a class or just
// a colletion of functions.
export default class UserHandler {
    getUser = async() => {
        try {
            const res = await axios.get("/api/v1/user")
            return res
        } catch(err) {
            return err
        }
    }

    createUser = async(username, email, password) => {
        const body = {
            username: username,
            email: email,
            password: password
        }

        try {
            const res = await axios.post("/api/v1/user", body)
            return res
        } catch(err) {
            return err
        }
    }
}