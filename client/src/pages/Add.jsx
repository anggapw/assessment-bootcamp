import React from 'react'
import { Container, Form, Button } from 'react-bootstrap'
import { useSelector, useDispatch } from 'react-redux'
import { useHistory } from 'react-router-dom'

import passwordAction from '../redux/password/passwordAction'

const Add = () => {
	const addNewPasswordData = useSelector((state) => state.webPassword)
	const dispatch = useDispatch()
	const history = useHistory()

	const handleSubmit = (e) => {
		e.preventDefault()

		dispatch(passwordAction.addNewPassword(addNewPasswordData.website, addNewPasswordData.password, history))
	}

	return (
		<Container className="mt-4">
			<Form onSubmit={handleSubmit}>
				<Form.Group controlId="formBasicWebsite">
					<Form.Label>Website</Form.Label>
					<Form.Control
						type="text"
						placeholder="Website"
						onChange={(e) => dispatch(passwordAction.setWebsite(e.target.value))}
					/>
				</Form.Group>

				<Form.Group controlId="formBasicPassword">
					<Form.Label>Password</Form.Label>
					<Form.Control
						type="password"
						placeholder="Password"
						onChange={(e) => dispatch(passwordAction.setPassword(e.target.value))}
					/>
				</Form.Group>
				<Form.Group controlId="formBasicCheckbox"></Form.Group>
				<Button variant="primary" type="submit">
					Submit
				</Button>
			</Form>
		</Container>
	)
}

export default Add
