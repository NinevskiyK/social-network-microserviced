FROM python:latest

WORKDIR /app

COPY requirements.txt .
RUN pip install --no-cache-dir -r requirements.txt
COPY src .
COPY .env .

ENTRYPOINT [ "python3", "main.py" ]