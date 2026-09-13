import "./Login.css"
import LabelInput from "@/components/label-input/LabelInput"
import logo from "@/assets/logo.svg"


export default function Login() {
	async function handleSubmit() {

	}
	return (
		<main className="login-page-center">
			<img src={logo} className="login-site-logo"/>
			<div className="login-center-container">
				<form action={handleSubmit} className="login-form">
					<LabelInput name="Username" type="text" required={true} />
					<LabelInput name="Password" type="text" required={true} />
					<button className="login-submit">Submit</button>
				</form>
			</div>
		</main>
	)
}
