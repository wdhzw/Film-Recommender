import numpy as np
import pandas as pd

if __name__ == '__main__':
    movie_list = [
        pd.read_csv("./files/movies_1712071845.csv"),
        pd.read_csv("./files/movies_1712072893.csv"),
        pd.read_csv("./files/movies_1712074316.csv"),
        pd.read_csv("./files/movies_1712076761.csv"),
        pd.read_csv("./files/movies_1712078514.csv"),
    ]
    movie = pd.concat(movie_list, ignore_index=True)
    res = set()
    for index, row in movie.iterrows():
        genres = row['genres']
        if isinstance(genres, str):
            genres = genres.split(',')
            for genre in genres:
                if genre not in res:
                    res.add(genre)

    # 输出所有类型
    print(res)
