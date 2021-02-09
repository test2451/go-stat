package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"github.com/pancakeswap/pancake-statas/statas"
	"github.com/pancakeswap/pancake-statas/util"
)

const (
	DefaultListenAddr = "0.0.0.0:8080"
)

type Server struct {
	config *util.Config

	statSvc *statas.StatasSvc
}

func NewServer(config *util.Config, statSvc *statas.StatasSvc) *Server {
	return &Server{
		config:  config,
		statSvc: statSvc,
	}
}

func (s *Server) Stat(w http.ResponseWriter, r *http.Request) {
		swapPiars, totalVolume, lockVolume, updateAt := s.statSvc.GetSwapPairInfos()
	_, t, _ := s.statSvc.GetSynup()
	resp := struct {
		UpdateAt            time.Time             `json:"update_at"`
		TotalVolume         float64               `json:"24h_total_volume"`
		LockVolume          float64               `json:"total_value_locked"`
		TradePairs          []statas.SwapPairInfo `json:"trade_pairs"`
		TotalValueLockedAll float64               `json:"total_value_locked_all"`
	}{
		updateAt,
		totalVolume,
		lockVolume,
		swapPiars,
			t + lockVolume,
	}
	jsonBytes, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_, err = w.Write(jsonBytes)
	if err != nil {
		util.Logger.Errorf("write response error, err=%s", err.Error())
	}
}

func (s *Server) Price(w http.ResponseWriter, r *http.Request) {
	prices, updateAt := s.statSvc.GetPrice()
	resp := struct {
		UpdateAt time.Time          `json:"update_at"`
		Prices   map[string]float64 `json:"prices"`
	}{
		updateAt,
		prices,
	}
	jsonBytes, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_, err = w.Write(jsonBytes)
	if err != nil {
		util.Logger.Errorf("write response error, err=%s", err.Error())
	}
}
func (s *Server) Syrup(w http.ResponseWriter, r *http.Request) {
	synups, tvl, updateAt := s.statSvc.GetSynup()
	resp := struct {
		UpdateAt time.Time  `json:"update_at"`
		TVL      float64    `json:"tvl"`
		Pools    []statas.SyrupTVL `json:"pools"`
	}{
		updateAt,
		tvl,
		synups,
	}
	jsonBytes, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_, err = w.Write(jsonBytes)
	if err != nil {
		util.Logger.Errorf("write response error, err=%s", err.Error())
	}
}

func (s *Server) Healthz(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (s *Server) Serve() {
	router := mux.NewRouter()

	router.HandleFunc("/healthz", s.Healthz).Methods("GET")
	router.HandleFunc("/api/v1/stat", s.Stat).Methods("GET")
	router.HandleFunc("/api/v1/price", s.Price).Methods("GET")
	router.HandleFunc("/api/v1/syrup", s.Syrup).Methods("GET")

	listenAddr := DefaultListenAddr
	if s.config.ServerConfig.ListenAddr != "" {
		listenAddr = s.config.ServerConfig.ListenAddr
	}
	srv := &http.Server{
		Handler:      router,
		Addr:         listenAddr,
		WriteTimeout: 3 * time.Second,
		ReadTimeout:  3 * time.Second,
	}

	util.Logger.Infof("start admin server at %s", srv.Addr)

	err := srv.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("start admin server error, err=%s", err.Error()))
	}
}
