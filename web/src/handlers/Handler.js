import axios from "axios"

export default class Handler {
    async request(method, endpoint, data) {
        var reqFunc = undefined
        switch(method) {
            case "GET":
                reqFunc = axios.get
                break
            case "POST":
                reqFunc = axios.post
                break
            case "PUT":
                reqFunc = axios.put
                break
            case "PATCH":
                reqFunc = axios.patch
                break
            case "DELETE":
                reqFunc = axios.delete
                break
        }

        try {
            const res = await reqFunc(endpoint, data)
            return [true, res.data]
        } catch(error) {
            if (error.response) {
                if (error.response.data !== undefined) {
                    return [false, "The server is not responding."]
                }
                return [false, err.response.data]
            } else if (error.request) {
                return [false, "Check your internet connection."]
            } else {
                return [false, "An unknown error occurred."]
            }
        }
    }
}