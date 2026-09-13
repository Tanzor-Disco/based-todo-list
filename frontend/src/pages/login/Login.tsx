import "./Login.css"
import LabelInput from "@/components/label-input/LabelInput"
import logo from "@/assets/logo.svg"
import { useNavigate } from "react-router-dom"


export default function Login() {
	const navigate = useNavigate()
	async function handleSubmit(formData:FormData) {
		const formDataObj = Object.fromEntries(formData)
		const response = await fetch("api/login", {
			method:"POST",
			headers: {
				"Content-Type": "application/json",
			},
			body:JSON.stringify(formDataObj)
		})

		if (response.ok) {
			navigate("/")
		}
	}
	return (
		<main className="login-page-center">
			<img src={logo} className="login-site-logo"/>
			<div className="login-center-container">
				<form action={handleSubmit} className="login-form">
					<LabelInput name="Username" type="text" required={true} />
					<LabelInput name="Password" type="text" required={true} />
					<button className="login-submit">Submit</button>
					<a href="/register/"> Sign Up </a>
				</form>
			</div>
		</main>
	)
}
