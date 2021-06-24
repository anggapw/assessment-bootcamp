import React, {useEffect} from 'react'
import { useSelector, useDispatch } from 'react-redux'

import userLoginAction from './redux/user/login/userLoginAction'

import './App.css'

import Routes from './routes/Routes'

function App() {
	const loginData = useSelector((state) => state.userLogin)
	const dispatch = useDispatch()

	useEffect(() => {
		const loggedInUser = localStorage.getItem("auth")
		if (loggedInUser) {
		  const foundUser = JSON.parse(loggedInUser)
		  dispatch(userLoginAction.setUser(foundUser))
		}
	  }, []);
	return (
		<>
			<Routes />
		</>
	)
}

export default App
