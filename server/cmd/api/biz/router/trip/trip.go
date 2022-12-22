// Code generated by hertz generator. DO NOT EDIT.

package Trip

import (
	"github.com/cloudwego/hertz/pkg/app/server"

	trip "github.com/CyanAsterisk/FreeCar/server/cmd/api/biz/handler/trip"
)

/*
 This file will register all the routes of the services in the master idl.
 And it will update automatically when you use the "update" command for the idl.
 So don't modify the contents of the file, or your code will be deleted when it is updated.
*/

// Register register routes based on the IDL 'api.${HTTP Method}' annotation.
func Register(r *server.Hertz) {

	root := r.Group("/", rootMw()...)
	{
		_v1 := root.Group("/v1", _v1Mw()...)
		_v1.GET("/trips", append(_gettripsMw(), trip.GetTrips)...)
		_v1.POST("/trip", append(_tripMw(), trip.CreateTrip)...)
		_trip := _v1.Group("/trip", _tripMw()...)
		_trip.GET("/:id", append(_gettripMw(), trip.GetTrip)...)
		_trip.PUT("/:id", append(_updatetripMw(), trip.UpdateTrip)...)
	}
}