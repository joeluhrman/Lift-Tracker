import axios from "axios"

const request = async(method, endpoint, data) => {
    var reqFunc
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
        const res = error.response
        return [res.status, res.headers, res.data]
    }
}

const handler = {
    request,
} 

export default handler