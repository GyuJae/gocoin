package rest

import (
	"fmt"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gyujae/gocoin/blockchain"
	"github.com/gyujae/gocoin/libs"
)

type AddBlockBody struct {
	Message string `json:"message"`
}

func block(c *fiber.Ctx) error {
	height, err := strconv.Atoi(c.Params("height"))
	libs.HandleErr(err)
	block, err := blockchain.GetBlockchain().GetBlock(height)
	libs.HandleErr(err)
	return c.JSON(block)
}

func homeHandler(c *fiber.Ctx) error {
	return c.JSON(c.App().Stack())
}

func blocks(c *fiber.Ctx) error {
	c.Set("Content-Type", "application/json")
	switch c.Method() {
	case "GET":
		return c.JSON(blockchain.GetBlockchain().AllBlocks())
	case "POST":
		addBlockBody := new(AddBlockBody)

		if err := c.BodyParser(addBlockBody); err != nil {
			return err
		}
		blockchain.GetBlockchain().AddBlock(addBlockBody.Message)
		return c.JSON(blockchain.GetBlockchain().AllBlocks())
	}

	return nil
}

func Start(port int) {
	app := fiber.New()
	app.Use(logger.New())
	app.Get("/", homeHandler).Name("See Documentation")
	app.Get("/blocks", blocks).Name("See All Blocks")
	app.Post("/blocks", blocks).Name("Add A Block")
	app.Get("/blocks/{hash:[a-f0-9]+}", block).Name("See A Block")

	log.Fatal(app.Listen(fmt.Sprintf(":%d", port)))
}
