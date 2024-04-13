const baseURL = `http://ec2-34-206-66-131.compute-1.amazonaws.com:8080/api/`

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
        const response = await fetch(baseURL + 'user_login', {
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
        const response = await fetch(baseURL + 'get_user_by_email', {
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
        const response = await fetch(baseURL + 'update_user_genre', {
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
        const response = await fetch(baseURL + 'user_sign_up', {
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
        const response = await fetch(baseURL + 'confirm_user_sign_up', {
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

export {
    login,
    getUserInfo,
    addGenre,
    signUp,
    confirmSignUp
}