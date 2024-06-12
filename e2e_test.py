import subprocess
import time
import requests
import datetime

def check_login():
    name = datetime.datetime.now().strftime('%Y-%m-%d_%H:%M:%S')
    payload = '{"user_name": "' + name + '", "user_password": "name"}'
    resp = requests.post("http://localhost:8080/user/register", data=payload)
    assert resp.status_code == 200

    resp = requests.post("http://localhost:8080/user/login", data=payload)
    assert resp.status_code == 200

def check_post():
    session = requests.Session()
    login_data='{"user_name": "name", "user_password": "name"}'
    resp = session.post("http://localhost:8080/user/login", data=login_data)
    session.cookies.set(*resp.headers['Set-Cookie'].split('=')[:2])
    assert resp.status_code == 200

    post_data='{"post_title": "title","post_text": "text"}'
    resp = session.post("http://localhost:8080/post/create", data=post_data)
    assert resp.status_code == 200
    id = resp.json()["id"]

    resp = session.get(f"http://localhost:8080/post/get/{id}")
    post = resp.json()
    assert post['post_title'] == 'title'
    assert post['post_text'] == 'text'


subprocess.Popen("docker compose up --build", shell=True, stdout = subprocess.DEVNULL, stderr = subprocess.DEVNULL)
time.sleep(10)
check_login()
check_post()
subprocess.Popen("docker compose down", shell=True, stdout = subprocess.DEVNULL, stderr=subprocess.DEVNULL)

