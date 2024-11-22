from flask import Flask, request, jsonify
import os
import psycopg2

app = Flask(__name__)

@app.route('/')
def hello():
    return "Hello from Flask!"

@app.route('/db-test')
def db_test():
    conn = psycopg2.connect(os.environ['DATABASE_URL'])
    cur = conn.cursor()
    cur.execute('SELECT 1')
    result = cur.fetchone()
    cur.close()
    conn.close()
    return f"Database connection successful: {result}"

@app.route('/items', methods = ['POST'])
def add_items():
    data = request.get_json()
    if not data or 'name' not in data:
        return jsonify({"error": "name is required"}), 400
    conn = psycopg2.connect(os.environ['DATABASE_URL'])
    cur = conn.cursor()

    try:
        cur.execute(
            "INSERT INTO test (name) VALUES (%s) RETURNING id, name",
            (data['name'],)
        )
        new_item = cur.fetchone()
        conn.commit()

        return jsonify({
            'id':new_item[0],
            "name": new_item[1]
        })
    except Exception as e:
        conn.rollback()
        return jsonify({"error": str(e)}), 500
    
    finally:
        cur.close()
        conn.close()

@app.route('/items', methods=['GET'])
def get_items():
    conn = psycopg2.connect(os.environ['DATABASE_URL'])
    cur = conn.cursor()
    
    try:
        cur.execute("SELECT id, name FROM test")
        items = cur.fetchall()
        
        return jsonify([
            {"id": item[0], "name": item[1]}
            for item in items
        ])
        
    finally:
        cur.close()
        conn.close()


if __name__ == '__main__':
    app.run(host='0.0.0.0', port=8080, debug=True)
