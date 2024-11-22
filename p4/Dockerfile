FROM python:3.9-slim
RUN pip install fastapi[standard] uvicorn
COPY work/ /work/
WORKDIR /work
EXPOSE 80
CMD ["uvicorn", "main:app", "--host", "0.0.0.0", "--port", "80", "--proxy-headers", "--log-level", "debug"]