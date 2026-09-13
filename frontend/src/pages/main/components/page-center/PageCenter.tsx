import "./PageCenter.css"
import TaskList from "./components/task-list/TaskList"
import TaskInput from "./components/task-input/TaskInput"
import type TaskData from "@/models/TaskData"
import {useState,useEffect} from "react"

export default function PageCenter() {
	const [tasks,setTasks] = useState<TaskData[]>([])

	const test = {
		ID:123,
		Description:"123"
	}

	useEffect(() => {
		setTasks([test])
	},[]) 
	

	return (
		<section className="page-center">
			<div className="tasks-container">
				<h1>Tasks:</h1>
				<TaskList Tasks={tasks}/>
				<TaskInput setTasks={setTasks} lenTasks={tasks.length}/>
			</div>
		</section>
	)
}
