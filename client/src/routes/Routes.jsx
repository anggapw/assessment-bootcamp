import { BrowserRouter as Router, Switch, Route } from 'react-router-dom'


import PrivateRoute from './PrivateRoute'

//components
import NavbarMenu from '../components/NavbarMenu'

///pages
import Login from '../pages/Login'
import Register from '../pages/Register'
import Homepage from '../pages/Homepage'
import Dashboard from '../pages/Dashboard'
import Add from '../pages/Add'

const Routes = () => {
	return (
		<Router>
			<NavbarMenu />
			<Switch>
				<Route path="/login" component={Login} />
				<Route path="/register" component={Register} />
				<PrivateRoute path="/dashboard">
					<Dashboard />
				</PrivateRoute>
				<PrivateRoute path="/add">
					<Add />
				</PrivateRoute>
				<Route path="/" component={Homepage} />
			</Switch>
		</Router>
	)
}

export default Routes
