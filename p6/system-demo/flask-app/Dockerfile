FROM python:3.9-slim
WORKDIR /app
RUN pip install flask psycopg2-binary python-dotenv
COPY . .
CMD [ "python", "app.py"]
