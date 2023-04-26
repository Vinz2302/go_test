create table driver_incentives (
	id serial primary key,
	booking_id int,
	incentive bigint,
	
	constraint booking_id foreign key (booking_id)
		references bookings (id)
);
