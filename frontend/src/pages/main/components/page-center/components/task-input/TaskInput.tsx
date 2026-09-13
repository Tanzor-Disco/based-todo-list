import "./TaskInput.css"
import type TaskData from "@/models/TaskData"
import type { Dispatch, SetStateAction } from "react"

interface TaskInputProps {
	setTasks:Dispatch<SetStateAction<TaskData[]>>;
	lenTasks: number;
}

export default function TaskInput({setTasks,lenTasks}:TaskInputProps) {
	async function handleSubmit(formData: FormData) {
		const formDataObj = Object.fromEntries(formData)
		const newTask = {
			ID:lenTasks,
			Description:formDataObj.description as string
		}
		setTasks((prevTasks) => [...prevTasks,newTask])

		
	}
	return (
		<form action={handleSubmit}>
			<input className="task-input" placeholder="Add a task" name="description"></input>
		</form>
	)
}
