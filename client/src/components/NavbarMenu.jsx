import React from 'react'
import { Navbar, Nav, Form } from 'react-bootstrap'
import { Link } from 'react-router-dom'

const NavbarMenu = () => {
	return (
		<div>
			<Navbar bg="dark" variant="dark">
				<Navbar.Brand as={Link} to="/">
					Navbar
				</Navbar.Brand>
				<Nav className="mr-auto">
					<Nav.Link as={Link} to="/">
						Home
					</Nav.Link>
					<Nav.Link as={Link} to="/dashboard">
						Dashboard
					</Nav.Link>
					{/* <Nav.Link href="#features">Features</Nav.Link>
					<Nav.Link href="#pricing">Pricing</Nav.Link> */}
				</Nav>
				<Form inline>
					<Nav.Link as={Link} to="/login">
						Login
					</Nav.Link>
					<Nav.Link as={Link} to="/register">
						Register
					</Nav.Link>
				</Form>
			</Navbar>
		</div>
	)
}

export default NavbarMenu
