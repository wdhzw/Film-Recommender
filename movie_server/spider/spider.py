import json
import random
import time
from datetime import datetime

import boto3
import pandas as pd

import requests

from fake_useragent import UserAgent
from googletrans import Translator
import tmdbsimple as tmdb

tmdb.API_KEY = '4f40401acacceee6b426290d18f6f38a'

ua = UserAgent()
translator = Translator()


def get_html(url):
    while True:
        try:
            time.sleep(random.randint(2, 3))
            res = requests.get(url, headers={'User-Agent': ua.chrome})
            if res.status_code == 200:
                return res.text
        except Exception as e:
            print("request failed, reason: %s" % e)


def crawl_movies(start, end):
    movies_lists = []
    cnt = 0
    for page_num in range(start, end + 1):
        if cnt == 250:
            print("wait for 5 seconds")
            time.sleep(5)
            cnt = 0
        response = tmdb.Movies().popular(page=page_num)
        movies = response['results']
        movies_lists.extend(movies)
        print('page {} done'.format(page_num))
        cnt += 1
    return movies_lists


search_engine = tmdb.Search()


def parse_movies(movies_lists):
    print("before wait for 5 seconds")
    movies = []
    for i in range(len(movies_lists)):
        movie_meta = movies_lists[i]
        if i % 250 == 0:
            print("wait for 5 seconds")
            time.sleep(5)
        movie_id = movie_meta['id']
        movie_details = tmdb.Movies(movie_id).info()

        # title
        title = movie_meta['title']

        # overview
        overview = movie_details['overview']

        average_rate = movie_meta['vote_average']

        popularity = movie_meta['popularity']

        # movie homepage url
        tmdb_movie_url = movie_details['homepage']
        if tmdb_movie_url == '':
            tmdb_movie_url = "https://themoviedb.org/movie/" + str(movie_id)
        # images
        poster_uri = movie_meta['poster_path']
        #  "https://image.tmdb.org/t/p/original" +

        # top casts
        credit_resp = tmdb.Movies(movie_id).credits()
        casts = []
        if len(credit_resp['cast']) > 10:
            casts = credit_resp['cast'][:10]
        else:
            casts = credit_resp['cast']

        actors = []
        for cast in casts:
            actor = {
                'character': cast['character'],
                'name': cast['name'],
                'profile_uri': cast['profile_path']
            }
            actors.append(actor)
        actors_str = json.dumps(actors)
        # director && story
        director = ''
        storiers = []
        for crew in credit_resp['crew']:
            if crew['job'] == 'Director':
                director = crew['name']
            elif crew['job'] == 'Story' or crew['job'] == 'Writer' or crew['job'] == 'Novel' or crew[
                'job'] == 'Screenplay':
                storiers.append(crew['name'])
        writers = ', '.join(storiers)
        # genre
        genres = []
        for genre in movie_details['genres']:
            genres.append(genre['name'])
        genre = ', '.join(genres)
        # production countries
        product_countries = [prd['name'] for prd in movie_details['production_countries']]
        countries = ', '.join(product_countries)
        # language
        languages = [l['english_name'] for l in movie_details['spoken_languages']]
        language = ', '.join(languages)
        # release_date
        release_date = movie_details['release_date']
        # time duration, unit is minutes
        duration = movie_details['runtime']
        # keywords
        keywords = [keyword['name'] for keyword in tmdb.Movies(movie_id).keywords()['keywords']]
        keyword = ', '.join(keywords)

        each_movie = [movie_id, title, overview, average_rate, popularity, tmdb_movie_url, poster_uri, actors_str,
                      director, writers, genre, countries, language, release_date, duration, keyword]
        movies.append(each_movie)
        print("finish 1 movie, num {}".format(i))
    return movies


def save_to_csv(movies):
    item = ['movie_id', 'title', 'overview', 'rate', 'popularity', 'homepage', 'poster_uri', 'actors', 'director',
            'writers', 'genres', 'production_country', 'language', 'release_date', 'duration', 'keyword']
    MOVIES = pd.DataFrame(data=movies, columns=item)  # 转换为DataFrame数据格式
    MOVIES.to_csv('movies.csv', mode='w', encoding='utf-8', index=False)  # 存入csv文件


if __name__ == '__main__':
    s3 = boto3.resource(
        service_name='s3',
        aws_access_key_id="AKIA47CRYK4VYOFU3AWF",
        aws_secret_access_key="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",
        region_name="us-east-1"
    )
    page_num = 600
    movies_lists = crawl_movies(501, page_num)
    movies = parse_movies(movies_lists)
    save_to_csv(movies)
    s3.Bucket('cs5224-movie').upload_file(Filename='movies.csv',
                                          Key="movie_data/movies_{}.csv".format(int(datetime.now().timestamp())))
    print("finish one file")
