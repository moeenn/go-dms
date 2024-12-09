create type
  user_role as enum(
    'admin',
    'subcontractor',
    'site-manager',
    'gatekeeper'
  );

create table
  users (
    user_id uuid default gen_random_uuid (),
    email text not null,
    name text not null,
    password text not null,
    role user_role not null,
    created_at timestamp default now(),
    deleted_at timestamp null,
    primary key (user_id),
    constraint email_unique unique (email)
  );

create table
  sites (
    site_id uuid default gen_random_uuid (),
    name text not null,
    business_id uuid not null,
    location point not null,
    created_at timestamp default now(),
    deleted_at timestamp null,
    primary key (site_id),
    foreign key (business_id) references users (user_id)
  );

create type
  weekday as enum(
    'monday',
    'tuesday',
    'wednesday',
    'thursday',
    'friday',
    'saturday',
    'sunday'
  );

create table
  site_days (
    site_day_id uuid default gen_random_uuid (),
    site_id uuid not null,
    day weekday not null,
    start_time time not null,
    end_time time not null,
    primary key (site_day_id),
    foreign key (site_id) references sites (site_id)
  );

create type
  site_resource_type as enum('forklift', 'crane', 'offloading-bay', 'trolley');

create table
  site_resources (
    site_resource_id uuid default gen_random_uuid (),
    site_id uuid not null,
    name text not null,
    resource_type site_resource_type not null,
    created_at timestamp default now(),
    deleted_at timestamp null,
    primary key (site_resource_id),
    foreign key (site_id) references sites (site_id)
  );

create table
  site_subcontractors (
    site_id uuid not null,
    subcontractor_id uuid not null,
    primary key (site_id, subcontractor_id),
    foreign key (site_id) references sites (site_id),
    foreign key (subcontractor_id) references users (user_id)
  );

create table
  vendors (
    vendor_id uuid default gen_random_uuid (),
    name text not null,
    description text not null,
    phone text not null,
    email text not null,
    primary key (vendor_id)
  );

create table
  site_vendors (
    site_id uuid not null,
    vendor_id uuid not null,
    primary key (site_id, vendor_id),
    foreign key (site_id) references sites (site_id),
    foreign key (vendor_id) references vendors (vendor_id)
  );

create table
  bookings (
    booking_id uuid default gen_random_uuid (),
    site_id uuid not null,
    start_timestamp timestamp not null,
    end_timestamp timestamp not null,
    duration text not null,
    created_at timestamp default now(),
    deleted_at timestamp null,
    primary key (booking_id),
    foreign key (site_id) references sites (site_id)
  );

create table
  booking_resources (
    booking_id uuid,
    site_resource_id uuid,
    primary key (booking_id, site_resource_id),
    foreign key (booking_id) references bookings (booking_id),
    foreign key (site_resource_id) references site_resources (site_resource_id)
  );

create type
  delivery_status as enum(
    'pending-approval',
    'approved',
    'arrived',
    'completed',
    'cancelled'
  );

create table
  delivery (
    delivery_id uuid default gen_random_uuid (),
    booking_id uuid not null,
    vendor_id uuid not null,
    status delivery_status not null,
    notes text null,
    primary key (delivery_id),
    foreign key (booking_id) references bookings (booking_id),
    foreign key (vendor_id) references vendors (vendor_id)
  );