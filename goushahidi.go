/*
Package goushahidi provides a client for accessing data from Ushahidi API.

Access different parts of the World Bank Open Data API using the various
services:
         client := goushahidi.NewClient(nil)

         // list all incidences
         incidences, err := client.Countries.GetCountries()


The full Ushahidi API is documented at https://wiki.ushahidi.com/display/WIKI/Ushahidi+3.x+REST+API.
*/
package goushahidi
