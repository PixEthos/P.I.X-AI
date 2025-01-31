// I use GPL2

/* Copyright (C) 2024, 2025 PixEthos */

/* This file is part of PixAI.

PixAI is free software: you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation, either version 2 of the License, or (at your option) any later version.

PixAI is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for more details.

You should have received a copy of the GNU General Public License along with PixAI. If not, see <https://www.gnu.org/licenses/>. */

// cache.go
package cache

import (
	"errors"
	"sync"
)

/* Found an article that gave me an idea, though this time I am splitting values */

type Cache[K comparable, V any] struct {

	// usable eventually
	mutex sync.Mutex
	mapping sync.Map

	// regular cache
	regcache map[K]V

	// 64/32bit values witin this area
	f64bit map[K]float64
	i64bit map[K]int64
	f32bit map[K]float32
	i32bit map[K]int32
}

// error handling, checking whether or not the cache are nil
func (c *Cache[K, V]) ErrorChecks(key K, value V) error {
	if c.f64bit == nil || c.i64bit == nil {
		return errors.New("nil values within the 64bit cache")
	}

	if c.f32bit == nil || c.i32bit == nil {
		return errors.New("nil values within the 32bit cache")
	}

	if c.regcache == nil {
		return errors.New("nil values in non-type specific cache")
	}

	return nil
}

func nCache[K comparable, V any]() *Cache[K, V] {
	c := Cache[K, V]{}
	return &c
}

// regular caching, not type specific
func RegCache[K comparable, V any]() *Cache[K, V] {
	return &Cache[K, V]{
		regcache: make(map[K]V),
	}
}

func (c* Cache[K, V]) SetReg(key K, value V) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.regcache[key] = value
}

func (c *Cache[K, V]) GetReg(key K) (V, bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	value, found := c.regcache[key]
	return value, found
}

/* Type specific caching, the reason I did this is for type safety. */
// creating float64 and int64 caches
func Newf64[K comparable, V any]() *Cache[K, V] {
	return &Cache[K, V]{
		f64bit: make(map[K]float64),
	}
}

func (c* Cache[K, V]) Setf64bit(key K, value float64) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.f64bit[key] = value
}

func (c *Cache[K, V]) Getf64bit(key K) (float64, bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	value, found := c.f64bit[key]
	return value, found
}


// int64
func Newi64[K comparable, V any]() *Cache[K, V] {
	return &Cache[K, V]{
		i64bit: make(map[K]int64),
	}
}

func (c* Cache[K, V]) Seti64bit(key K, value int64) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.i64bit[key] = value
}

func (c *Cache[K, V]) Geti64bit(key K) (int64, bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	value, found := c.i64bit[key]
	return value, found
}


// creating float32 and int32 caches
func Newf32[K comparable, V any]() *Cache[K, V] {
	return &Cache[K, V]{
		f32bit: make(map[K]float32),
	}
}

// setting 32bit caches
func (c* Cache[K, V]) Setf32bit(key K, value float32) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.f32bit[key] = value
}

// getting 32bit caches
func (c *Cache[K, V]) Getf32bit(key K) (float32, bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	value, found := c.f32bit[key]
	return value, found
}

// int32
func Newi32[K comparable, V any]() *Cache[K, V] {
	return &Cache[K, V]{
		i32bit: make(map[K]int32),
	}
}

func (c* Cache[K, V]) Seti32bit(key K, value int32) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.i32bit[key] = value
}

func (c *Cache[K, V]) Geti32bit(key K) (int32, bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	value, found := c.i32bit[key]
	return value, found
}

/* So, I found the code for this basically here: https://www.alexedwards.net/blog/implementing-an-in-memory-cache-in-go
 credit where credit is due, and it's ideal for me to be honest where I find the code I use

 That being said, it's just simple caching, and this can work for anything, I just split
 the values between 32bit and 64bit for the fact of usability and better safety.

 Even if the solution provided is simple; I prefer safty first and foremost.

 Because even if you set an array/value to equal or adjust to an array - you'd need to
 deallocate for each array. Each array is separate; so why only use a singlular value? */

// removing caches
func (c *Cache[K, V]) Remove(key K) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	delete(c.regcache, key)
	delete(c.f64bit, key)
	delete(c.i64bit, key)
	delete(c.f32bit, key)
	delete(c.i32bit, key)
}

func (c *Cache[K, V]) PopReg(key K) (V, bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	value, found := c.regcache[key]

	if found {
		delete(c.regcache, key)
	}

	return value, found
}

// deleting 64bit caches
func (c *Cache[K, V]) Popf64bit(key K) (float64, bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	value, found := c.f64bit[key]

	if found {
		delete(c.f64bit, key)
	}

	return value, found
}

func (c *Cache[K, V]) Popi64bit(key K) (int64, bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	value, found := c.i64bit[key]

	if found {
		delete(c.i64bit, key)
	}

	return value, found
}

// deleting 32bit caches
func (c *Cache[K, V]) Popf32bit(key K) (float32, bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	value, found := c.f32bit[key]

	if found {
		delete(c.f32bit, key)
	}

	return value, found
}

func (c *Cache[K, V]) Popi32bit(key K) (int32, bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	value, found := c.i32bit[key]

	if found {
		delete(c.i32bit, key)
	}

	return value, found
}
