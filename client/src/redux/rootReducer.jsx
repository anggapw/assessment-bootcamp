import userRegisterReducer from './user/register/userRegisterReducer'
import passwordReducer from './password/passwordReducer'

const rootReducer = {
	userRegister: userRegisterReducer,
	webPassword: passwordReducer
}

export default rootReducer
