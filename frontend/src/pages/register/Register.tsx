import "./Register.css"
import LabelInput from "@/components/label-input/LabelInput"
import logo from "@/assets/logo.svg"

export default function Register() {
	async function handleSubmit(formData:FormData) {
		
	}
	return (
		<main className="register-page-center">
			<img src={logo} className="register-site-logo"/>
			<div className="register-center-container">
				<form action={handleSubmit} className="register-form">
					<LabelInput name="Username" type="text" required={true} />
					<LabelInput name="Password" type="text" required={true} />
					<button className="register-submit">Submit</button>
				</form>
			</div>
		</main>
	)
}
