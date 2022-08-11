FROM python:3.7-slim

WORKDIR /app

RUN apt-get update \
    && apt-get install gcc dnsutils -y \
    && apt-get clean

RUN python3 -m pip install --upgrade pip
COPY ./requirements.txt .
RUN python3 -m pip install -r requirements.txt

COPY /cms-api /app/cms-api

CMD ["uvicorn", "cms-api.main:app", "--host", "0.0.0.0", "--port", "8000"]
