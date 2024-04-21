const userURL = `http://ec2-44-217-97-83.compute-1.amazonaws.com:8080/api/`
const movieURL = `http://cs5224-movie-service.us-east-1.elasticbeanstalk.com/movie_server/`
const searchURL = `http://cs5224-movie-service.us-east-1.elasticbeanstalk.com/movie_server/search`
const recommendURL = `http://recommendation-service-env.eba-p8im7vcj.us-east-1.elasticbeanstalk.com/recommendations`


const login = async (username, email, pwd) => {
    const headers = new Headers();
    headers.append("Content-Type", "application/json");
    headers.append("Accept", "*/*")
    headers.append("Accept-Encoding", "gzip, deflate, br")

    const raw = JSON.stringify({
      "user_name": username,
      "email": email,
      "password": pwd
    });

    try {
        const response = await fetch(userURL + 'user_login', {
            method: 'POST',
            headers: headers,
            body: raw,
        })
        if (response.ok) {
            return await response.json()
        } else {
            console.error('Login failed', await response.text())
            return null
        }
    } catch (error) {
        console.error('Error:', error)
        return null
    }
}

const getUserInfo= async (email) => {
    const headers = new Headers();
    headers.append("Content-Type", "application/json");
    headers.append("Accept", "*/*")
    headers.append("Accept-Encoding", "gzip, deflate, br")

    const raw = JSON.stringify({
      "email": email,
    });

    try {
        const response = await fetch(userURL + 'get_user_by_email', {
            method: 'POST',
            headers: headers,
            body: raw,
        })
        if (response.ok) {
            return await response.json()
        } else {
            console.error('Fetch user info failed', await response.text())
            return null
        }
    } catch (error) {
        console.error('Error:', error)
        return null
    }
}

const addGenre = async(username, email, newGenre) => {
    const headers = new Headers();
    headers.append("Content-Type", "application/json");
    headers.append("Accept", "*/*")
    headers.append("Accept-Encoding", "gzip, deflate, br")

    const raw = JSON.stringify({
      "user_name": username,
      "email": email,
      "genre": newGenre,
    });


    try {
        const response = await fetch(userURL + 'update_user_genre', {
            method: 'POST',
            headers: headers,
            body: raw,
        })
        if (response.ok) {
            return await response.json()
        } else {
            console.error('Add user genre failed', await response.text())
            return null
        }
    } catch (error) {
        console.error('Error:', error)
        return null
    }
}

const signUp = async (username, email, pwd) => {
    const headers = new Headers();
    headers.append("Content-Type", "application/json");
    headers.append("Accept", "*/*")
    headers.append("Accept-Encoding", "gzip, deflate, br")

    const raw = JSON.stringify({
      "user_name": username,
      "email": email,
      "password": pwd,
    });

    try {
        const response = await fetch(userURL + 'user_sign_up', {
            method: 'POST',
            headers: headers,
            body: raw,
        })
        if (response.ok) {
            return await response.json()
        } else {
            console.error('Sign up failed', await response.text())
            return null
        }
    } catch (error) {
        console.error('Error:', error)
        return null
    }
}

const confirmSignUp = async (username, email, code) => {
    const headers = new Headers();
    headers.append("Content-Type", "application/json");
    headers.append("Accept", "*/*")
    headers.append("Accept-Encoding", "gzip, deflate, br")

    const raw = JSON.stringify({
      "user_name": username,
      "email": email,
      "code": code,
    });

    try {
        const response = await fetch(userURL + 'confirm_user_sign_up', {
            method: 'POST',
            headers: headers,
            body: raw,
        })
        if (response.ok) {
            return await response.json()
        } else {
            console.error('Sign up failed', await response.text())
            return null
        }
    } catch (error) {
        console.error('Error:', error)
        return null
    }
}

const getMovieDetail = async(movieID) => {
    const headers = new Headers();
    headers.append("Content-Type", "application/json");
    headers.append("Accept", "*/*")
    headers.append("Accept-Encoding", "gzip, deflate, br")

    try {
        const response = await fetch(movieURL + movieID, {
            method: 'GET',
            headers: headers,
        })
        if (response.ok) {
            return await response.json()
        } else {
            console.error('Get movie detail failed', await response.text())
            return null
        }
    } catch (error) {
        console.error('Error:', error)
        return null
    }
}

const searchMovie = async(keywords) => {
    const headers = new Headers();
    headers.append("Content-Type", "application/json");
    headers.append("Accept", "*/*")
    headers.append("Accept-Encoding", "gzip, deflate, br")

    try {
        const response = await fetch(searchURL + "?search_word=" + keywords, {
            method: 'GET',
            headers: headers,
        })
        if (response.ok) {
            return await response.json()
        } else {
            console.error('Get search list failed', await response.text())
            return null
        }
    } catch (error) {
        console.error('Error:', error)
        return null
    }
}

const getRecommendations = async (email, pageID) => {
    const headers = new Headers();
    headers.append("Content-Type", "application/json");
    headers.append("Accept", "*/*")
    headers.append("Accept-Encoding", "gzip, deflate, br")

    try {
        const response = await fetch(recommendURL + "?email=" + email + "&page=" + pageID, {
            method: 'GET',
            headers: headers,
        })
        if (response.ok) {
            return await response.json()
        } else {
            console.error('Get recommendation list failed', await response.text())
            return null
        }
    } catch (error) {
        console.error('Error:', error)
        return null
    }
}


export {
    login,
    getUserInfo,
    addGenre,
    signUp,
    confirmSignUp,
    getMovieDetail,
    searchMovie,
    getRecommendations,
}