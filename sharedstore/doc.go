// Package sharedstore provides shared storage where only one thread (local or on another instance) works on preparing the data for a key.
//
// While another Store instance has the lock, the store will poll periodically until it’s unlocked.
// To reduce the stress on the Client, other threads of the local store will wait for the polling thread to complete.
//
// Depends on memcached, but in a minimal fashion and could be decoupled easily.
package sharedstore
