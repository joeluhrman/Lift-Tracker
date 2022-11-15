import { API_V1_POST_LOGIN } from "../../endpoints"

export const SubmitLoginForm = async () => {
    fetch(API_V1_POST_LOGIN, {
        method: "POST"
    })
    .then((response) => {
        return response.code 
    })
    .catch((error) => {
        console.log(error)
        return error
    })
}