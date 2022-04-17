import requests
import time

def create(n):
    start_time = time.time()
    for i in range(n):
        res = requests.post("http://localhost:8081/db/create", json={
            "name": "testUser1",
            "password": "123",
            "email": "testuser1@gmail.com"
        })
        print(i, res.content.decode(), end='\r')
    print(f'\nTotal time: {time.time() - start_time}')

def update(n):
    start_time = time.time()
    for i in range(n):
        res = requests.get(f"http://localhost:8081/db/update?email=testuser1@gmail.com&field=name&value=Rachel{i}")
        print(i, res.content.decode(), end='\r')
    print(f'\nTotal time: {time.time() - start_time}')


def delete(n):
    start_time = time.time()
    for i in range(n):
        res = requests.get(f"http://localhost:8081/db/delete?email=testuser1@gmail.com")
        print(i, res.content.decode(), end='\r')
    print(f'\nTotal time: {time.time() - start_time}')


if __name__ == "__main__":
    create(200)
    update(200)
    delete(200)
