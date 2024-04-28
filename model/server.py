
from flask import Flask, app, request
from model import getRate
import json

app = Flask(__name__)

def getTagsRate(jsonStr):
    
    return getRate(jsonStr["text"])


@app.route('/send', methods=['POST'])
def treatmentVideo():
    return getTagsRate(request.json)


app.run(host='0.0.0.0', port=5000)