import uvicorn
from fastapi import FastAPI
import requests

app = FastAPI()

@app.get("/fast/api/http")
async def root():
    print('Hello, World!')
    r = requests.get('http://fastapi-server-service:80')

    return {"message": r.status_code}


@app.get("/hello/{name}")
async def say_hello(name: str):
    print('Hello, World!')
    return {"message": f"Hello {name}"}

if __name__ == "__main__":
    uvicorn.run(app, host="0.0.0.0", port=3003)