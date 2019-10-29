import users from "../../models/User"
import React from "react";
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import { Provider, Subscribe } from 'unstated';
import ItemCreate from './new'

class ItemsTable extends React.Component {
	componentDidMount(){
		users.search();
	}
	render(){
		return (
			<Table>
				<TableHead>
        			<TableRow><TableCell>ID</TableCell><TableCell>Name</TableCell><TableCell>DOB</TableCell></TableRow>
				</TableHead>
				<TableBody>
					{users.state.items.map(function(item, index){
	                	return (
	                		<TableRow key={ index }>
	                			<TableCell>{item.id}</TableCell>
	                			<TableCell>{item.name}</TableCell>
	                			<TableCell>{item.dob}</TableCell>
                			</TableRow>
	            		);
	              	})}
				</TableBody>            	
			</Table>
		)
	}
}

export default function Users() {
	return (
		<div>
			<ItemCreate />
			<Provider>
				<Subscribe to={[users]}>
					{users => (
						<ItemsTable />
					)}
				</Subscribe>
			</Provider>
		</div>
	);
}