create table payment
(
    id             int auto_increment,
    user_id        int            not null,
    order_id       varchar(100)   not null,
    transaction_id varchar(100)   not null,
    amount         decimal(10, 2) not null,
    pay_at         datetime       not null,
    created_at     datetime       not null default current_timestamp,
    updated_at     datetime       not null default current_timestamp on update current_timestamp,
    constraint payment_pk primary key (id)
);