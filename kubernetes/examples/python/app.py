from flask import Flask
import requests
import json
app = Flask(__name__)

@app.route("/")
def hello():
  return "Hello"

@app.route("/btcprice")
def btcprice():
  url = 'https://api.coindesk.com/v1/bpi/currentprice.json'
  response = requests.get(url)
  data = response.json()

  return json.dumps(data)
