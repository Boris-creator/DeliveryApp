package models

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type Workshop struct {
	Id        int     `db:"id"         sql:"omit_on_insert"`
	Name      *string `db:"name"`
	AddressId int     `db:"address_id"`
	CreatedAt string  `db:"created_at"`
}

func FindNearestWorkshop(db *sqlx.DB, lat, lon float32) (Workshop, error) {
	// using https://github.com/zachasme/h3-pg
	resolution := 5
	q := `WITH target as (
		select 
		  h3_lat_lng_to_cell(
			POINT($1,$2), 
			$3
		  )
	  ) 
	  SELECT 
		workshops.*, 
		h3_grid_distance(
		  h3_lat_lng_to_cell(
			POINT(geo_lat, geo_lon), 
			$3
		  ), 
		  (
			SELECT 
			  * 
			from 
			  target
		  )
		) as distance 
	  FROM 
		workshops 
		JOIN address ON workshops.address_id = address.id 
	  WHERE 
		h3_lat_lng_to_cell(
		  POINT(geo_lat, geo_lon), 
		  $3
		) in (
		  SELECT 
			h3_grid_disk(
			  (
				SELECT 
				  * 
				from 
				  target
			  ), 
			  3
			)
		) 
	  order by 
		distance 
	  limit 
		1;
	  `

	var w Workshop
	err := db.Unsafe().Get(&w, q, fmt.Sprintf("%.6f", lat), fmt.Sprintf("%.6f", lon), resolution)

	return w, err
}
