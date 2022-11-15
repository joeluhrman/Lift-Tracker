import { SubmitLoginForm } from "../../../event-handlers/login/login-handlers"

const LoginForm = () => {
    return (
        <section>
            <h1>Login</h1>

            <form>
                <input type="text" name="username" placeholder="Username" />
                <input type="password" name="password" placeholder="Password" />

                <input type="submit" value="Login" onClick={SubmitLoginForm} />
            </form>
        </section>
    )
}

export default LoginForm