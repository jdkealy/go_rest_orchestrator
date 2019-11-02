package templates

var JsLFormTemplate = `
import React, { useState } from "react";
import {{.PluralLowerModel}} from "../../models/{{.Model}}"
import fields from './fields'
import { makeStyles } from '@material-ui/core/styles';
import TextField from '@material-ui/core/TextField';
import Button from '@material-ui/core/Button';

const useStyles = makeStyles(theme => ({
	container: {
		display: 'flex',
		flexWrap: 'wrap',
	},
	fullWidth: {
		width: '100%'
	},
	textField: {
		marginLeft: theme.spacing(1),
		width: '300px',
		marginRight: theme.spacing(1),
	},
	dense: {
		marginTop: theme.spacing(2),
	},
	menu: {
		width: 200,
	},
}));



export default function ItemCreate(props) {
	var initialState = {}
	fields.newFormFields.map((field)=>{return (initialState[field.field] = "");})
	const [node, setNode] = useState(initialState);

	const classes = useStyles();

	const onChange = function(e, item){
		node[item.field] = e.target.value;
		setNode(node);
	}

	const onSubmit = function(e){
		{{.PluralLowerModel}}.create(node)
		e.preventDefault();
	}

	return (
		<form onSubmit={onSubmit} className={classes.fullWidth}>
			<div>
				{fields.newFormFields.map(function(item, index){
					return (
						<div key={item.field}>
						<TextField
							id="filled-name"
							label={item.label}
							onChange={(e) => {onChange(e, item)}}
					        className={classes.textField}
							margin="normal"
						/>
						</div>
					);
				})}
			</div>
			<div>
				<Button type="submit">SUBMIT</Button>
			</div>
		</form>
	)
}
`