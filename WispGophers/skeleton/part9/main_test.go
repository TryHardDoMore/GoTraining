package main

import (
	"net"
	"reflect"
	"sync"
	"testing"
)

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		{name:"xxx"},
		{name:"yyy"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}

func TestPeers_Add(t *testing.T) {
	type fields struct {
		m  map[string]chan<- Message
		mu sync.RWMutex
	}
	type args struct {
		addr string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   <-chan Message
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Peers{
				m:  tt.fields.m,
				mu: tt.fields.mu,
			}
			if got := p.Add(tt.args.addr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Peers.Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPeers_Remove(t *testing.T) {
	type fields struct {
		m  map[string]chan<- Message
		mu sync.RWMutex
	}
	type args struct {
		addr string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Peers{
				m:  tt.fields.m,
				mu: tt.fields.mu,
			}
			p.Remove(tt.args.addr)
		})
	}
}

func TestPeers_List(t *testing.T) {
	type fields struct {
		m  map[string]chan<- Message
		mu sync.RWMutex
	}
	tests := []struct {
		name   string
		fields fields
		want   []chan<- Message
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Peers{
				m:  tt.fields.m,
				mu: tt.fields.mu,
			}
			if got := p.List(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Peers.List() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_broadcast(t *testing.T) {
	type args struct {
		m Message
	}
	tests := []struct {
		name string
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			broadcast(tt.args.m)
		})
	}
}

func Test_serve(t *testing.T) {
	type args struct {
		c net.Conn
	}
	tests := []struct {
		name string
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			serve(tt.args.c)
		})
	}
}

func Test_readInput(t *testing.T) {
	tests := []struct {
		name string
	}{
	// TODO: Add test cases.
	}
	for range tests {
		t.Run(tt.name, func(t *testing.T) {
			readInput()
		})
	}
}

func Test_dial(t *testing.T) {
	type args struct {
		addr string
	}
	tests := []struct {
		name string
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dial(tt.args.addr)
		})
	}
}

func TestSeen(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Seen(tt.args.id); got != tt.want {
				t.Errorf("Seen() = %v, want %v", got, tt.want)
			}
		})
	}
}
