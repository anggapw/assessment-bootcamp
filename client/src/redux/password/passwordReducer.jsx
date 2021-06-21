import { SET_WEB_PASSWORD, SET_WEBSITE, SET_PASSWORD } from './passwordActionTypes'

const initialState = []

const userRegisterReducer = (state = initialState, action) => {
	switch (action.type) {
		case SET_WEB_PASSWORD:
			return action.payload.data
		case SET_WEBSITE:
			return {
				...state,
				website: action.payload.website
			}
		case SET_PASSWORD:
			return {
				...state,
				password: action.payload.password
			}
		default:
			return state
	}
}

export default userRegisterReducer
