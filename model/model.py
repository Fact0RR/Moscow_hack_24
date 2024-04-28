import re
import psycopg2
from thefuzz import fuzz
import json

class Course:
        def __init__(self,id,name,link,type,duration,score):
            self.id = id
            self.name = name
            self.link = link
            self.type = type
            self.duration = duration
            self.score = score

def getRate(text):
    words = handlerForText(text)
    conn = psycopg2.connect(dbname="tilt_db", user="tilt", password="tilt_pass", host="db")
    cursor = conn.cursor()

    cursor.execute("SELECT name FROM tags")

    rateTags = dict()

    for course in cursor.fetchall():
        fuzzsum = 0
        for word in words:
            flf = float(fuzz.ratio(word.lower(),course[0].lower()))/100
            flf2 = flf*flf*flf
            fuzzsum +=flf2
        #print(tag[0], fuzzsum)
        rateTags.setdefault(course[0],fuzzsum)

    cursor.execute("SELECT name,id,link,type,duration FROM courses")
        
    сoursesList = []

    for course in cursor.fetchall():
        fuzzsum = 0
        for word in words:
            flf = float(fuzz.ratio(word.lower(),course[0].lower()))/100
            flf2 = flf*flf
            fuzzsum +=flf2
        сoursesList.append(Course(course[1],course[0],course[2],course[3],course[4],fuzzsum))
    сoursesList.sort(key=lambda сourse: сourse.score, reverse=True)
        
    cursor.close()
    conn.close()

    countsum = 0.0

    for course in сoursesList:
        countsum += course.score

    countsum = countsum / len(сoursesList)

    finalCoursesList = []

    for course in сoursesList:
        if course.score > countsum:
            finalCoursesList.append(course)

    return objToJSON(finalCoursesList)

def handlerForText(text):
    text = re.sub(r'[^\w\s]', '', text)
    words = text.split()
    return words

def objToJSON(list):
    
    final_string = json.dumps([ob.__dict__ for ob in list])
    return final_string
