package main

import (
	"fmt"
	"log"
	"net/http"
	"piskworky/tictactoe"

	"github.com/gin-gonic/gin"
)

var game tictactoe.Game

type playgroundSizeDto struct {
	Rows    uint16 `json:"rows"`
	Columns uint16 `json:"columns"`
}

type PositionDto struct {
	Row    uint16 `json:"row"`
	Column uint16 `json:"column"`
}

type TurnDto struct {
	Position PositionDto `json:"position"`
	Status   string      `json:"status"`
}

func startGameHandler(context *gin.Context) {
	size := playgroundSizeDto{}
	if err := context.BindJSON(&size); err != nil {
		log.Println(err.Error())
		context.String(http.StatusOK, err.Error())
	} else {
		game = tictactoe.NewGame(size.Rows, size.Columns)
		context.JSON(http.StatusOK, size)
	}
}

func oponentMoveHandler(context *gin.Context) {

	if game == nil {
		context.String(http.StatusBadRequest, "Game not started")
		return
	}

	move := &PositionDto{}

	if err := context.BindJSON(&move); err != nil {
		context.String(http.StatusBadRequest, "Json { row: X, column: Y } expected.")
		return
	}

	status := game.OpponentMove(tictactoe.Position{move.Row, move.Column})
	context.JSON(http.StatusOK, TurnDto{*move, translateStatus(status)})
}

func translateStatus(status tictactoe.GameStatus) string {
	switch status {
	case tictactoe.OPONNENT_WON:
		return "OPONENT_WON"
	case tictactoe.AI_WON:
		return "AI_WON"
	case tictactoe.TIE:
		return "TIE"
	case tictactoe.PLAYING:
		return "PLAYING"
	default:
		panic(fmt.Sprintf("Unknown status %d", status))

	}
}

func aiMoveHandler(context *gin.Context) {

	if game == nil {
		context.String(http.StatusBadRequest, "Game not started")
		return
	}

	status, pos := game.AiMove()

	context.JSON(http.StatusOK, &TurnDto{
		Position: PositionDto{pos.Row, pos.Column},
		Status:   translateStatus(status),
	})
}

func nextPositions(context *gin.Context) {
	if game == nil {
		context.String(http.StatusBadRequest, "Game not started")
		return
	}

	positions := game.NextPositions()
	dtos := make([]PositionDto, 0, len(positions))

	for _, p := range positions {
		dtos = append(dtos, PositionDto{Row: p.Row, Column: p.Column})
	}

	context.JSON(http.StatusOK, dtos)
}
