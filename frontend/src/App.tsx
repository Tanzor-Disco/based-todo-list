import Main from "./pages/main/Main"
import Register from "./pages/register/Register"
import Login from "./pages/login/Login"
import { BrowserRouter, Routes, Route } from 'react-router'
import './App.css'


export default function App() {
	return (
        <BrowserRouter>
            <Routes>
                <Route path="/" element={<Main />} />
				<Route path="/register" element={<Register />} />
				<Route path="/login" element={<Login />}/>
            </Routes>
        </BrowserRouter>
    )
}
