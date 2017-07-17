
export function action(text) {
	return {
		type: 'ADD',
		payload: {
			name: text,
			checked: false
		}
	}
}

export function remove(id) {
	return {
		type: "REMOVE",
		payload: {
			id
		}
	}
}


