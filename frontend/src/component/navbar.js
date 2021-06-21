import React, { useEffect, useState } from "react"
import { useHistory } from 'react-router-dom'
import { useDispatch } from 'react-redux'
import { logoutUser } from '../store/action/userAction'

const Navbar = () => {
    const history = useHistory()

    const dispatch = useDispatch()

    const [token, setToken] = useState("")

    useEffect(() => {
        setToken(localStorage.getItem("access_token"))
    })

    const logout = () => {
        dispatch(logoutUser())
        history.push("/login")
    }

    return (
        <header>
            <nav className="navbar navbar-expand-md navbar-light fixed-top bg-light">
                <div className="container-fluid">
                    <a className="navbar-brand" href="#">Password Manager</a>
                    <button className="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#navbarCollapse" aria-controls="navbarCollapse" aria-expanded="false" aria-label="Toggle navigation">
                        <span clasName="navbar-toggler-icon"></span>
                    </button>

                    {!token ? (
                        <div clasName="collapse navbar-collapse" id="navbarCollapse">
                            <ul clasName="navbar-nav me-auto mb-2 mb-md-0">
                                <li clasName="nav-item" onClick={e => history.push("/")}>
                                    <a clasName="nav-link active" aria-current="page">Home</a>
                                </li>
                                <li clasName="nav-item" onClick={e => history.push("/register")}>
                                    <a clasName="nav-link active" aria-current="page">Register</a>
                                </li>
                                <li clasName="nav-item" onClick={e => history.push("/login")}>
                                    <a clasName="nav-link active" aria-current="page">Login</a>
                                </li>

                            </ul>
                        </div>
                    ) : (
                        <div clasName="collapse navbar-collapse" id="navbarCollapse">
                            <ul clasName="navbar-nav me-auto mb-2 mb-md-0">
                                <li clasName="nav-item" onClick={e => history.push("/")}>
                                    <a clasName="nav-link active" aria-current="page">Home</a>
                                </li>
                                <li clasName="nav-item" onClick={logout}>
                                    <a clasName="nav-link active" aria-current="page" href="#">Logout</a>
                                </li>
                            </ul>
                        </div>
                    )}
                </div>
            </nav>
        </header>
    )
}

export default Navbar