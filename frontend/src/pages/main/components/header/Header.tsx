import logo from "@/assets/logo.svg"
import "./Header.css"

export default function Header() {
	return (
		<header className="site-header">
			<img src={logo} className="site-logo"/>
			<h1>Based To-do List</h1>
		</header>
	)
}
