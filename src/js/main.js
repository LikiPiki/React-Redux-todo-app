import React from 'react';
import ReactDOM from 'react-dom';

import {Provider} from 'react-redux';
import {createStore, combineReducers} from 'redux';

import App from './components/App.js';
// import Element from './components/Element.js';


const inittialState = {
	'hello': true,
	todos: [
		{
			name: 'kek',
			checked: false,
			id: 0
		}
	] 
}

let reducer = (state=inittialState, action) => {
	switch (action.type) {
		case 'ADD': {
			console.log('addded', action.payload.name);
			return {
				...state,
				todos: [...state.todos, {
					name: action.payload.name,
					checked: false,
					id: state.todos.length
				}]
			}
		}
		case 'REMOVE': {
			return {
				...state,
				todos: state.todos.map(
					(element, i) => i === action.payload.id ? {...element, checked: true} : element)
			}
		}
		default: {
			return state;
		}
	}
}


const store = createStore(reducer)


ReactDOM.render(
	<Provider store={store}>
		<App/>
	</Provider>,
	document.getElementById('app'));

