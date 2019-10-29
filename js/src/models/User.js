import { Container } from 'unstated';

const axios = require('axios');

class CounterContainer extends Container<CounterState> {
	state = {
		items: []
	};

	create(data) {
		var cmp = this;
		axios.post('/users', data)
			.then(function (response) {
				cmp.search();
			})
			.catch(function (error) {
				console.log(error);
			});
	}

	search() {
		var cmp = this;
		axios.get('/users')
			.then(function (response) {
				cmp.setState({items: response.data});
			})
			.catch(function (error) {
				console.log(error);
			});
	}

	update() {
		this.setState({ count: this.state.count - 1 });
	}

	destroy() {
		this.setState({ count: this.state.count - 1 });
	}

}

const groups = new CounterContainer();
export default groups