from flask import Flask, jsonify, redirect, url_for

app = Flask(__name__)

@app.route('/favicon.ico')
def favicon():
    return redirect(url_for('static', filename='favicon.ico'))

@app.route("/")
def index():
    return jsonify({
        "success": True
    })

if __name__ == "__main__":
    app.run(host="0.0.0.0", port=8080, debug=True, threaded=True)