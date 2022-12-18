create table e_data
(
    id                    serial primary key,
    chat_id               bigint        not null,
    message_id            int           not null,
    created_at             bigint        not null,
    next_interval_time      bigint        not null,
    increasing_coefficient NUMERIC(5, 4) not null,
    active                bool,
    CONSTRAINT chat_mess_un UNIQUE (chat_id,message_id)
);
create table message_to_send
(
    data_id             int,
    sending_num_of_data int,
    time_to_send        bigint,
    CONSTRAINT data_massage_pk
        FOREIGN KEY (data_id) REFERENCES e_data (id),
    CONSTRAINT data_snod_un UNIQUE (data_id, sending_num_of_data)
);
