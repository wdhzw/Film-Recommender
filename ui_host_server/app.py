import os
from flask import Flask, send_from_directory
from flask_cors import CORS

flask_app = Flask(
    __name__,
    static_folder='./static/build/static'
)
flask_app.secret_key = os.urandom(24)  # Use a secure random key in production
CORS(flask_app)

@flask_app.route("/")
def home_page():
    return send_from_directory('static/build', 'index.html')


if __name__ == '__main__':
    flask_app.run("0.0.0.0", debug=True, port=8080)