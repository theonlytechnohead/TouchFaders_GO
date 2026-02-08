package main

type parameter struct {
	system_exclusive byte
	manufacturer_id  byte
	device_id        byte
	group_id         byte
	model_id         byte
	data_category    byte
	element          [2]byte
	index            [2]byte
	channel          [2]byte
	data             []byte
	end_exclusive    byte
}
