import "./Task.css"

interface TaskProps {
	Description: String;
}

export default function Task({Description}:TaskProps) {
	return (
		<div className="task">
			<p>● {Description}</p>
		</div>
	)
}
