import passwordAPI from '../../../api/passwordAPI'
import Swal from 'sweetalert2'

import { SET_EMAIL, SET_PASSWORD, SET_AUTH, SET_USER} from '../userActionTypes'
import passwordAction from '../../password/passwordAction'

const setEmail = (email) => {
	return {
		type: SET_EMAIL,
		payload: {
			email: email
		}
	}
}

const setPassword = (password) => ({
	type: SET_PASSWORD,
	payload: {
		password: password
	}
})

const setAuth = (isAuth) => ({
	type: SET_AUTH,
	payload: {
		isAuth: isAuth
	}
})

const setUser = (user) => ({
	type: SET_USER,
	payload: {
		user: user
	}
})

const login = (email, password, history, location) => async (dispatch) => {
	try {
		const loginData = {
			email: email,
			password: password
		}

		const postData = await passwordAPI({
			method: 'POST',
			url: '/users/login',
			data: loginData
		})

		const getData = await passwordAPI({
			method: 'GET',
			url: '/password',
			headers: {
				Authorization: postData.data.token
			}
		})

		console.log(getData)

		dispatch(passwordAction.setWebPassword(getData.data))

		if (postData.status === 200) {
			Swal.fire({
				title: 'Success Login',
				icon: 'success',
				timer: 2000
			})
		}

		let token = postData.data.token
		localStorage.setItem('token', token)

		dispatch(setUser(postData.data))
		localStorage.setItem('user', postData.data)

		let { from } = location.state || { from: { pathname: '/dashboard' } }

		history.replace(from)

	} catch (error) {
		console.log(error);
		let code = error.response.status

		if (code === 401) {
			Swal.fire({
				title: 'Email/Password Invalid',
				icon: 'warning',
				timer: 2000
			})
		}
	}
}

const userLoginAction = {
	setEmail,
	setPassword,
	setAuth,
	login
}

export default userLoginAction
