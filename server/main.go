package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"strings"
	
	pb "app/currency"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedCurrencyServiceServer
}


func (s *server) GetCurrency(ctc context.Context, req *pb.Currency) (*pb.GetAllCurrencyRespone, error) {

	r, err := http.Get("https://cbu.uz/oz/arkhiv-kursov-valyut/json/")
	if err != nil{
		return nil, err
	}

	data, err := ioutil.ReadAll(r.Body)
	if err != nil{
		return nil, err
	}

	var allcurrency []*pb.CurrencyResponse

	err = json.Unmarshal(data, &allcurrency)
	if err != nil{
		return nil, err
	}

	resp := &pb.GetAllCurrencyRespone{}

	for _, v := range allcurrency{
		if strings.Contains(v.Ccy, req.Currency) {
			resp.Currencies = append(resp.Currencies, &pb.CurrencyResponse{
				Id: v.Id,
				Code: v.Code,
				Ccy: v.Ccy,
				CcyNm_RU: v.CcyNm_RU,
				CcyNm_UZ: v.CcyNm_UZ,
				CcyNm_UZC: v.CcyNm_UZC,
				CcyNm_EN: v.CcyNm_EN,
				Nominal: v.Nominal,
				Rate: v.Rate,
				Diff: v.Diff,
				Date: v.Date,
			})
		}
	}

	return resp, nil
}

func main() {

	lis, err := net.Listen("tcp", "localhost:9001")
	if err != nil{
		log.Fatalf("failed to listen: %+v", err)
	}

	s := grpc.NewServer()
	pb.RegisterCurrencyServiceServer(s, &server{})


	fmt.Println("Liste GRPC server ...")

	err = s.Serve(lis)
	if err != nil{
		log.Fatalf("failed to serve: %+v", err)
	}

}