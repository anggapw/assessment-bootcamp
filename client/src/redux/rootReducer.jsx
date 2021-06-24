import userRegisterReducer from './user/register/userRegisterReducer'
import passwordReducer from './password/passwordReducer'
import userLoginReducer from './user/login/userLoginReducer'

const rootReducer = {
	userRegister: userRegisterReducer,
	webPassword: passwordReducer,
	userLogin: userLoginReducer
}

export default rootReducer
