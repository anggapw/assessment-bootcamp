import { SET_EMAIL, SET_PASSWORD, SET_ADDRESS, SET_FULLNAME } from '../userActionTypes'

const initialState = {
	email: '',
	password: '',
	address: '',
	fullName: ''
}

const userRegisterReducer = (state = initialState, action) => {
	switch (action.type) {
		case SET_EMAIL:
			return {
				...state,
				email: action.payload.email
			}
		case SET_PASSWORD:
			return {
				...state,
				password: action.payload.password
			}
		case SET_FULLNAME:
			return {
				...state,
				fullName: action.payload.fullName
			}
		case SET_ADDRESS:
			return {
				...state,
				address: action.payload.address
			}
		default:
			return state
	}
}

export default userRegisterReducer
