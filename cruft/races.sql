CREATE TABLE odds (
    id SERIAL,
    value FLOAT NOT NULL,
    recorded_at TIMESTAMP NOT NULL DEFAULT NOW(),
    horse_id INTEGER,
    race_id INTEGER,
    provider_id INTEGER,
    PRIMARY KEY(id),
    FOREIGN KEY (horse_id) REFERENCES horses(id),
    FOREIGN KEY (provider_id) REFERENCES providers(id),
    FOREIGN KEY (race_id) REFERENCES races(id),
    CONSTRAINT only_record_each_bet_once UNIQUE(horse_id, race_id, provider_id, recorded_at)
)

CREATE TABLE bets (
    id SERIAL,
    odd_id INTEGER,
    amount INTEGER NOT NULL DEFAULT 0,
    to_win BOOL NOT NULL DEAFAULT TRUE,
    made_at TIMESTAMP NOT NULL DEFAULT NOW(),
    PRIMARY KEY(id),
    FOREIGN KEY (odd_id) REFERENCES odds(id),
    CONSTRAINT one_bet_per_horse_per_race_per_provider UNIQUE(horse_id, race_id, provider_id)
)

CREATE TABLE horses (
    id SERIAL,
    name VARCHAR(100),
    recorded_at TIMESTAMP NOT NULL DEFAULT NOW(),
    PRIMARY KEY(id),
    CONSTRAINT unique_horse_names UNIQUE(name)
)

CREATE TABLE providers (
    id SERIAl, 
    name VARCHAR(100),
    PRIMARY KEY(id),
    CONSTRAINT unique_provider_names UNIQUE(name)
)

CREATE TABLE venues (
    id SERIAl, 
    name VARCHAR(100),
    PRIMARY KEY(id),
    CONSTRAINT unique_venue_names UNIQUE(name)
)

CREATE TABLE races (
    id SERIAL,
    name VARCHAR(100),
    finished BOOL NOT NULL DEFAULT TRUE,
    venue_id INTEGER,
    occuring_at TIMESTAMP NOT NULL DEFAULT NOW(),
    PRIMARY KEY(id),
    FOREIGN KEY (venue_id) REFERENCES venues(id),
    CONSTRAINT unique_race_time_name UNIQUE(occuring_at, name)
)

CREATE TABLE placements (
    id SERIAL, 
    horse_id INTEGER,
    race_id INTEGER,
    placement INTEGER,
    PRIMARY KEY(id),
    FOREIGN KEY (horse_id) REFERENCES horses(id),
    FOREIGN KEY (race_id) REFERENCES races(id),
    CONSTRAINT one_placement_per_horse_per_race UNIQUE(horse_id, race_id, placement)
)