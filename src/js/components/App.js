import React from 'react';

import {Component} from 'react';
import {connect} from 'react-redux';
import { bindActionCreators } from 'redux';
import * as actions from '../actions.js';

import Element from './Element.js'


class App extends Component {
	constructor(props) {
		super(props);
		this.addnew = this.addnew.bind(this);
	}

	addnew(event) {
		console.log(this.refs.input.value);
		this.props.actions.action(this.refs.input.value);
	}

	render() {
		return (
			<div>
				<h1>Header here</h1>
				<input type="text" ref="input"/>
				<button onClick={this.addnew}>Add new</button>
				<div>
				{this.props.todos.map(item => 
					<Element key={item.id} content={item}/>
				)}
					
				</div>
			</div>
		)
	}
}

function mapStateToProps(state) {
	return {
		todos: state.todos
	}
}

function mapDispatchToProps(dispatch) {
	return {
		actions: bindActionCreators(actions, dispatch)
	}
}

export default connect(mapStateToProps, mapDispatchToProps)(App);
