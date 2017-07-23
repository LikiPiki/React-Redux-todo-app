import React from 'react';

import {Component} from 'react';
import {connect} from 'react-redux';

import { bindActionCreators } from 'redux';
import * as actions from '../actions';

class Element extends Component {
	constructor(props) {
		super(props);
		this.delete = this.delete.bind(this);
	}


	delete(event) {
		console.log(this.props.content.id);
		this.props.actions.remove(
			this.props.content.id
		);
	}

	render() {
		return (
			<li>
				<div className="complete-button"></div>
				{this.props.content.checked == false ? this.props.content.name: "DELETED"}
				<div onClick={this.delete} className="remove-button"></div>
			</li>
		);
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


export default connect(mapStateToProps ,mapDispatchToProps)(Element);