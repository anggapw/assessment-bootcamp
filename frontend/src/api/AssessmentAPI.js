import Axios from 'axios'

const AssessmentAPI = Axios.create({
    baseURL : "http://localhost:8080"
})

export default AssessmentAPI