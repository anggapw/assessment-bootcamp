import AssessmentAPI from '../../api/AssessmentAPI'

export const registerUser = (payload) => {
    return async (dispatch) => {
        try {
            dispatch({ type : "USER_LOADING"})

            const { data } = await AssessmentAPI({
                method : "POST",
                url : "/users/register",
                data : payload,
            })

            console.log(data)

            return dispatch({ type : "USER_REGISTER", payload : data})

        } catch(err) {
            console.log(err.respose)
        }
    }
}

export const loginUser = (payload) => {
    return async (dispatch) => {
        try {
            dispatch({ type : "USER_LOADING"})

            const { data } = await AssessmentAPI({
                method : "POST",
                url : "/users/login",
                data : payload,
            })

            localStorage.setItem("access_token", data.authorization)

            console.log(data)

            return dispatch({ type : "USER_LOGIN", payload : data})

        } catch(err) {
            console.log(err.respose)
        }
    }
}

export const logoutUser = () => ({
    type : "LOGOUT_USER"
})