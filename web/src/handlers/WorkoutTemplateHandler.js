import axios from "axios"
import Handler from "./Handler"

export default class WorkoutTemplateHandler extends Handler {
    async getAll() {
        const res = await this.request("GET", "/api/v1/workout-template")
        return res
    }
}