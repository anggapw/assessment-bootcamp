import { BrowserRouter as Router, Switch, Route } from 'react-router-dom'

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
				<Route path="/dashboard" component={Dashboard} />
				<Route path="/add" component={Add} />
				<Route path="/" component={Homepage} />
			</Switch>
		</Router>
	)
}

export default Routes
