import React, { useState, useEffect } from 'react'
import { Table, Container, Button, Modal, Form, Row } from 'react-bootstrap'
import { useSelector, useDispatch } from 'react-redux'
import { useHistory } from 'react-router-dom'

import passwordAction from '../redux/password/passwordAction'

const ResultTable = () => {
	const [showModal, setShowModal] = useState(false)
	const dispatch = useDispatch()
	const history = useHistory()
	const webPasswordData = useSelector((state) => state.webPassword)

	useEffect(() => {
		dispatch(passwordAction.setWebPassword())
	}, [dispatch])

	const handleClose = () => setShowModal(false)

	const handleShow = (e) => {
		e.preventDefault()

		setShowModal(true)
	}

	const handleAdd = (e) => {
		e.preventDefault()

		history.push('/add')
	}

	return (
		<Container className="mt-3">
			<Row className="mb-3">
				<Button variant="success" onClick={handleAdd}>
					Add
				</Button>
			</Row>
			<Modal show={showModal} onHide={handleClose}>
				<Modal.Header>
					<Modal.Title>Update</Modal.Title>
				</Modal.Header>
				<Modal.Body>
					<Form>
						<Form.Group controlId="formUpdateemail">
							<Form.Label>Email address</Form.Label>
							<Form.Control type="text" placeholder="Enter email" />
						</Form.Group>

						<Form.Group controlId="formUpdatePassword">
							<Form.Label>Password</Form.Label>
							<Form.Control type="text" placeholder="Password" />
						</Form.Group>
					</Form>
				</Modal.Body>
				<Modal.Footer>
					<Button variant="secondary" onClick={handleClose}>
						Close
					</Button>
					<Button variant="primary" onClick={handleClose}>
						Save Changes
					</Button>
				</Modal.Footer>
			</Modal>

			<Table striped bordered hover>
				<thead>
					<tr>
						<th>No</th>
						<th>Website</th>
						<th>Password</th>
						<th>Created At</th>
						<th>Updated At</th>
						<th>Action</th>
					</tr>
				</thead>
				{webPasswordData &&
					webPasswordData.map((value, index) => {
						return (
							<tbody>
								<tr>
									<td>{value.ID}</td>
									<td>{value.website}</td>
									<td>{value.title}</td>
									<td>{value.CreatedAt}</td>
									<td>{value.UpdatedAt}</td>
									<td>
										<Button variant="primary" onClick={handleShow}>
											Update
										</Button>
										<Button
											variant="danger"
											onClick={() => dispatch(passwordAction.deletePassword(value.ID, history))}
										>
											Delete
										</Button>
									</td>
								</tr>
							</tbody>
						)
					})}
			</Table>
		</Container>
	)
}

export default ResultTable
