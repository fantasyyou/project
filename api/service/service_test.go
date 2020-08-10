package service

import (
	"gin/api/dao"
	"gin/api/model"
	"github.com/gin-gonic/gin"
	"reflect"
	"testing"
)

func TestService_ExcelDownLoad(t *testing.T) {
	type fields struct {
		service *Service
		dao     *dao.Dao
	}
	type args struct {
		c *gin.Context
	}
	var tests []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := &Service{
				service: tt.fields.service,
				dao:     tt.fields.dao,
			}
			if err := service.ExcelDownLoad(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("ExcelDownLoad() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestService_Get(t *testing.T) {
	type fields struct {
		service *Service
		dao     *dao.Dao
	}
	type args struct {
		id int
	}
	var tests []struct {
		name      string
		fields    fields
		args      args
		wantOrder model.DemoOrder
		wantErr   bool
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := &Service{
				service: tt.fields.service,
				dao:     tt.fields.dao,
			}
			gotOrder, err := service.Get(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotOrder, tt.wantOrder) {
				t.Errorf("Get() gotOrder = %v, want %v", gotOrder, tt.wantOrder)
			}
		})
	}
}

func TestService_Insert(t *testing.T) {
	type fields struct {
		service *Service
		dao     *dao.Dao
	}
	type args struct {
		order *model.DemoOrder
	}
	var tests []struct {
		name    string
		fields  fields
		args    args
		wantId  int
		wantErr bool
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := &Service{
				service: tt.fields.service,
				dao:     tt.fields.dao,
			}
			gotId, err := service.Insert(tt.args.order)
			if (err != nil) != tt.wantErr {
				t.Errorf("Insert() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotId != tt.wantId {
				t.Errorf("Insert() gotId = %v, want %v", gotId, tt.wantId)
			}
		})
	}
}

func TestService_Query(t *testing.T) {
	type fields struct {
		service *Service
		dao     *dao.Dao
	}
	var tests []struct {
		name       string
		fields     fields
		wantOrders []model.DemoOrder
		wantErr    bool
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := &Service{
				service: tt.fields.service,
				dao:     tt.fields.dao,
			}
			gotOrders, err := service.Query()
			if (err != nil) != tt.wantErr {
				t.Errorf("Query() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotOrders, tt.wantOrders) {
				t.Errorf("Query() gotOrders = %v, want %v", gotOrders, tt.wantOrders)
			}
		})
	}
}

func TestService_SortQuery(t *testing.T) {
	type fields struct {
		service *Service
		dao     *dao.Dao
	}
	type args struct {
		sort string
	}
	var tests []struct {
		name       string
		fields     fields
		args       args
		wantOrders []model.DemoOrder
		wantErr    bool
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := &Service{
				service: tt.fields.service,
				dao:     tt.fields.dao,
			}
			gotOrders, err := service.SortQuery(tt.args.sort)
			if (err != nil) != tt.wantErr {
				t.Errorf("SortQuery() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotOrders, tt.wantOrders) {
				t.Errorf("SortQuery() gotOrders = %v, want %v", gotOrders, tt.wantOrders)
			}
		})
	}
}

func TestService_UpLoad(t *testing.T) {
	type fields struct {
		service *Service
		dao     *dao.Dao
	}
	type args struct {
		c *gin.Context
	}
	var tests []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := &Service{
				service: tt.fields.service,
				dao:     tt.fields.dao,
			}
			if err := service.UpLoad(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("UpLoad() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestService_Update(t *testing.T) {
	type fields struct {
		service *Service
		dao     *dao.Dao
	}
	type args struct {
		order *model.DemoOrder
	}
	var tests []struct {
		name            string
		fields          fields
		args            args
		wantUpdateOrder model.DemoOrder
		wantErr         bool
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := &Service{
				service: tt.fields.service,
				dao:     tt.fields.dao,
			}
			gotUpdateOrder, err := service.Update(tt.args.order)
			if (err != nil) != tt.wantErr {
				t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotUpdateOrder, tt.wantUpdateOrder) {
				t.Errorf("Update() gotUpdateOrder = %v, want %v", gotUpdateOrder, tt.wantUpdateOrder)
			}
		})
	}
}

func TestService_VagueQuery(t *testing.T) {
	type fields struct {
		service *Service
		dao     *dao.Dao
	}
	type args struct {
		username string
	}
	var tests []struct {
		name       string
		fields     fields
		args       args
		wantOrders []model.DemoOrder
		wantErr    bool
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := &Service{
				service: tt.fields.service,
				dao:     tt.fields.dao,
			}
			gotOrders, err := service.VagueQuery(tt.args.username)
			if (err != nil) != tt.wantErr {
				t.Errorf("VagueQuery() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotOrders, tt.wantOrders) {
				t.Errorf("VagueQuery() gotOrders = %v, want %v", gotOrders, tt.wantOrders)
			}
		})
	}
}