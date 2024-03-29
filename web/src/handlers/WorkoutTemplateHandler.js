import Handler from "./Handler"

export default class WorkoutTemplateHandler extends Handler {
    async getAll() {
        const res = await this.request("GET", "/api/v1/workout-template")
        return res
    }

    async create(workoutTemplate) {
        const res = await this.request("POST", "/api/v1/workout-template", workoutTemplate)
        return res
    }
}