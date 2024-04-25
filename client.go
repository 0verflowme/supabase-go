package supabase

import (
	"errors"

	"github.com/supabase-community/postgrest-go"
	storage "github.com/supabase-community/storage-go"
)

const (
	REST_URL      = "/rest/v1"
	STORAGE_URL   = "/storage/v1"
	DefaultSchema = "public"
)

type Client struct {
	Rest    *postgrest.Client
	Storage *storage.Client
}

type Options struct {
	Headers map[string]string
	Db      *OptionsDb
}

type OptionsDb struct {
	Schema string
}

// NewClient creates a new Supabase client.
func NewClient(url, key string, options *Options) (*Client, error) {
	if url == "" || key == "" {
		return nil, errors.New("url and key are required")
	}

	headers := make(map[string]string)
	headers["Authorization"] = "Bearer " + key
	headers["apikey"] = key

	if options != nil && options.Headers != nil {
		for k, v := range options.Headers {
			headers[k] = v
		}
	}

	schema := DefaultSchema
	if options != nil && options.Db != nil {
		schema = options.Db.Schema
	}

	client := &Client{
		Rest:    postgrest.NewClient(url+REST_URL, schema, headers),
		Storage: storage.NewClient(url+STORAGE_URL, key, headers),
	}

	return client, nil
}

// Wrap postgrest From method
func (c *Client) From(table string) *postgrest.QueryBuilder {
	return c.Rest.From(table)
}

// Wrap postgrest Rpc method
func (c *Client) Rpc(name, count string, rpcBody interface{}) string {
	return c.Rest.Rpc(name, count, rpcBody)
}
