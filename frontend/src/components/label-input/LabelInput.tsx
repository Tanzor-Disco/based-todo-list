import "./LabelInput.css"

interface LabelInputProps {
	name:string;
	type:string;
	required:boolean;
}

export default function LabelInput({name,type,required}:LabelInputProps) {
	return (
		<div className="label-input">
			<label htmlFor={name}>{name}</label>
			<input type={type} name={name} required={required} />
		</div>
	)
}
