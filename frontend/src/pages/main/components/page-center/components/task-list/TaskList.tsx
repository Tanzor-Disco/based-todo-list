import "./TaskList.css"
import Task from "./components/task/Task"
import type TaskData from "@/models/TaskData"

interface TasksListProps {
	Tasks: TaskData[];
}

export default function TaskList({Tasks}:TasksListProps) {
	const taskComponents = Tasks.map((task) => {
		return <Task key = {task.TaskID} Description={task.Description} /> 
	})
	return (
		taskComponents
	)
}
