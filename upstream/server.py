
from flask import Flask, request, abort, jsonify

import uuid

app = Flask(__name__)

uid = uuid.uuid4()

@app.route('/')
def index():
   return jsonify({
      "uuid": str(uid)
   })

@app.route('/healthz')
def health():
   return jsonify({
      "message": "ok"
   })

if __name__ == "__main__":
   app.run(host='0.0.0.0', port=8081, debug=True)