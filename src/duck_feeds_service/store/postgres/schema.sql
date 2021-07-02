CREATE TABLE IF NOT EXISTS duck_feeds(
    id SERIAL PRIMARY KEY,
    food TEXT NOT NULL,
    food_type TEXT NOT NULL,
    food_quantity INT NOT NULL CHECK(food_quantity > 0),
    park_name TEXT NOT NULL,
    country TEXT NOT NULL,
    ducks_count INT NOT NULL CHECK(ducks_count > 0),
    time TIMESTAMP WITH TIME ZONE NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT now(),
    UNIQUE (id)
);

CREATE INDEX IF NOT EXISTS idx_food_duck_feeds
ON duck_feeds(food);

CREATE INDEX IF NOT EXISTS idx_food_type_duck_feeds
ON duck_feeds(food_type);

CREATE INDEX IF NOT EXISTS idx_park_name_duck_feeds
ON duck_feeds(park_name);

CREATE INDEX IF NOT EXISTS idx_country_duck_feeds
ON duck_feeds(country);

CREATE TABLE IF NOT EXISTS duck_feed_schedules(
    id SERIAL PRIMARY KEY,
    duck_feed_id SERIAL,
    CONSTRAINT fk_duck_feed_id
      FOREIGN KEY(duck_feed_id) 
	    REFERENCES duck_feeds(id)
);
