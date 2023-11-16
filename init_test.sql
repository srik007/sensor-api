CREATE TABLE IF NOT EXISTS sensors (
    id bigserial PRIMARY KEY,
    name text NOT NULL,
    group_id bigint,
    x numeric,
    y numeric,
    z numeric,
    value bigint,
    format text
);

CREATE INDEX IF NOT EXISTS idx_member ON sensors (name, group_id);

INSERT INTO sensors (name, group_id, x, y, z, value, format)
VALUES
    ('Group1', 1, 10.0, 20.0, 30.0, 100, 'Celsius'),
    ('Group1', 2, 10.0, 20.0, 30.0, 100, 'Celsius'),
    ('Group2', 1, 15.0, 25.0, 35.0, 150, 'Celsius'),
    ('Group2', 2, 10.0, 20.0, 30.0, 100, 'Celsius');

CREATE TABLE IF NOT EXISTS sensor_data (
    id bigserial PRIMARY KEY,
    created_at timestamptz,
    updated_at timestamptz,
    deleted_at timestamptz,
    transparency bigint,
    value numeric,
    scale text,
    species text,
    sensor_id bigint NOT NULL,
    FOREIGN KEY (sensor_id) REFERENCES sensors(id) ON UPDATE CASCADE ON DELETE SET NULL
);

INSERT INTO sensor_data (created_at, updated_at, transparency, value, scale, species, sensor_id)
VALUES
    (CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, 50, 200.0, 'Scale1', '[{"Name":"Fish1","Count":5},{"Name":"Fish2","Count":10}]', 1),
    (CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, 75, 300.0, 'Scale2', '[{"Name":"Fish3","Count":2},{"Name":"Fish4","Count":3}]', 2),
    (CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, 50, 200.0, 'Scale1', '[{"Name":"Fish5","Count":5},{"Name":"Fish6","Count":10}]', 3),
    (CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, 75, 300.0, 'Scale2', '[{"Name":"Fish7","Count":5},{"Name":"Fish8","Count":1}]', 4);

CREATE TABLE IF NOT EXISTS sensor_groups (
    name text PRIMARY KEY,
    sensor_count bigint
);

INSERT INTO sensor_groups (name, sensor_count)
VALUES
    ('Group1', 2),
    ('Group2', 0);
