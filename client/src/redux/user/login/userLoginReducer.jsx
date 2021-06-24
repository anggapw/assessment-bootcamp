import { SET_PASSWORD, SET_EMAIL, SET_AUTH, SET_USER } from '../userActionTypes'

const initialState = {
	email: '',
	password: '',
	isAuth: false,
	user: ''
}

const userLoginReducer = (state = initialState, action) => {
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
		case SET_AUTH:
			return {
				...state,
				isAuth: action.payload.isAuth
			}
		case SET_USER:
		return {
			...state,
			isAuth: action.payload.user
		}
		default:
			return state
	}
}

export default userLoginReducer
