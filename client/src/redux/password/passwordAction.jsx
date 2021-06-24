import passwordAPI from '../../api/passwordAPI'
import Swal from 'sweetalert2'

import { SET_WEB_PASSWORD, SET_WEBSITE, SET_PASSWORD } from './passwordActionTypes'

const setWebsite = (website) => {
	return {
		type: SET_WEBSITE,
		payload: {
			website: website
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

const setWebPassword = (passwordData) => {
	return {
		type: SET_WEB_PASSWORD,
		payload: {
			data: passwordData
		}
	}
}

const addNewPassword = (website, password, history) => async (dispatch) => {
	try {
		const data = {
			website: website,
			title: password
		}

		const postData = await passwordAPI({
			method: 'POST',
			url: '/password',
			data: data,
			headers: {
				Authorization: localStorage.getItem('token')
			}

		})

		if (postData.status === 200) {
			Swal.fire({ title: 'Success Add New Password', icon: 'success', timer: 1500 })
		}

		history.push('/')
	} catch (err) {
		console.log(err.response)
	}
}

const deletePassword = (id, history) => async (dispatch) => {
	try {
		// eslint-disable-next-line
		const postData = await passwordAPI({
			method: 'DELETE',
			url: `/password/${id}`,
			headers: {
				Authorization: localStorage.getItem('token')
			}
		})

		if (postData.status === 200) {
			Swal.fire({ title: 'Success Remove', icon: 'success', timer: 1500 })

			history.push('/dashboard')
		}
	} catch (err) {
		console.log(err.response)
	}
}

const passwordAction = {
	setWebPassword,
	setWebsite,
	setPassword,
	addNewPassword,
	deletePassword
}

export default passwordAction
