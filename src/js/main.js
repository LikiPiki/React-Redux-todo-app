import React from 'react';
import ReactDOM from 'react-dom';

import {Provider} from 'react-redux';
import {createStore, combineReducers} from 'redux';

import App from './components/App.js';


import "../scss/main.scss";


let inittialState = {
	'hello': true,
	todos: [] 
}

let data;

var xhrGet = new XMLHttpRequest();
xhrGet.open('GET', 'get_all', false);
xhrGet.send();

if (xhrGet.status != 200) {
  console.log( xhrGet.status + ': ' + xhrGet.statusText );
} else {
  data =  JSON.parse(xhrGet.responseText);
}
inittialState.todos = data;


// TODO: реализовать удаление по id елемента!!!


let reducer = (state=inittialState, action) => {
	switch (action.type) {
		case 'ADD': {
			let xhrPost = new XMLHttpRequest();
			xhrPost.open('POST', 'add_new', true)		// });
			let data = {
				"name": action.payload.name
			}
			console.log(data);
			xhrPost.send(JSON.stringify(data));

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

