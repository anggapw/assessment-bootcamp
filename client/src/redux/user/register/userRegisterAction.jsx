import passwordAPI from '../../../api/passwordAPI'
import Swal from 'sweetalert2'

import { SET_EMAIL, SET_PASSWORD, SET_FULLNAME, SET_ADDRESS } from '../userActionTypes'

const setEmail = (email) => {
	return {
		type: SET_EMAIL,
		payload: {
			email: email
		}
	}
}

const setPassword = (password) => {
	return {
		type: SET_PASSWORD,
		payload: {
			password: password
		}
	}
}

const setFullName = (fullName) => {
	return {
		type: SET_FULLNAME,
		payload: {
			fullName: fullName
		}
	}
}

const setAddress = (address) => {
	return {
		type: SET_ADDRESS,
		payload: {
			address: address
		}
	}
}

const register = (email, password, fullName, address, history) => async (dispatch) => {
	try {
		const data = {
			email: email,
			password: password,
			full_name: fullName,
			address: address
		}

		// eslint-disable-next-line
		const postData = await passwordAPI({
			method: 'POST',
			url: '/users/register',
			data: data
		})

		console.log(postData)

		if (postData.status === 201) {
			Swal.fire({ title: 'Register Success', icon: 'success', timer: 1500 })

			history.push('/login')
		}
	} catch (err) {
		console.log(err.response)

		if (err.response.status === 500) {
			Swal.fire({ title: 'Email sudah pernah didaftarkan', icon: 'warning', timer: 1500 })
		}
	}
}

const userRegisterAction = {
	setEmail,
	setPassword,
	setFullName,
	setAddress,
	register
}

export default userRegisterAction
