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
		this.enterPress = this.enterPress.bind(this);
	}

	enterPress(event) {
		if (event.which == 13) {
			this.addnew();
		}
	}

	addnew(event) {
		if (this.refs.input.value) {
			this.props.actions.action(this.refs.input.value);
			this.refs.input.value = '';
		}
	}

	render() {
		return (
			<div>
				<div className="header">
					<div className="content">
						<input type="text" onKeyPress={this.enterPress} ref="input"/>
						<button onClick={this.addnew}>+</button>
					</div>
				</div>
				<ul className="todo">
					{this.props.todos.map(item => 
						<Element key={item.id} content={item}/>
					)}
				</ul>
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
