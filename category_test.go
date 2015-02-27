// Copyright 2014, 2015 Zac Bergquist
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package spotify

import (
	"net/http"
	"testing"
)

func TestGetCategories(t *testing.T) {
	client := testClientString(http.StatusOK, getCategories)
	page, err := client.GetCategories()
	if err != nil {
		t.Fatal(err)
	}
	if l := len(page.Categories); l != 2 {
		t.Fatalf("Expected 2 categories, got %d\n", l)
	}
	if name := page.Categories[1].Name; name != "Mood" {
		t.Errorf("Expected 'Mood', got '%s'", name)
	}
}

func TestGetCategory(t *testing.T) {
	client := testClientString(http.StatusOK, getCategory)
	cat, err := client.GetCategory("dinner")
	if err != nil {
		t.Fatal(err)
	}
	if cat.ID != "dinner" || cat.Name != "Dinner" {
		t.Errorf("Invalid name/id (%s, %s)\n", cat.Name, cat.ID)
	}
}

func TestGetCategoryPlaylists(t *testing.T) {
	client := testClientString(http.StatusOK, getCategoryPlaylists)
	page, err := client.GetCategoryPlaylists("dinner")
	if err != nil {
		t.Fatal(err)
	}
	if l := len(page.Playlists); l != 2 {
		t.Fatalf("Expected 2 playlists, got %d\n", l)
	}
	if name := page.Playlists[0].Name; name != "Dinner with Friends" {
		t.Errorf("Expected 'Dinner with Friends', got '%s'\n", name)
	}
	if tracks := page.Playlists[1].Tracks.Total; tracks != 91 {
		t.Errorf("Expected 'Dinner Music' to have 91 tracks, but got %d\n", tracks)
	}
	if page.Total != 36 {
		t.Errorf("Expected 26 playlists in category 'dinner' - got %d\n", page.Total)
	}
}

var getCategories = `
{
  "categories" : {
    "href" : "https://api.spotify.com/v1/browse/categories?country=CA&offset=0&limit=2",
    "items" : [ {
      "href" : "https://api.spotify.com/v1/browse/categories/toplists",
      "icons" : [ {
        "height" : 275,
        "url" : "https://datsnxq1rwndn.cloudfront.net/media/derived/toplists_11160599e6a04ac5d6f2757f5511778f_0_0_275_275.jpg",
        "width" : 275
      } ],
      "id" : "toplists",
      "name" : "Top Lists"
    }, {
      "href" : "https://api.spotify.com/v1/browse/categories/mood",
      "icons" : [ {
        "height" : 274,
        "url" : "https://datsnxq1rwndn.cloudfront.net/media/original/mood-274x274_976986a31ac8c49794cbdc7246fd5ad7_274x274.jpg",
        "width" : 274
      } ],
      "id" : "mood",
      "name" : "Mood"
    } ],
    "limit" : 2,
    "next" : "https://api.spotify.com/v1/browse/categories?country=CA&offset=2&limit=2",
    "offset" : 0,
    "previous" : null,
    "total" : 31
  }
}`

var getCategory = `
{
  "href" : "https://api.spotify.com/v1/browse/categories/dinner",
  "icons" : [ {
    "height" : 274,
    "url" : "https://datsnxq1rwndn.cloudfront.net/media/original/dinner_1b6506abba0ba52c54e6d695c8571078_274x274.jpg",
    "width" : 274
  } ],
  "id" : "dinner",
  "name" : "Dinner"
}`

var getCategoryPlaylists = `
{
  "playlists" : {
    "href" : "https://api.spotify.com/v1/browse/categories/dinner/playlists?offset=0&limit=2",
    "items" : [ {
      "collaborative" : false,
      "external_urls" : {
        "spotify" : "http://open.spotify.com/user/spotify/playlist/59ZbFPES4DQwEjBpWHzrtC"
      },
      "href" : "https://api.spotify.com/v1/users/spotify/playlists/59ZbFPES4DQwEjBpWHzrtC",
      "id" : "59ZbFPES4DQwEjBpWHzrtC",
      "images" : [ {
        "height" : 300,
        "url" : "https://i.scdn.co/image/68b6a65573a55095e9c0c0c33a274b18e0422736",
        "width" : 300
      } ],
      "name" : "Dinner with Friends",
      "owner" : {
        "external_urls" : {
          "spotify" : "http://open.spotify.com/user/spotify"
        },
        "href" : "https://api.spotify.com/v1/users/spotify",
        "id" : "spotify",
        "type" : "user",
        "uri" : "spotify:user:spotify"
      },
      "public" : null,
      "tracks" : {
        "href" : "https://api.spotify.com/v1/users/spotify/playlists/59ZbFPES4DQwEjBpWHzrtC/tracks",
        "total" : 98
      },
      "type" : "playlist",
      "uri" : "spotify:user:spotify:playlist:59ZbFPES4DQwEjBpWHzrtC"
    }, {
      "collaborative" : false,
      "external_urls" : {
        "spotify" : "http://open.spotify.com/user/spotify/playlist/1WDw5izv4UhpobNdGXQug7"
      },
      "href" : "https://api.spotify.com/v1/users/spotify/playlists/1WDw5izv4UhpobNdGXQug7",
      "id" : "1WDw5izv4UhpobNdGXQug7",
      "images" : [ {
        "height" : 300,
        "url" : "https://i.scdn.co/image/acdcc5e1aa4e9c1db523d684a35f9c0785e50152",
        "width" : 300
      } ],
      "name" : "Dinner Music",
      "owner" : {
        "external_urls" : {
          "spotify" : "http://open.spotify.com/user/spotify"
        },
        "href" : "https://api.spotify.com/v1/users/spotify",
        "id" : "spotify",
        "type" : "user",
        "uri" : "spotify:user:spotify"
      },
      "public" : null,
      "tracks" : {
        "href" : "https://api.spotify.com/v1/users/spotify/playlists/1WDw5izv4UhpobNdGXQug7/tracks",
        "total" : 91
      },
      "type" : "playlist",
      "uri" : "spotify:user:spotify:playlist:1WDw5izv4UhpobNdGXQug7"
    } ],
    "limit" : 2,
    "next" : "https://api.spotify.com/v1/browse/categories/dinner/playlists?offset=2&limit=2",
    "offset" : 0,
    "previous" : null,
    "total" : 36
  }
}`
