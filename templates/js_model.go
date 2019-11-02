package templates

var JsModelTemplate = `
import { Container } from 'unstated';

const axios = require('axios');

class CounterContainer extends Container<CounterState> {
	state = {
		items: []
	};

	create(data) {
		var cmp = this;
		axios.post('/{{.PluralLowerModel}}', data)
			.then(function (response) {
				cmp.search();
			})
			.catch(function (error) {
				console.log(error);
			});
	}

	search() {
		var cmp = this;
		axios.get('/{{.PluralLowerModel}}')
			.then(function (response) {
				cmp.setState({items: response.data});
			})
			.catch(function (error) {
				console.log(error);
			});
	}
}

const groups = new CounterContainer();
export default groups
`