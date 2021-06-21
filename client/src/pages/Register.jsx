import React from 'react'
import { Container, Form, Button } from 'react-bootstrap'
import { useSelector, useDispatch } from 'react-redux'
import { useHistory } from 'react-router-dom'

import userRegisterAction from '../redux/user/register/userRegisterAction'

const Register = () => {
	const registerData = useSelector((state) => state.userRegister)
	const dispatch = useDispatch()
	const history = useHistory()

	const handleSubmit = (e) => {
		e.preventDefault()

		dispatch(
			userRegisterAction.register(
				registerData.email,
				registerData.password,
				registerData.fullName,
				registerData.address,
				history
			)
		)
	}
	return (
		<Container>
			<Form onSubmit={handleSubmit}>
				<Form.Group controlId="formFullName">
					<Form.Label>Full Name</Form.Label>
					<Form.Control
						type="text"
						placeholder="Enter your name"
						onChange={(e) => dispatch(userRegisterAction.setFullName(e.target.value))}
					/>
				</Form.Group>

				<Form.Group controlId="formAddress">
					<Form.Label>Address</Form.Label>
					<Form.Control
						type="text"
						placeholder="Enter your address"
						onChange={(e) => dispatch(userRegisterAction.setAddress(e.target.value))}
					/>
				</Form.Group>

				<Form.Group controlId="formBasicEmail">
					<Form.Label>Email address</Form.Label>
					<Form.Control
						type="email"
						placeholder="Enter email"
						onChange={(e) => dispatch(userRegisterAction.setEmail(e.target.value))}
					/>
				</Form.Group>

				<Form.Group controlId="formBasicPassword">
					<Form.Label>Password</Form.Label>
					<Form.Control
						type="password"
						placeholder="Password"
						onChange={(e) => dispatch(userRegisterAction.setPassword(e.target.value))}
					/>
				</Form.Group>
				<Button variant="primary" type="submit">
					Submit
				</Button>
			</Form>
		</Container>
	)
}

export default Register
