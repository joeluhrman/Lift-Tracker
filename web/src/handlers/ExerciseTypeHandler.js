import Handler from "./Handler"

export default class ExerciseTypeHandler extends Handler {
    async getAll() {
        const res = await this.request("GET", "/api/v1/exercise-type")
        return res
    }
}