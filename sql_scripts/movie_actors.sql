CREATE TABLE movie_actors (
    movie_id INT NOT NULL,
    actor_id INT NOT NULL,
    FOREIGN KEY (movie_id) REFERENCES movies(movie_id) ON DELETE CASCADE,
    FOREIGN KEY (actor_id) REFERENCES actors(actor_id) ON DELETE CASCADE,
    PRIMARY KEY (movie_id, actor_id)
);