from pydantic import BaseModel

class EBLInRequest(BaseModel):
    query: str

