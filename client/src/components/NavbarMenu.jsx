import React from 'react'
import { Navbar, Nav, Form, Button } from 'react-bootstrap'
import { Link, useHistory } from 'react-router-dom'


const NavbarMenu = () => {
	let token = localStorage.getItem('token')
	const history = useHistory()

	const handleLogout = () => {
		localStorage.clear();
		history.push('/')
	  };
	
	return (
		<div>
			<Navbar bg="dark" variant="dark">
				<Navbar.Brand as={Link} to="/">
					Logo
				</Navbar.Brand>
				<Nav className="mr-auto">
					<Nav.Link as={Link} to="/">
						Home
					</Nav.Link>
				</Nav>
				{token ? (
					<Form inline>
						<Nav.Link as={Link} to="/dashboard">
							Dashboard
						</Nav.Link>
					<Button onClick={handleLogout}>
						Logout
					</Button>
					</Form>
					) : (
					<Form inline>
						<Nav.Link as={Link} to="/login">
							Login
						</Nav.Link>
						<Nav.Link as={Link} to="/register">
							Register
						</Nav.Link>
					</Form>
					)}
			</Navbar>
		</div>
	)
}

export default NavbarMenu
