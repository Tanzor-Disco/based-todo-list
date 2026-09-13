import "./TaskInput.css"
import type TaskData from "@/models/TaskData"
import type { Dispatch, SetStateAction } from "react"

interface TaskInputProps {
	setTasks:Dispatch<SetStateAction<TaskData[]>>;
}

export default function TaskInput({setTasks}:TaskInputProps) {
	async function handleSubmit(formData: FormData) {
		const formDataObj = Object.fromEntries(formData)
		const response = await fetch("api/main/task/create",{
			method:"POST",
			headers: {
				"Content-Type": "application/json",
    		},
			body:JSON.stringify(formDataObj)
		})
		
		if (!response.ok) {
			console.warn("The response from the server is not OK")
			return
		}

		const taskID = await response.json()
		
		const newTask = {
			TaskID:taskID as number,
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
