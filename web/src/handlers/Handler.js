import axios from "axios"

export default class Handler {
    // returns status code, headers, and body of response
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
            return [res.status, res.headers, res.data]
        } catch(error) {
           // if (error.response) {
                const res = error.response
                return [res.status, res.headers, res.data]
           // }
        }
    }
}