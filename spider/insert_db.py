import csv

import boto3
import pandas as pd
import pymysql

sql_create = """
CREATE TABLE IF NOT EXISTS movies (
    id INT AUTO_INCREMENT PRIMARY KEY,
    movie_id BIGINT UNIQUE,
    title VARCHAR(255),
    overview TEXT,
    rate FLOAT,
    popularity FLOAT,
    homepage VARCHAR(255),
    poster_uri VARCHAR(255),
    actors JSON,
    director VARCHAR(255),
    writers VARCHAR(1000),
    genres VARCHAR(1000),
    production_country VARCHAR(255),
    language VARCHAR(255),
    release_date DATE,
    duration INT,
    keyword VARCHAR(255),
    create_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    update_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);"""

insert_sql = """
INSERT IGNORE INTO movies (
    movie_id, title, overview, rate, popularity, homepage, poster_uri,
    actors, director, writers, genres, production_country, language,
    release_date, duration, keyword
) VALUES (
    %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s
)
"""

if __name__ == '__main__':
    s3 = boto3.client(
        service_name='s3',
        aws_access_key_id="AKIA47CRYK4VYOFU3AWF",
        aws_secret_access_key="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",
        region_name="us-east-1"
    )

    # rds_client = boto3.client(
    #     service_name='rds',
    #     aws_access_key_id="AKIA47CRYK4VYOFU3AWF",
    #     aws_secret_access_key="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",
    #     region_name='us-east-1')
    #
    # resp = rds_client.describe_db_instances(DBInstanceIdentifier='cs5224-movie')

    db_conn = pymysql.connect(
        host='cs5224-movie-meta.c3o8o8eyqv2b.us-east-1.rds.amazonaws.com',
        port=3306,
        user='guohaonan',
        password='ghn980421',
        database='movie',
        charset='utf8mb4'
    )

    cur = db_conn.cursor()
    # cur.execute("create database movie")  # create database
    # cur.execute(sql_create)

    resp = s3.list_objects_v2(Bucket='cs5224-movie', Prefix='movie_data/')

    for obj in resp.get('Contents'):
        if obj['Key'] == 'movie_data/':
            continue

        object_key = obj['Key'].split('/')[1]
        local_file_path = f'./files/{object_key}'
        s3.download_file('cs5224-movie', obj['Key'], local_file_path)

        movies = pd.read_csv(local_file_path)
        movies.fillna('', inplace=True)
        movies.drop_duplicates(subset=['movie_id'], inplace=True)

        cnt = 1
        for index, row in movies.iterrows():
            data_to_insert = (
                row['movie_id'], row['title'], row['overview'], row['rate'], row['popularity'],
                row['homepage'], row['poster_uri'], row['actors'], row['director'], row['writers'],
                row['genres'], row['production_country'], row['language'], row['release_date'],
                row['duration'], row['keyword']
            )

            cur.execute(insert_sql, data_to_insert)
            print("The {} row inserted".format(cnt))
            if cnt % 200 == 0:
                print("every 200 rows submit the transactions")
                db_conn.commit()
            cnt += 1
        db_conn.commit()
        print("file {} handle done~".format(object_key))

    db_conn.close()
