import React from 'react'
import { Form, Button, Container } from 'react-bootstrap'
import { useHistory, useLocation } from 'react-router-dom'
import { useSelector, useDispatch } from 'react-redux'

import userLoginAction from '../redux/user/login/userLoginAction'

const Login = () => {
	const loginData = useSelector((state) => state.userLogin)
	const dispatch = useDispatch()
	const history = useHistory()
	let location = useLocation()

	const loginHandler = (e) => {
		e.preventDefault()

		dispatch(userLoginAction.login(loginData.email, loginData.password, history, location))
	}

	return (
		<Container>
			<Form onSubmit={loginHandler}>
				<Form.Group controlId="formBasicEmail">
					<Form.Label>Email address</Form.Label>
					<Form.Control type="email" placeholder="Enter email" onChange={(e) => {
											dispatch(userLoginAction.setEmail(e.target.value))
										}}/>
				</Form.Group>

				<Form.Group controlId="formBasicPassword">
					<Form.Label>Password</Form.Label>
					<Form.Control type="password" placeholder="Password" onChange={(e) => {
											dispatch(userLoginAction.setPassword(e.target.value))
										}}/>
				</Form.Group>
				<Form.Group controlId="formBasicCheckbox"></Form.Group>
				<Button variant="primary" type="submit">
					Submit
				</Button>
			</Form>
		</Container>
	)
}

export default Login
