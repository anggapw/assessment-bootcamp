import axios from 'axios'

const passwordAPI = axios.create({
	baseURL: 'http://localhost:8080'
})

export default passwordAPI
