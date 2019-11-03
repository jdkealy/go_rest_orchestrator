package templates

var JsListTemplate = `
import React from "react";
import {{.PluralLowerModel}} from "../../models/{{.Model}}"
import fields from './fields'
import Table from '@material-ui/core/Table';
import TableHeadCmp from '../../components/table_head'
import TableBodyCmp from '../../components/table_body'
import { Provider, Subscribe } from 'unstated';
import ItemCreate from './new'

class ItemsTable extends React.Component {
	componentDidMount(){
		{{.PluralLowerModel}}.search();
	}
	render(){
		return (
			<Table>
				<TableHeadCmp fields={fields.listFields} />
				<TableBodyCmp items={ {{.PluralLowerModel}}.state.items } fields={fields.listFields} />
			</Table>
		)
	}
}

export default function {{.PluralModel}}() {
	return (
		<div>
			<ItemCreate />
			<Provider>
				<Subscribe to={[{{.PluralLowerModel}}]}>
					{ 
					{{.PluralLowerModel}} => (
						<ItemsTable />
					)}
				</Subscribe>
			</Provider>
		</div>
	);
}
`