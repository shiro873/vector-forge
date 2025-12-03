from fastapi import FastAPI
app = FastAPI(title="AI Embedding Service")

@app.get("/health")
def health():
    return {"status":"ok"}

@app.post("/embed")
def embed(doc: dict):
    # stub: return a fixed vector for now
    return {"embedding":[0.0]*768}
