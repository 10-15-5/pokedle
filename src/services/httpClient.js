import axios from 'axios'

const apiClient = () => {
    const instance = axios.create({
        baseURL: getBackendApiUrl(),
        withCredentials: true // Send cookies
    })

    return instance
}

const getBackendApiUrl = () => {
    
    // This is coming from VITE as defined replacable variable.
    const url = BUILD_API_URL

    if(!process.env.NODE_ENV === 'development'){
        return `${window.location.origin}${url}`
    }

    return url
}

export {
    apiClient
}