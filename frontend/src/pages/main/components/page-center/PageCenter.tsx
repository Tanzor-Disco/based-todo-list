import "./PageCenter.css"
import TaskList from "./components/task-list/TaskList"
import TaskInput from "./components/task-input/TaskInput"
import type TaskData from "@/models/TaskData"
import {useState,useEffect} from "react"
import { useNavigate } from "react-router-dom"

export default function PageCenter() {
	const navigate = useNavigate()
	const [tasks,setTasks] = useState<TaskData[]>([])


	useEffect(() => {
		async function fetchTasks() {
			const response = await fetch("/api/main/tasks")
			if (!response.ok) {
				console.warn("Code of the response is not OK")
				navigate("/login")
				return
			}
			const respObj = await response.json()
			const tasks = respObj.data
			setTasks(tasks)
		}
		fetchTasks()
	},[]) 
	

	return (
		<section className="tasks-container">
			<h1>Tasks:</h1>
			<TaskList Tasks={tasks}/>
			<TaskInput setTasks={setTasks} />
		</section>
	)
}
