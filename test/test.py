import requests


numbers = [1, 2, 4, 9]

for num in numbers:
    resp = requests.post("http://127.0.0.1:8080", json={
        "value": num,
    })

    if resp.status_code != 200:
        print(f"error adding number: {num}\n error: {resp.content}")
        exit(1)

    print(f"added number: {num}\napi returned: {resp.json()}")
